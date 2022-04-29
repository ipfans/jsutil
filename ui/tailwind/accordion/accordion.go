package accordion

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type AccordionItem struct {
	app.Compo
	Body       app.UI
	Title      app.UI
	ID         string
	ParentID   string
	Expandable bool
	AlwaysOpen bool
}

func (a *AccordionItem) Render() app.UI {
	collapseID := "collapse" + a.ID
	headingID := "heading" + a.ID
	expandable := "false"
	if a.Expandable {
		expandable = "true"
	}
	body := app.Div().
		ID("collapse"+a.ID).
		Class("accordion-collapse collapse").
		Aria("labelledby", headingID).
		Body(
			app.Div().
				Class("accordion-body py-4 px-5").
				Body(
					a.Body,
				),
		)
	if a.ParentID != "" {
		body = body.DataSet("bs-parent", "#"+a.ParentID)
	}
	app.Div()
	return app.Div().
		Class("accordion-item bg-white border border-gray-200").
		Body(
			app.H2().
				Class("accordion-header mb-0").
				ID(headingID).
				Body(
					app.Button().
						Class("accordion-button collapsed relative flex items-center w-full py-4 px-5 text-base text-gray-800 text-left bg-white border-0 rounded-none transition focus:outline-none").
						Type("button").
						DataSet("bs-toggle", "collapse").
						DataSet("bs-target", "#"+collapseID).
						Aria("expanded", expandable).
						Aria("controls", collapseID).
						Attr("aria-expanded", expandable).
						Body(
							a.Title,
						),
				),
			body,
		)
}

type Accordion struct {
	app.Compo

	ID       string
	Children []app.UI
}

func (a *Accordion) Render() app.UI {
	div := app.Div().Class("accordion")
	if a.ID != "" {
		div = div.ID(a.ID)
	}
	return div.Body(a.Children...)
}
