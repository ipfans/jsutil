package jsutil

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Await equivalent for js await statement.
func Await(promise app.Value) (res app.Value, err error) {
	ch := make(chan bool)
	promise.Call("then",
		Callback1(func(arg app.Value) interface{} {
			res = arg
			close(ch)
			return nil
		}),
	).Call("catch",
		Callback1(func(arg app.Value) interface{} {
			err = UnwrapError(arg)
			close(ch)
			return nil
		}),
	)
	<-ch
	return
}
