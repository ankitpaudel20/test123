package main

import (
	"net/http"

	"example.com/assets"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(app.notFound)

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.Handler("GET", "/static/*filepath", fileServer)

	mux.HandlerFunc("GET", "/", app.home)
	mux.HandlerFunc("GET", "/healthz", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(`hello world`)) })

	mux.Handler("GET", "/basic-auth-protected", app.requireBasicAuthentication(http.HandlerFunc(app.protected)))

	return app.logAccess(app.recoverPanic(app.securityHeaders(mux)))
}
