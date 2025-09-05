package views

import (
	"net/http"

	"github.com/ts22082/logging-service/templates/pages"
)

func Login(w http.ResponseWriter, r *http.Request) {
	component := pages.LoginPage()
	err := component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
