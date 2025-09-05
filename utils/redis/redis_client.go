package redis_client

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

// Config holds Redis configuration
type Config struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
}

// DefaultConfig returns default Redis configuration
func DefaultConfig() *Config {
	return &Config{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
	}
}

// Init initializes the Redis connection
func Init(config *Config) error {
	if config == nil {
		config = DefaultConfig()
	}

	// Create Redis client
	client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,

		// Connection pool settings
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,

		// Timeouts
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,

		// Retry settings
		MaxRetries:      3,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,
	})

	// Test the connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	fmt.Printf("Connected to Redis: %s\n", pong)
	return nil
}

// GetClient returns the Redis client
func GetClient() *redis.Client {
	return client
}

// GetContext returns the context for Redis operations
func GetContext() context.Context {
	return ctx
}

// Close closes the Redis connection
func Close() error {
	if client != nil {
		err := client.Close()
		if err != nil {
			return fmt.Errorf("error closing Redis connection: %v", err)
		}
		fmt.Println("Redis connection closed")
	}
	return nil
}

// SetupGracefulShutdown sets up graceful shutdown for Redis
func SetupGracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("\nShutting down server...")

		if err := Close(); err != nil {
			log.Printf("Error during Redis shutdown: %v", err)
		}

		fmt.Println("Server stopped")
		os.Exit(0)
	}()
}

// Health checks Redis connectivity
func Health() (string, error) {
	if client == nil {
		return "", fmt.Errorf("Redis client not initialized")
	}

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return "", fmt.Errorf("Redis ping failed: %v", err)
	}

	return pong, nil
}
