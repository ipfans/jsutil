//go:build !wasm

package jsutil

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// JS2Bytes convert from TypedArray for JS to byte slice for Go.
func JS2Bytes(v app.Value) []byte {
	panic("not implemented")
	return nil
}

// Bytes2JS convert from byte slice for Go to Uint8Array for app.
func Bytes2JS(b []byte) app.Value {
	panic("not implemented")
	return nil
}

// Callback0 make auto-release callback without params.
func Callback0(fn func() interface{}) app.Func {
	panic("not implemented")
	return nil
}

// Callback1 make auto-release callback with 1 param.
func Callback1(fn func(res app.Value) interface{}) app.Func {
	panic("not implemented")
	return nil
}

// CallbackN make auto-release callback with multiple params.
func CallbackN(fn func(res []app.Value) interface{}) app.Func {
	panic("not implemented")
	return nil
}

// WrapError to wrap golang standard error to app.Value.
func WrapError(err error) app.Value {
	panic("not implemented")
	return nil
}

// UnwrapError to unwrap app.Value to golang standard error.
func UnwrapError(v app.Value) error {
	panic("not implemented")
	return nil
}

// IsArray checking value is array type.
func IsArray(item app.Value) bool {
	panic("not implemented")
	return false
}

// JS2Go JS values convert to Go values.
func JS2Go(obj app.Value) interface{} {
	panic("not implemented")
	return nil
}

// Form2Go retrieve form values from form element.
func Form2Go(form app.Value) map[string]interface{} {
	panic("not implemented")
	return nil
}
