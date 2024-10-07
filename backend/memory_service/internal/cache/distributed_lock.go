package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

)

// DistributedLock that handles acquiring and releasing distributed locks using Redis
type DistributedLock struct {
	redisClient *redis.Client
}

// NewDistributedLock initializes a new DistributedLock with the Redis client
func NewDistributedLock(redisClient *redis.Client) *DistributedLock {
	return &DistributedLock{redisClient: redisClient}
}

// AcquireLock tries to set a lock with a unique ID and retries until max_wait_time is exceeded
func (lock *DistributedLock) AcquireLock(ctx context.Context, key string, expTime, maxWaitTime int) (string, error) {
	uniqueID := uuid.NewString()

	startTime := time.Now()
	for {
		// Try to acquire the lock
		success, err := lock.redisClient.SetNX(ctx, key, uniqueID, time.Duration(expTime)*time.Second).Result()
		if err != nil {
			return "", fmt.Errorf("failed to acquire lock: %v", err)
		}

		if success {
			return uniqueID, nil
		}

		// Check if we've exceeded the maximum wait time
		elapsedTime := time.Since(startTime)
		if elapsedTime.Seconds() >= float64(maxWaitTime) {
			return "", fmt.Errorf("failed to acquire lock: exceeded max wait time of %d seconds", maxWaitTime)
		}

		// Sleep for 0.5 seconds before retrying
		time.Sleep(500 * time.Millisecond)
	}
}

// ReleaseLock tries to release the lock if the unique ID matches
func (lock *DistributedLock) ReleaseLock(ctx context.Context, key, uniqueID string) error {
	// Use a Lua script to ensure atomicity: check if the value matches, and delete if it does
	luaScript := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`
	// Run the Lua script with the key and unique ID
	result, err := lock.redisClient.Eval(ctx, luaScript, []string{key}, uniqueID).Result()
	if err != nil {
		return fmt.Errorf("failed to release lock: %v", err)
	}

	if result == int64(0) {
		return errors.New("failed to release lock: lock not owned or expired")
	}

	return nil
}
