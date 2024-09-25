package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"authentication_service/internal/proto/authenticator"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type SessionStore struct {
	redisClient *redis.Client
}

// NewSessionController initializes a new SessionStore
func NewSessionStore(redisClient *redis.Client) *SessionStore {
	return &SessionStore{redisClient: redisClient}
}

// CreateSession creates a new session with a unique session ID and stores the UserSession in Redis
func (c *SessionStore) CreateSession(ctx context.Context, userSession *authenticator.UserSession, expTimeInMinutes int) (string, error) {
	sessionID := uuid.NewString()

	// Marshal the UserSession struct into JSON to store in Redis
	sessionData, err := json.Marshal(userSession)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user session: %v", err)
	}

	// Store the session in Redis with an expiration time
	err = c.redisClient.Set(ctx, sessionID, sessionData, time.Duration(expTimeInMinutes)*time.Minute).Err()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}

	return sessionID, nil
}

// GetSession checks if the session ID exists in Redis and returns the UserSession
func (c *SessionStore) GetSession(ctx context.Context, sessionID string) (*authenticator.UserSession, error) {
	// Retrieve the session data from Redis
	sessionData, err := c.redisClient.Get(ctx, sessionID).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("session not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get session: %v", err)
	}

	// Unmarshal the JSON back into the UserSession struct
	var userSession authenticator.UserSession
	err = json.Unmarshal([]byte(sessionData), &userSession)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user session: %v", err)
	}

	return &userSession, nil
}

// DeleteSession deletes the session by session ID
func (c *SessionStore) DeleteSession(ctx context.Context, sessionID string) error {
	// Delete the session from Redis
	err := c.redisClient.Del(ctx, sessionID).Err()
	if err != nil {
		return fmt.Errorf("failed to delete session: %v", err)
	}
	return nil
}
