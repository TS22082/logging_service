package views

import (
	"net/http"
	"strings"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	apiUrl := strings.Split(r.Host, ":")[0]

	prod := false

	if apiUrl != "localhost" {
		prod = true
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   prod,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
