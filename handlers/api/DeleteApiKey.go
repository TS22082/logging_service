package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteApiKey(w http.ResponseWriter, r *http.Request) {
	keysCollection := mongodb_client.GetCollection("ApiKey")

	vars := mux.Vars(r)
	apiKeyId, exists := vars["apiKeyId"]

	if !exists {
		http.Error(w, "Wrong input", http.StatusExpectationFailed)
	}

	var apiIdHex, err = primitive.ObjectIDFromHex(apiKeyId)

	if err != nil {
		http.Error(w, "Cant parse item", http.StatusExpectationFailed)
	}

	deleteKeyFilter := bson.D{{Key: "_id", Value: apiIdHex}}
	_, err = keysCollection.DeleteOne(r.Context(), deleteKeyFilter)

	if err != nil {
		http.Error(w, "Cant parse item", http.StatusExpectationFailed)
	}

	response := map[string]interface{}{
		"success": "true",
		"msg":     "key successfully deleted",
	}

	json.NewEncoder(w).Encode(response)
}
