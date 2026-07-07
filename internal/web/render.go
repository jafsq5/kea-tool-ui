package web

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, name string, data any) error {

	tpl, err := template.ParseFS(
		Files,
		"templates/layout.html",
		"templates/"+name,
	)

	if err != nil {
		return err
	}

	return tpl.ExecuteTemplate(w, "layout", data)
}

func Static() http.Handler {

	return http.StripPrefix(
		"/static/",
		http.FileServer(http.FS(Files)),
	)

}
