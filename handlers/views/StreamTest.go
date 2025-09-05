package views

import (
	"net/http"

	templates "github.com/ts22082/logging-service/templates/pages"
)

func StreamTest(w http.ResponseWriter, r *http.Request) {
	component := templates.StreamTestPage()
	err := component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
