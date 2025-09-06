package project_types

type LogEntry struct {
	Message   string `json:"message" bson:"message"`
	Type      string `json:"type" bson:"type"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
	TraceID   string `json:"traceId,omitempty"`
	ProjectId string `json:"projectId,omitempty" bson:"projectId"`
}
