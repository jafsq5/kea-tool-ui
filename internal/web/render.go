package web

import (
	"html/template"
	"io/fs"
	"net/http"
)

func Render(w http.ResponseWriter, page string, data any) error {

	tmpl, err := template.ParseFS(
		Assets,
		"templates/layout.html",
		"templates/"+page,
	)

	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout", data)
}

func Static() http.Handler {

	staticFS, err := fs.Sub(Assets, "static")
	if err != nil {
		panic(err)
	}

	return http.StripPrefix(
		"/static/",
		http.FileServer(http.FS(staticFS)),
	)
}
