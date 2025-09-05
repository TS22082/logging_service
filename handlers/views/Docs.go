package views

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ts22082/logging-service/templates/pages"
)

func Docs(w http.ResponseWriter, r *http.Request) {

	subjects := [5]string{"accounts", "projects", "apiKeys", "logs", "invites"}

	vars := mux.Vars(r)
	subject := vars["subject"]

	component := pages.DocsPage(subject, subjects)
	err := component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
