//go:build !wasm

package jsutil

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Fetch wrapper fetch function.
func Fetch(url string, opt map[string]interface{}) (app.Value, error) {
	panic("not implemented")
	return nil, nil
}
