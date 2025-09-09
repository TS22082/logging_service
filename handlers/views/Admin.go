package views

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ts22082/logging-service/templates/pages"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Admin(w http.ResponseWriter, r *http.Request) {

	mongoDbContext, cancel := mongodb_client.GetContext(10 * time.Second)
	defer cancel()

	userProjectRelCollection := mongodb_client.GetCollection("User_Project_Rel")
	keyProjectRelCollection := mongodb_client.GetCollection("ApiKey")
	projectCollection := mongodb_client.GetCollection("Projects")

	userEmail := "ts2208@gmail.com"

	var userProjectRel []ProjectUserRel
	projectsMap := make(map[string]pages.ProjectRobust)

	var project pages.ProjectRobust

	filter := bson.M{"userEmail": userEmail}

	cursor, err := userProjectRelCollection.Find(mongoDbContext, filter)

	if err != nil {
		http.Error(w, "Cannot Get Keys", http.StatusExpectationFailed)
	}

	if err := cursor.All(mongoDbContext, &userProjectRel); err != nil {
		http.Error(w, "Cannot Get Keys", http.StatusExpectationFailed)
	}

	for _, v := range userProjectRel {
		var projectIdHex, err = primitive.ObjectIDFromHex(v.ProjectId)

		if err != nil {
			fmt.Println("Problem parsing project id")
		}

		filter := bson.D{{Key: "_id", Value: projectIdHex}}

		err = projectCollection.FindOne(mongoDbContext, filter).Decode(&project.Project)

		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "No Documents", http.StatusExpectationFailed)
		}

		userCountFilter := bson.D{{Key: "projectId", Value: v.ProjectId}}
		userCount, err := userProjectRelCollection.CountDocuments(mongoDbContext, userCountFilter)

		if err != nil {
			http.Error(w, "Cannot Get User", http.StatusExpectationFailed)
		}

		keyCount, err := keyProjectRelCollection.CountDocuments(mongoDbContext, userCountFilter)

		if err != nil {
			http.Error(w, "Cannot Get Keys", http.StatusExpectationFailed)
		}

		project.KeyCount = keyCount
		project.UserCount = userCount
		projectsMap[v.ProjectId] = project
	}

	component := pages.ProjectsPage(projectsMap)
	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
