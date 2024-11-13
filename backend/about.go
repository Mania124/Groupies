package backend

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400: Bad Request Method", http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseGlob("frontend/*.html")
	if err != nil {
		http.Error(w, "500: Oops! Something Went Wrong.", http.StatusInternalServerError)
	}
	templ.ExecuteTemplate(w, "about.html", nil)
}
