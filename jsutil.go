//go:build wasm

package jsutil

import (
	"errors"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	global = app.Window()
	array  = global.Get("Array")
	object = global.Get("Object")
)

// JS2Bytes convert from TypedArray for JS to byte slice for Go.
func JS2Bytes(v app.Value) []byte {
	b := make([]byte, v.Get("byteLength").Int())
	app.CopyBytesToGo(b, global.Get("Uint8Array").New(v.Get("buffer")))
	return b
}

// Bytes2JS convert from byte slice for Go to Uint8Array for app.
func Bytes2JS(b []byte) app.Value {
	res := global.Get("Uint8Array").New(len(b))
	app.CopyBytesToJS(res, b)
	return res
}

// Callback0 make auto-release callback without params.
func Callback0(fn func() interface{}) app.Func {
	var cb app.Func
	cb = app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		defer cb.Release()
		return fn()
	})
	return cb
}

// Callback1 make auto-release callback with 1 param.
func Callback1(fn func(res app.Value) interface{}) app.Func {
	var cb app.Func
	cb = app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		defer cb.Release()
		return fn(args[0])
	})
	return cb
}

// CallbackN make auto-release callback with multiple params.
func CallbackN(fn func(res []app.Value) interface{}) app.Func {
	var cb app.Func
	cb = app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		defer cb.Release()
		return fn(args)
	})
	return cb
}

// WrapError to wrap golang standard error to app.Value.
func WrapError(err error) app.Value {
	return global.Get("Error").New(err.Error())
}

// UnwrapError to unwrap app.Value to golang standard error.
func UnwrapError(v app.Value) error {
	return errors.New(v.Call("toString").String())
}

// IsArray checking value is array type.
func IsArray(item app.Value) bool {
	return array.Call("isArray", item).Bool()
}

// JS2Go JS values convert to Go values.
func JS2Go(obj app.Value) interface{} {
	switch obj.Type() {
	default:
		return obj
	case app.TypeBoolean:
		return obj.Bool()
	case app.TypeNumber:
		return obj.Float()
	case app.TypeString:
		return obj.String()
	case app.TypeObject:
		if IsArray(obj) {
			res := []interface{}{}
			for i := 0; i < obj.Length(); i++ {
				res = append(res, obj.Index(i))
			}
			return res
		}
		res := map[string]interface{}{}
		entries := object.Call("entries", obj)
		for i := 0; i < entries.Length(); i++ {
			entry := entries.Index(i)
			key, value := entry.Index(0).String(), entry.Index(1)
			res[key] = JS2Go(value)
		}
		return res
	}
}

// Form2Go retrieve form values from form element.
func Form2Go(form app.Value) map[string]interface{} {
	obj := map[string]interface{}{}
	formData := global.Get("FormData").New(form)
	fp := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		value, key := JS2Go(args[0]), args[1].String()
		if _, ok := obj[key]; !ok {
			obj[key] = value
			return nil
		}
		v, ok := obj[key].([]interface{})
		if !ok {
			v = []interface{}{obj[key]}
		}
		v = append(v, value)
		obj[key] = v
		return nil
	})
	defer fp.Release()
	formData.Call("forEach", fp)
	return obj
}
