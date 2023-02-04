package main

import (
	"github.com/fragforce/event-manager/pkg/event"
	"github.com/fragforce/event-manager/pkg/webserver"
	"log"
)

func main() {
	err := event.LoadEvents()
	if err != nil {
		log.Panicln(err)
	}
	webserver.Start()
}
