package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// Redis struct holds the Redis client
type Redis struct {
	Client *redis.Client
}

// NewRedis initializes a new Redis connection using environment variables
func NewRedis() (*Redis, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	fmt.Println(redisHost)
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	return &Redis{Client: client}, nil
}

// GetClient returns the Redis client
func (r *Redis) GetClient() *redis.Client {
	return r.Client
}

// Close closes the Redis connection
func (r *Redis) Close() error {
	return r.Client.Close()
}
