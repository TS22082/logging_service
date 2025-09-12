package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
)

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projectsCollection := mongodb_client.GetCollection("Projects")
	userProjectRelCollection := mongodb_client.GetCollection("User_Project_Rel")
	apiKeyCollection := mongodb_client.GetCollection("ApiKey")

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	var projectIdHex, err = primitive.ObjectIDFromHex(projectId)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	g, _ctx := errgroup.WithContext(r.Context())

	deleteProjectFilter := bson.D{{Key: "_id", Value: projectIdHex}}
	deleteRelFilter := bson.D{{Key: "projectId", Value: projectId}}

	g.Go(func() error {
		_, err := projectsCollection.DeleteOne(_ctx, deleteProjectFilter)
		return err
	})

	g.Go(func() error {
		_, err := userProjectRelCollection.DeleteOne(_ctx, deleteRelFilter)
		return err
	})

	g.Go(func() error {
		_, err := apiKeyCollection.DeleteOne(_ctx, deleteRelFilter)
		return err
	})

	if err := g.Wait(); err != nil {
		http.Error(w, "Error deleting Items", http.StatusBadRequest)
	}

	response := map[string]bool{
		"success": true,
	}

	json.NewEncoder(w).Encode(response)
}
