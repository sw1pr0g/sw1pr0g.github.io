package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
	"os"
)

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.P().
			Text("sw1prog's home.. maybe you were waiting for the design, but I don't care about that"),
		app.P().Body(
			app.Text("i have a "),
			app.A().
				Text("telegram channel").
				Href("https://t.me/sw1pr0g_channel").
				Target("_blank"),
		),
		app.P().Body(
			app.Text("check out my "),
			app.A().
				Text("github").
				Href("https://github.com/sw1pr0g").
				Target("_blank"),
		),
	)
}

func main() {
	app.Route("/", &Home{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("docs", &app.Handler{
		Author:      "Alex Poryadin",
		Description: "sw1pr0g personal page",
		Name:        "sw1pr0g",
		Title:       "sw1pr0g",
	})
	if err != nil {
		log.Fatalf("Failed to generate static website: %v", err)
		os.Exit(1)
	}

	h := app.Handler{
		Author:      "Alex Poryadin",
		Description: "sw1pr0g personal page",
		Icon: app.Icon{
			Default: "/web/logo.png",
		},
		Name:  "sw1pr0g",
		Title: "sw1pr0g",
	}

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", "8080"),
		Handler: &h,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
