package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ts22082/logging-service/utils"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	redis_client "github.com/ts22082/logging-service/utils/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogEntry struct {
	Id        string `json:"id" bson:"id"`
	Message   string `json:"message" bson:"message"`
	Type      string `json:"type" bson:"type"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
	TraceID   string `json:"traceId,omitempty"`
	ProjectId string `json:"projectId,omitempty" bson:"projectId"`
}

type ApiKey struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectId   string             `json:"projectId" bson:"projectId"`
	Token       string             `json:"token" bson:"token"`
	Count       int64              `json:"count" bson:"count"`
	DateCreated time.Time          `json:"dateCreated" bson:"dateCreated"`
}

func CreateLog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	authToken, authErr := utils.RequireBearerToken(r)

	if authErr != nil {
		http.Error(w, "Unauthorized: Missing authorization header", http.StatusUnauthorized)
		return
	}

	keysCollection := mongodb_client.GetCollection("ApiKey")

	var logEntry LogEntry

	if err := json.NewDecoder(r.Body).Decode(&logEntry); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if logEntry.Message == "" || logEntry.Type == "" || logEntry.ProjectId != "" {
		http.Error(w, "Missing required fields: message and type", http.StatusBadRequest)
		return
	}

	logEntry.Timestamp = time.Now().UTC().Format(time.RFC3339)

	filter := bson.D{{Key: "token", Value: authToken}}
	var apiKey ApiKey

	if err := keysCollection.FindOne(r.Context(), filter).Decode(&apiKey); err != nil {
		http.Error(w, "Issue querying DB", http.StatusInternalServerError)
		return
	}

	logEntry.ProjectId = apiKey.ProjectId

	logsCollection := mongodb_client.GetCollection("Logs")
	result, err := logsCollection.InsertOne(r.Context(), logEntry)

	logEntry.Id = result.InsertedID.(primitive.ObjectID).Hex()

	if err != nil {
		http.Error(w, "Issue adding log to DB", http.StatusInternalServerError)
	}

	redisClient := redis_client.GetClient()
	ctx := redis_client.GetContext()

	redisResponseJSON, err := json.Marshal(map[string]interface{}{
		"data":      logEntry,
		"type":      "success",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})

	if err != nil {
		http.Error(w, "JSON Parsing Error", http.StatusInternalServerError)
	}

	redisChannel := "project_logs/" + apiKey.ProjectId

	if err := redisClient.Publish(ctx, redisChannel, redisResponseJSON).Err(); err != nil {
		log.Printf("Failed to publish to Redis: %v", err)
	}

	response := map[string]interface{}{
		"data":      logEntry,
		"type":      "message",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}
