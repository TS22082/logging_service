package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"github.com/ts22082/logging-service/utils/project_types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostApiKeyRequestBody struct {
	ProjectId string `json:"projectId" bson:"projectId"`
}

func CreateApiKey(w http.ResponseWriter, r *http.Request) {

	request := new(PostApiKeyRequestBody)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	mongodbContext, cancel := mongodb_client.GetContext(10 * time.Second)
	defer cancel()

	apiKeyCollection := mongodb_client.GetCollection("ApiKey")

	apiKey := new(project_types.ApiKey)
	token := uuid.New()

	apiKey.Count = 0
	apiKey.ProjectId = strings.TrimRight(request.ProjectId, "?")
	apiKey.Token = token.String()
	apiKey.DateCreated = time.Now()

	newApiKeyRes, err := apiKeyCollection.InsertOne(mongodbContext, apiKey)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	apiKey.Id = newApiKeyRes.InsertedID.(primitive.ObjectID)

	json.NewEncoder(w).Encode(apiKey)
}
