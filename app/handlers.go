package app

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func (a *App) handlerHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			errorHandler(w, r, http.StatusNotFound)

			return
		}

		viewCount, err := a.db.queryHomeViewData()
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)

			return
		}

		data := struct {
			Title     string
			ViewCount int
		}{Title: "Home Page", ViewCount: viewCount}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		err = a.template.ExecuteTemplate(w, "home.go.tmpl", data)
		if err != nil {
			log.Println(err)
		}
	}
}

// Handles static files, like CSS.
func (a *App) handlerStatic() http.HandlerFunc {
	staticFS := http.FS(staticFiles)
	fs := http.FileServer(staticFS)

	return func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "4😢4")
	}
}
