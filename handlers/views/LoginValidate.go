package views

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailLoginToken struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email,omitempty"`
	Token string             `json:"token" bson:"token,omitempty"`
	Exp   time.Time          `json:"exp" bson:"exp,omitempty"`
}

type User struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email,omitempty"`
}

func LoginValidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	loginToken := vars["login_token"]

	emailTokenCollection := mongodb_client.GetCollection("Email_Login_Token")
	mongoDbContext, cancel := mongodb_client.GetContext(10 * time.Second)
	defer cancel()

	var emailLoginToken EmailLoginToken
	tokenFilter := bson.D{{Key: "token", Value: loginToken}}

	if err := emailTokenCollection.FindOne(mongoDbContext, tokenFilter).Decode(&emailLoginToken); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if emailLoginToken.Exp.Before(time.Now().UTC()) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	userCollection := mongodb_client.GetCollection("Users")

	emailFilter := bson.D{{Key: "email", Value: emailLoginToken.Email}}
	var user User

	err := userCollection.FindOne(mongoDbContext, emailFilter).Decode(&user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		user.Email = emailLoginToken.Email
		res, err := userCollection.InsertOne(mongoDbContext, user)
		user.Id = res.InsertedID.(primitive.ObjectID)

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	} else if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	expirationTime := time.Now().UTC().Add(time.Hour * 24).Unix()

	claims := jwt.MapClaims{
		"id":    user.Id.Hex(),
		"email": user.Email,
		"exp":   expirationTime,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	apiUrl := strings.Split(r.Host, ":")[0]

	prod := false

	if apiUrl != "localhost" {
		prod = true
	}

	tokenString, err := jwtToken.SignedString(secret)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	_, tokenDeleteErr := emailTokenCollection.DeleteOne(mongoDbContext, tokenFilter)

	if tokenDeleteErr != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   prod,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
