package controller

import (
	"authentication_service/internal/cache"
	"authentication_service/internal/proto/authenticator"
	"authentication_service/internal/repository"
	"context"
	"fmt"
	"log"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthenticationController represents the controller for managing authentication
// Holds a reference to the user repository
type AuthenticationController struct {
	// UserRepo
	userRepo *repository.UserRepository

	// Lock
	lock *cache.DistributedLock

	// SessionStore
	sessionStore *cache.SessionStore
}

// NewAuthenticationController creates a new AuthenticationController
func NewAuthenticationController(userRepo *repository.UserRepository, lock *cache.DistributedLock, sessionStore *cache.SessionStore) *AuthenticationController {
	return &AuthenticationController{userRepo: userRepo, lock: lock, sessionStore: sessionStore}
}

// RegisterUser registers a new user
func (authenticationController *AuthenticationController) RegisterUser(ctx context.Context, createUserRequest *authenticator.CreateUserRequest) (*authenticator.CreateUserResponse, error) {
	// Parse request
	username, err := parseUsername(createUserRequest.GetUsername())
	if err != nil {
		log.Printf("Failed to parse username: %v\n", err)
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}
	password, err := parsePassword(createUserRequest.GetPassword())
	if err != nil {
		log.Printf("Failed to parse password: %v\n", err)
		return nil, status.Error(codes.InvalidArgument, "invalid password")
	}

	// Hash password
	hashedSaltedPassword, err := hashSaltPassword(password)
	if err != nil {
		log.Printf("Failed to hash password: %v\n", err)
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	maxWaitTime := 2
	expTime := 5
	// Acquire lock for consistency
	uniqueID, err := authenticationController.lock.AcquireLock(ctx, username, expTime, maxWaitTime)
	if err != nil {
		log.Printf("Failed to acquire lock for user %s: %v\n", username, err)
		return nil, status.Error(codes.ResourceExhausted, "failed to acquire lock")
	}else{
		defer func() {
			// Release lock
			_ = authenticationController.lock.ReleaseLock(ctx, username, uniqueID)
			// if error, then lock only released after expiration of 5(expTime) seconds
		}()
	}

	// Check if user already exists
	_, err = authenticationController.userRepo.GetUserByUsername(ctx, username)
	if err == nil {
		// User already exists
		log.Printf("User %s already exists\n", username)
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	// Create user
	err = authenticationController.userRepo.CreateUser(ctx, username, hashedSaltedPassword)
	if err != nil {
		log.Printf("Failed to create user %s: %v\n", username, err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	log.Printf("User %s created successfully\n", username)
	return &authenticator.CreateUserResponse{}, nil
}

// CreateSession creates a new session
func (authenticationController *AuthenticationController) CreateSession(ctx context.Context, createSessionRequest *authenticator.CreateSessionRequest) (*authenticator.CreateSessionResponse, error) {
	// Parse request
	username, err := parseUsername(createSessionRequest.GetUsername())
	if err != nil {
		log.Printf("Failed to parse username: %v\n", err)
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}
	password, err := parsePassword(createSessionRequest.GetPassword())
	if err != nil {
		log.Printf("Failed to parse password: %v\n", err)
		return nil, status.Error(codes.InvalidArgument, "invalid password")
	}

	// Get user
	user, err := authenticationController.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Failed to get user %s: %v\n", username, err)
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// Compare password
	err = compareHashAndPassword(user.GetHashSaltPassword(), password)
	if err != nil {
		log.Printf("Failed to authenticate user %s: %v\n", username, err)
		return nil, status.Error(codes.Unauthenticated, "failed to authenticate user")
	}

	// Create session
	userSession := &authenticator.UserSession{Id: user.GetId(), Username: username}
	expTimeInMinutes := 60 * 24 //TODO: Move to env
	sessionID, err := authenticationController.sessionStore.CreateSession(ctx, userSession, expTimeInMinutes)
	if err != nil {
		log.Printf("Failed to create session for user %s: %v\n", username, err)
		return nil, status.Error(codes.Internal, "failed to create session")
	}

	log.Printf("Session %s created successfully for user %s\n", sessionID, username)
	return &authenticator.CreateSessionResponse{SessionId: sessionID}, nil
}



// AuthenticateSession authenticates a session
func (authenticationController *AuthenticationController) AuthenticateSession(ctx context.Context, authenticateSessionRequest *authenticator.AuthenticateSessionRequest) (*authenticator.AuthenticateSessionResponse, error) {
	// Parse request
	sessionID := authenticateSessionRequest.GetSessionId()

	// Check if session exists
	userSession, err := authenticationController.sessionStore.GetSession(ctx, sessionID)
	if err != nil {
		log.Printf("Failed to get session %s: %v\n", sessionID, err)
		return nil, status.Error(codes.NotFound, "session not found")
	}

	log.Printf("Session %s authenticated successfully\n", sessionID)
	return &authenticator.AuthenticateSessionResponse{UserSession: userSession}, nil
}

// DeleteSession deletes a session
func (authenticationController *AuthenticationController) DeleteSession(ctx context.Context, deleteSessionRequest *authenticator.DeleteSessionRequest) (*authenticator.DeleteSessionResponse, error) {
	// Parse request
	sessionID := deleteSessionRequest.GetSessionId()

	// Delete session
	err := authenticationController.sessionStore.DeleteSession(ctx, sessionID)
	if err != nil {
		log.Printf("Failed to delete session %s: %v\n", sessionID, err)
		return nil, status.Error(codes.Internal, "failed to delete session")
	}

	log.Printf("Session %s deleted successfully\n", sessionID)
	return &authenticator.DeleteSessionResponse{}, nil
}

func parseUsername(username string) (string, error) {
	// Checks for
	// 1. no spaces
	// 2. Alphanum + . only
	// 3. Change to lowercase
	// 4. 4 to 60 characters only

	isAlphanumeric := func(s string) bool {
		for _, r := range s {
			if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '.' {
				return false
			}
		}
		return true
	}

	if len(username) < 4 || len(username) > 60 {
		return "", fmt.Errorf("username must be between 4 and 60 characters")
	}

	if !isAlphanumeric(username) {
		return "", fmt.Errorf("username can only contain alphanumeric characters and periods")
	}

	return strings.ToLower(username), nil
}

func parsePassword(password string) (string, error) {
	// between 4 and 71 characters
	if len(password) < 4 || len(password) > 71 {
		return "", fmt.Errorf("password must be between 8 and 71 characters")
	}

	return password, nil
}



// hashSaltPassword hashes the password using bcrypt
func hashSaltPassword(password string) (string, error) {
	// Hash & salt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
	
}

// compareHashAndPassword compares the hashed password with the plaintext password
func compareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}