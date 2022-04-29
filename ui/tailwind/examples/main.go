package main

import (
	"log"
	"net/http"

	"github.com/ipfans/jsutil/ui/tailwind"
	"github.com/ipfans/jsutil/ui/tailwind/accordion"
	"github.com/ipfans/jsutil/ui/tailwind/alert"
	"github.com/ipfans/jsutil/ui/tailwind/icons"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

func (a *hello) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		app.Log("Update available. Reloading...")
		ctx.Reload()
	}
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	accordionID := "accordionExample"
	example := &accordion.Accordion{
		ID: accordionID,
		Children: []app.UI{
			&accordion.AccordionItem{
				ID:         "One",
				ParentID:   accordionID,
				Expandable: true,
				Title:      app.Text("One"),
				Body: app.Div().Body(
					app.Strong().
						Body(
							app.Text("This is the second item's accordion body."),
						),
					app.Text("It is hidden by default,\n      until the collapse plugin adds the appropriate classes that we use to style each\n      element. These classes control the overall appearance, as well as the showing and\n      hiding via CSS transitions. You can modify any of this with custom CSS or overriding\n      our default variables. It's also worth noting that just about any HTML can go within\n      the"),
					app.Code().
						Body(
							app.Text(".accordion-body"),
						),
					app.Text(", though the transition does limit overflow."),
				),
			},
			&accordion.AccordionItem{
				ID:       "Two",
				ParentID: accordionID,
				Title:    app.Text("Two"),
				Body: app.Div().Body(
					app.Strong().
						Body(
							app.Text("This is the second item's accordion body."),
						),
					app.Text("It is hidden by default,\n      until the collapse plugin adds the appropriate classes that we use to style each\n      element. These classes control the overall appearance, as well as the showing and\n      hiding via CSS transitions. You can modify any of this with custom CSS or overriding\n      our default variables. It's also worth noting that just about any HTML can go within\n      the"),
					app.Code().
						Body(
							app.Text(".accordion-body"),
						),
					app.Text(", though the transition does limit overflow."),
				),
			},
		},
	}

	return app.Div().Body(
		example,
		&alert.Alert{
			Icon:     icons.IconSuccess,
			Body:     app.Text("This is an alert"),
			Color:    alert.AlertColorWarning,
			Closable: true,
		},
		tailwind.TailwindElementsJS,
	)
}

func main() {
	app.Route("/", &hello{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles:      tailwind.TailwindElementsCSS,
		RawHeaders: []string{
			tailwind.TailwindScript,
			tailwind.TailwindElementsConfig,
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
