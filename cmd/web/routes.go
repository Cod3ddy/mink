package main

import (
	"net/http"

	"github.com/cod3ddy/mink/ui"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.HandleFunc("GET /ping", ping)

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /url/shorten", app.shortenUrl)

	standard := alice.New(app.logRequest, app.recoverPanic, commonHeaders)
	return standard.Then(mux)
}