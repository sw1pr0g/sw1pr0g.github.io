package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sw1pr0g/sw1pr0g.github.io/src/components"
	"github.com/sw1pr0g/sw1pr0g.github.io/src/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Route("/", &components.Home{})
	app.RunWhenOnBrowser()

	h := app.Handler{
		Author:      "Alex Poryadin (sw1pr0g)",
		Description: "sw1pr0g's personal page",
		Icon: app.Icon{
			Default: "/web/logo.png",
		},
		LoadingLabel: "sw1pr0g's personal page loading..",
		Name:         "sw1pr0g",
		Image:        "/web/logo.png",
		Title:        "sw1pr0g",
	}

	switch cfg.Launch.Mode {
	case "local":
		runLocal(&h, cfg.HTTP.Port)

	case "github":
		generateGithubPages(cfg.App.StaticDir, &h)
	}
}

func runLocal(h http.Handler, port int) {
	s := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: h,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func generateGithubPages(dir string, h *app.Handler) {
	if err := app.GenerateStaticWebsite(dir, h); err != nil {
		panic(err)
	}
}
