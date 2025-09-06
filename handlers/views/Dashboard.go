package views

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ts22082/logging-service/templates/pages"
	"github.com/ts22082/logging-service/utils"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectUserRel struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserEmail string             `json:"userEmail" bson:"userEmail"`
	ProjectId string             `json:"projectId" bson:"projectId"`
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ValidateToken(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	userProjectRelCollection := mongodb_client.GetCollection("User_Project_Rel")
	projectCollection := mongodb_client.GetCollection("Projects")
	keyProjectRelCollection := mongodb_client.GetCollection("Key_Project_Rel")

	filter := bson.D{{Key: "userEmail", Value: user.Email}}

	var projectArr []pages.ProjectRobust
	var userProjectRel []ProjectUserRel
	var project pages.ProjectRobust

	mongoDbContext, cancel := mongodb_client.GetContext(10 * time.Second)
	defer cancel()

	cursor, err := userProjectRelCollection.Find(mongoDbContext, filter)

	if err != nil {
		http.Error(w, "Error Iterating", http.StatusInternalServerError)
		return
	}

	if err := cursor.All(mongoDbContext, &userProjectRel); err != nil {
		http.Error(w, "Error Iterating", http.StatusInternalServerError)
	}

	for _, value := range userProjectRel {
		var projectIdHex, err = primitive.ObjectIDFromHex(value.ProjectId)

		if err != nil {
			http.Error(w, "Error Iterating", http.StatusInternalServerError)
		}

		filter := bson.D{{Key: "_id", Value: projectIdHex}}

		err = projectCollection.FindOne(mongoDbContext, filter).Decode(&project.Project)

		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "Error Iterating", http.StatusInternalServerError)
		}

		userCountFilter := bson.D{{Key: "projectId", Value: value.ProjectId}}
		userCount, err := userProjectRelCollection.CountDocuments(mongoDbContext, userCountFilter)

		if err != nil {
			fmt.Println("problem here")
		}

		keyCount, err := keyProjectRelCollection.CountDocuments(mongoDbContext, userCountFilter)

		if err != nil {
			fmt.Println("problem getting keys")
		}

		project.KeyCount = keyCount
		project.UserCount = userCount

		projectArr = append(projectArr, project)
	}

	component := pages.DashboardPage(projectArr)
	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed rendering template", http.StatusInternalServerError)
		return
	}
}
