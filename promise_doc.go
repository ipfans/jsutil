//go:build !wasm

package jsutil

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Await equivalent for js await statement.
func Await(promise app.Value) (res app.Value, err error) {
	panic("not implemented")
	return
}
