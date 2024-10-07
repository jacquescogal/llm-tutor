package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

type CacheStore struct {
	redisClient *redis.Client
}

// NewSessionController initializes a new SessionStore
func NewCacheStore(redisClient *redis.Client) *CacheStore {
	return &CacheStore{redisClient: redisClient}
}

// StoreData stores the data in Redis with the specified key and expiration time
func (c *CacheStore) StoreData(ctx context.Context, key string, message proto.Message, expTimeInMinutes int) error {
	// Marshal the proto message into JSON
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal proto message: %v", err)
	}

	// Store the data in Redis with an expiration time
	err = c.redisClient.Set(ctx, key, data, time.Duration(expTimeInMinutes)*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("failed to store data: %v", err)
	}

	return nil
}

// RetrieveData retrieves the data from Redis with the specified key
func (c *CacheStore) RetrieveData(ctx context.Context, key string, message proto.Message) error {
	// Retrieve the data from Redis
	data, err := c.redisClient.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return fmt.Errorf("data not found")
	} else if err != nil {
		return fmt.Errorf("failed to get data: %v", err)
	}
	
	// Unmarshal the JSON back into the proto message
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("failed to unmarshal proto message: %v", err)
	}

	return nil
}

// UpdateData updates the data in Redis with the specified key and expiration time
func (c *CacheStore) UpdateData(ctx context.Context, key string, message proto.Message, expTimeInMinutes int) error {
	// Marshal the proto message into JSON
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal proto message: %v", err)
	}

	// Update the data in Redis with an expiration time
	err = c.redisClient.Set(ctx, key, data, time.Duration(expTimeInMinutes)*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("failed to update data: %v", err)
	}

	return nil
}

// DeleteData deletes the data from Redis with the specified key
func (c *CacheStore) DeleteData(ctx context.Context, key string) error {
	// Delete the data from Redis
	// If error then the redis had an internal error
	err := c.redisClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete data: %v", err)
	}

	return nil
}