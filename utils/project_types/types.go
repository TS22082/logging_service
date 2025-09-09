package project_types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogEntry struct {
	Message   string `json:"message" bson:"message"`
	Type      string `json:"type" bson:"type"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
	TraceID   string `json:"traceId,omitempty"`
	ProjectId string `json:"projectId,omitempty" bson:"projectId"`
}

type ProjectUserRel struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserEmail string             `json:"userEmail" bson:"userEmail"`
	ProjectId string             `json:"projectId" bson:"projectId"`
}

type Project struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Plan        string             `json:"plan" bson:"plan"`
	ApiKeyCount int                `json:"apiKeyCount" bson:"apiKeyCount"`
}

type ProjectRobust struct {
	Project
	KeyCount  int64 `json:"keyCount" bson:"keyCount"`
	UserCount int64 `json:"userCount" bson:"userCount"`
}

type User struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email,omitempty"`
}

type EmailBody struct {
	Email string `json:"email" bson:"email,omitempty"`
}

type EmailLoginToken struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email,omitempty"`
	Token string             `json:"token" bson:"token,omitempty"`
	Exp   time.Time          `json:"exp" bson:"exp,omitempty"`
}

type ApiKey struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectId   string             `json:"projectId" bson:"projectId"`
	Token       string             `json:"token" bson:"token"`
	Count       int64              `json:"count" bson:"count"`
	DateCreated time.Time          `json:"dateCreated" bson:"dateCreated"`
}
