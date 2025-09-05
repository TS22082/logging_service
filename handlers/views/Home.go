package views

import (
	"net/http"

	"github.com/ts22082/logging-service/templates/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	component := pages.HomePage()
	err := component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
