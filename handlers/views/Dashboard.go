package views

import (
	"fmt"
	"net/http"

	"github.com/ts22082/logging-service/templates/pages"
	"github.com/ts22082/logging-service/utils"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ValidateToken(r)

	if err != nil {
		fmt.Println("Error validating token =>", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	fmt.Println("User ==>", user)

	component := pages.DashboardPage()
	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
