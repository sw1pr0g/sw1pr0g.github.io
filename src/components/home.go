package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.P().Body(
			app.Text("sw1prog's personal page.. if you want the design, feel free to make a pull request "),
			app.A().
				Text("here").
				Href("https://github.com/sw1pr0g/sw1pr0g.github.io"),
			app.Text(", haha"),
		),
		app.P().Body(
			app.Text("i have own blog in "),
			app.A().
				Text("telegram").
				Href("https://t.me/gopherlog"),
		),
		app.P().Body(
			app.Text("check out my "),
			app.A().
				Text("github").
				Href("https://github.com/sw1pr0g"),
		),
	)
}
