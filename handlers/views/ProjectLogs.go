package views

import (
	"net/http"
	"slices"
	"time"

	"github.com/gorilla/mux"
	"github.com/ts22082/logging-service/templates/pages"
	"github.com/ts22082/logging-service/utils"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"github.com/ts22082/logging-service/utils/project_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProjectLogs(w http.ResponseWriter, r *http.Request) {
	_, err := utils.ValidateToken(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	projectCollection := mongodb_client.GetCollection("Projects")
	logsCollection := mongodb_client.GetCollection("Logs")
	mongoDbContext, cancel := mongodb_client.GetContext(10 * time.Second)

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	defer cancel()

	parsedProjectId, err := primitive.ObjectIDFromHex(projectId)

	if err != nil {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	}

	projectFilter := bson.D{{Key: "_id", Value: parsedProjectId}}
	logsFilter := bson.D{{Key: "projectId", Value: projectId}}

	var project pages.ProjectRobust
	var logs []project_types.LogEntry

	if err := projectCollection.FindOne(mongoDbContext, projectFilter).Decode(&project.Project); err != nil {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	cursor, err := logsCollection.Find(mongoDbContext, logsFilter)

	if err != nil {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	if err := cursor.All(mongoDbContext, &logs); err != nil {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	slices.Reverse(logs)

	component := pages.ProjectLogsPage(logs)
	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}

}
