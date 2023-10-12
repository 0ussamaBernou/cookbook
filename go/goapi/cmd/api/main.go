package main

import (
	"cookbook/go/goapi/internal/handlers"
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // logrus aliased as log
	"net/http"
)

func main() {
	log.SetReportCaller(true)

	// route multiplexer that parses the req url
	var r *chi.Mux = chi.NewRouter()

	// handle the request in the handler
	handlers.Handler(r)

	fmt.Println("Starting Go API...")
	// Asked phind for "GO API" ascii art gave me this
	fmt.Println(`
 _____                     ___     ____
|  ___|   ___     ___     |__ \   | __ )    ____   ____
| |_     / _ \   / _ \      / /   |  _ \   / _  | | '__|
|  _|   | (_) | | (_) |    |_|    | |_) | | (_| | | |
|_|      \___/   \___/     (_)    |____/   \____| |_| `)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
