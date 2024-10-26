package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sw1pr0g/sw1pr0g.github.io/src/sw1pr0g/components"
	"log"
	"net/http"
)

func main() {
	app.Route("/", &components.Home{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{})

	// log.Printf("App running at: http://localhost:%s", cfg.HTTP.Port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
