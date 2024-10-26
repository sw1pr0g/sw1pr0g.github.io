package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sw1pr0g/sw1pr0g.github.io/src/sw1pr0g/components"
	"net/http"
)

func main() {
	app.Route("/", &components.Home{})
	app.RunWhenOnBrowser()

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
