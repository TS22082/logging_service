package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	redis_client "github.com/ts22082/logging-service/utils/redis"
)

type StreamMessage struct {
	Type      string    `json:"type"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func ProjectLogsStream(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Cache-Control")

	client := redis_client.GetClient()
	ctx := redis_client.GetContext()

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	redisChannel := "project_logs/" + projectId

	pubsub := client.Subscribe(ctx, redisChannel)
	defer pubsub.Close()

	clientGone := r.Context().Done()

	fmt.Fprintf(w, "data: %s\n\n", toSSEData(StreamMessage{
		Type:      "connected",
		Data:      "Connected to stream",
		Timestamp: time.Now(),
	}))

	w.(http.Flusher).Flush()

	heartbeat := time.NewTicker(30 * time.Second)
	defer heartbeat.Stop()

	ch := pubsub.Channel()

	for {
		select {
		case <-clientGone:
			log.Println("SSE client disconnected")
			return

		case <-heartbeat.C:
			fmt.Fprintf(w, "data: %s\n\n", toSSEData(StreamMessage{
				Type:      "heartbeat",
				Data:      "ping",
				Timestamp: time.Now(),
			}))
			w.(http.Flusher).Flush()

		case msg := <-ch:
			fmt.Fprintf(w, "data: %s\n\n", msg.Payload)
			w.(http.Flusher).Flush()

		case <-time.After(5 * time.Minute):
			fmt.Fprintf(w, "data: %s\n\n", toSSEData(StreamMessage{
				Type:      "timeout",
				Data:      "Stream timeout",
				Timestamp: time.Now(),
			}))
			w.(http.Flusher).Flush()
			return
		}
	}
}

func toSSEData(msg StreamMessage) string {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling SSE message: %v", err)
		return `{"type":"error","data":"Failed to marshal message","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`
	}
	return string(data)
}
