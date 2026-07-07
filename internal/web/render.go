package web

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, page string, data any) error {

	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/"+page,
	)

	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout", data)
}

func Static() http.Handler {

	return http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./static")),
	)

}
