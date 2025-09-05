package mongodb_client

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

type Config struct {
	URI      string
	Database string
	Timeout  time.Duration
}

func Init(config *Config) error {
	// Default configuration
	if config == nil {
		config = &Config{
			URI:      getEnvOrDefault("MONGODB_URI", "mongodb://localhost:27017"),
			Database: getEnvOrDefault("MONGODB_DATABASE", "log_interpreter"),
			Timeout:  30 * time.Second,
		}
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.URI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// Set global variables
	Client = client
	Database = client.Database(config.Database)

	log.Printf("Connected to MongoDB database: %s", config.Database)
	return nil
}

func SetupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutting down MongoDB connection...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := Client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		} else {
			log.Println("MongoDB connection closed.")
		}

		os.Exit(0)
	}()
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Helper functions for common operations
func GetCollection(name string) *mongo.Collection {
	return Database.Collection(name)
}

func GetContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
