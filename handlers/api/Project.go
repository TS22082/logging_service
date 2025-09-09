package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ts22082/logging-service/utils"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	"github.com/ts22082/logging-service/utils/project_types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostProjectRequestBody struct {
	Project string `json:"project" bson:"project"`
	Plan    string `json:"plan" bson:"plan"`
}

func PostProject(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ValidateToken(r)

	mongodbContext, cancel := mongodb_client.GetContext(10 * time.Second)
	defer cancel()

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	projectsCollection := mongodb_client.GetCollection("Projects")
	userProjectRelCollection := mongodb_client.GetCollection("User_Project_Rel")

	request := new(PostProjectRequestBody)
	userEmail := user.Email

	newProject := new(project_types.Project)
	newUserProjectRel := new(project_types.ProjectUserRel)

	newUserProjectRel.UserEmail = userEmail

	newProject.Name = request.Project
	newProject.Plan = request.Plan

	newProjectRes, err := projectsCollection.InsertOne(mongodbContext, newProject)

	if err != nil {
		http.Error(w, "Error Iterating", http.StatusInternalServerError)
	}

	newProject.Id = newProjectRes.InsertedID.(primitive.ObjectID)
	newUserProjectRel.ProjectId = newProject.Id.Hex()

	_, err = userProjectRelCollection.InsertOne(mongodbContext, newUserProjectRel)

	if err != nil {
		http.Error(w, "Error Iterating", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(newProject)
}
