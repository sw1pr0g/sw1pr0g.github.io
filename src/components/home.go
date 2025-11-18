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
				Href("https://github.com/sw1pr0g/sw1pr0g.github.io/pulls"),
			app.Text(", heh"),
		),
		app.P().Body(
			app.Text("i have own "),
			app.A().
				Text("blog").
				Href("https://t.me/dev_unplugged"),
			app.Text(" where i write about coding, sometimes music and other stuff"),
		),
		app.P().Body(
			app.Text("all my talks about software you can find "),
			app.A().
				Text("here").
				Href("https://github.com/sw1pr0g/talks"),
		),
		app.P().Body(
			app.Text("check out my "),
			app.A().
				Text("github").
				Href("https://github.com/sw1pr0g"),
		),
	)
}
