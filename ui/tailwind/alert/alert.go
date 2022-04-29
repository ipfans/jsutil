package alert

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type AlertColor string

var (
	AlertColorIno     AlertColor = "blue"
	AlertColorDanger  AlertColor = "red"
	AlertColorSuccess AlertColor = "green"
	AlertColorWarning AlertColor = "yellow"
)

type Alert struct {
	app.Compo

	Icon     app.UI
	Body     app.UI
	Color    AlertColor
	Closable bool
}

func (c *Alert) Render() app.UI {
	color := string(c.Color)
	body := []app.UI{
		c.Body,
	}
	if c.Closable {
		body = append(body,
			app.Button().
				Type("button").
				Class("btn-close").
				Class("box-content").
				Class("w-4").
				Class("h-4").
				Class("p-1").
				Class("ml-auto").
				Class("text-"+color+"-900").
				Class("border-none").
				Class("rounded-none").
				Class("opacity-50").
				Class("focus:shadow-none").
				Class("focus:outline-none").
				Class("focus:opacity-100").
				Class("hover:text-"+color+"-900").
				Class("hover:opacity-75").
				Class("hover:no-underline").
				DataSet("bs-dismiss", "alert").
				Aria("label", "Close"),
		)
	}
	return app.Div().
		Class("alert").
		Class("bg-"+color+"-100").
		Class("rounded-lg").
		Class("py-5").
		Class("px-6").
		Class("mb-3").
		Class("text-base").
		Class("text-"+color+"-700").
		Class("inline-flex").
		Class("items-center").
		Class("w-full").
		Class("alert-dismissible").
		Class("fade").
		Class("show").
		Aria("role", "alert").
		Body(
			body...,
		)
}
