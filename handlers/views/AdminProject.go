package views

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ts22082/logging-service/templates/pages"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"github.com/ts22082/logging-service/utils/project_types"
	"go.mongodb.org/mongo-driver/bson"
)

func AdminProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId, exists := vars["projectId"]

	if !exists {
		http.Error(w, "Incorrect input", http.StatusExpectationFailed)
	}

	mongoDbContext, cancel := mongodb_client.GetContext(10 * time.Second)
	defer cancel()

	var apiKeys []project_types.ApiKey
	keysCollection := mongodb_client.GetCollection("ApiKey")
	keyFilter := bson.D{{Key: "projectId", Value: projectId}}

	cursor, err := keysCollection.Find(mongoDbContext, keyFilter)

	if err != nil {
		http.Error(w, "Error Iterating", http.StatusInternalServerError)
	}

	if err := cursor.All(mongoDbContext, &apiKeys); err != nil {
		http.Error(w, "DB parse error", http.StatusInternalServerError)
	}

	// Do not expose full api key to front end
	for i, v := range apiKeys {
		apiKeys[i].Token = v.Token[len(v.Token)-4:]
	}

	component := pages.AdminProjectPage(apiKeys)
	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
