package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
)

type EmailLoginToken struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email,omitempty"`
	Token string             `json:"token" bson:"token,omitempty"`
	Exp   time.Time          `json:"exp" bson:"exp,omitempty"`
}

func SendLoginLink(w http.ResponseWriter, r *http.Request) {
	emailLoginTokenCollection := mongodb_client.GetCollection("Email_Login_Token")

	emailLoginToken := new(EmailLoginToken)

	if err := json.NewDecoder(r.Body).Decode(&emailLoginToken); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	token := uuid.New().String()
	emailLoginToken.Token = token

	now := time.Now().UTC()
	expiresAt := now.Add(15 * time.Minute)
	emailLoginToken.Exp = expiresAt

	m := gomail.NewMessage()
	m.SetHeader("From", "ts22082@gmail.com")
	m.SetHeader("To", emailLoginToken.Email)
	m.SetHeader("Subject", "Hello!")

	emailPw := os.Getenv("SERVICE_EMAIL_PW")
	apiUrl := os.Getenv("API_URL")
	linkToLogin := apiUrl + "/email/login_validate/" + emailLoginToken.Token

	body := fmt.Sprintf(`Hello, <a href="%s">Login</a>!`, linkToLogin)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "ts22082@gmail.com", emailPw)

	if err := d.DialAndSend(m); err != nil {
		http.Error(w, "Email Error", http.StatusInternalServerError)
		return
	}

	_, err := emailLoginTokenCollection.InsertOne(r.Context(), emailLoginToken)

	if err != nil {
		http.Error(w, "DB Write Error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data":      "success",
		"type":      "message",
		"success":   true,
		"timestamp": time.Now().UTC(),
	}

	json.NewEncoder(w).Encode(response)
}
