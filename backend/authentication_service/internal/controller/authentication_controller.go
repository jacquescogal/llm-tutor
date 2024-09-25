package controller

import (
	"authentication_service/internal/cache"
	"authentication_service/internal/proto/authenticator"
	"authentication_service/internal/repository"
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
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
	username := createUserRequest.GetUsername()
	password := createUserRequest.GetPassword()

	// Hash password
	hashedSaltedPassword, err := hashSaltPassword(password)
	if err != nil {
		log.Printf("Failed to hash password: %v\n", err)
		return &authenticator.CreateUserResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_INVALID_PASSWORD, Message: "failed to hash password"}}, err
	}

	maxWaitTime := 2
	expTime := 5
	// Acquire lock for consistency
	uniqueID, err := authenticationController.lock.AcquireLock(ctx, username, expTime, maxWaitTime)
	if err != nil {
		log.Printf("Failed to acquire lock for user %s: %v\n", username, err)
		return &authenticator.CreateUserResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_FAILED_TO_ACQUIRE_LOCK, Message: "failed to acquire lock"}}, err
	}else{
		defer func() {
			// Release lock
			err := authenticationController.lock.ReleaseLock(ctx, username, uniqueID)
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}

	// Check if user already exists
	_, err = authenticationController.userRepo.GetUserByUsername(ctx, username)
	if err == nil {
		// User already exists
		log.Printf("User %s already exists\n", username)
		return &authenticator.CreateUserResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_USER_ALREADY_EXISTS, Message: "user already exists"}}, fmt.Errorf("user already exists")
	}

	// Create user
	err = authenticationController.userRepo.CreateUser(ctx, username, hashedSaltedPassword)
	if err != nil {
		log.Printf("Failed to create user %s: %v\n", username, err)
		return &authenticator.CreateUserResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_FAILED_TO_CREATE_USER, Message: "failed to create user"}}, err
	}

	log.Printf("User %s created successfully\n", username)
	return &authenticator.CreateUserResponse{Error: nil}, nil
}

// CreateSession creates a new session
func (authenticationController *AuthenticationController) CreateSession(ctx context.Context, createSessionRequest *authenticator.CreateSessionRequest) (*authenticator.CreateSessionResponse, error) {
	// Parse request
	username := createSessionRequest.GetUsername()
	password := createSessionRequest.GetPassword()

	// Get user
	user, err := authenticationController.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Failed to get user %s: %v\n", username, err)
		return &authenticator.CreateSessionResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_USER_NOT_FOUND, Message: "user not found"}}, err
	}

	// Compare password
	err = compareHashAndPassword(user.GetHashSaltPassword(), password)
	if err != nil {
		log.Printf("Failed to authenticate user %s: %v\n", username, err)
		return &authenticator.CreateSessionResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_INVALID_PASSWORD, Message: "invalid password"}}, err
	}

	// Create session
	userSession := &authenticator.UserSession{Id: user.GetId()}
	expTimeInMinutes := 60 * 24 //TODO: Move to env
	sessionID, err := authenticationController.sessionStore.CreateSession(ctx, userSession, expTimeInMinutes)
	if err != nil {
		log.Printf("Failed to create session for user %s: %v\n", username, err)
		return &authenticator.CreateSessionResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_FAILED_TO_CREATE_SESSION, Message: "failed to create session"}}, err
	}

	log.Printf("Session %s created successfully for user %s\n", sessionID, username)
	return &authenticator.CreateSessionResponse{SessionId: sessionID, Error: nil}, nil
}



// AuthenticateSession authenticates a session
func (authenticationController *AuthenticationController) AuthenticateSession(ctx context.Context, authenticateSessionRequest *authenticator.AuthenticateSessionRequest) (*authenticator.AuthenticateSessionResponse, error) {
	// Parse request
	sessionID := authenticateSessionRequest.GetSessionId()

	// Check if session exists
	userSession, err := authenticationController.sessionStore.GetSession(ctx, sessionID)
	if err != nil {
		log.Printf("Failed to get session %s: %v\n", sessionID, err)
		return &authenticator.AuthenticateSessionResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_SESSION_NOT_FOUND, Message: "session not found"}}, err
	}

	log.Printf("Session %s authenticated successfully\n", sessionID)
	return &authenticator.AuthenticateSessionResponse{UserSession: userSession, Error: nil}, nil
}

// DeleteSession deletes a session
func (authenticationController *AuthenticationController) DeleteSession(ctx context.Context, deleteSessionRequest *authenticator.DeleteSessionRequest) (*authenticator.DeleteSessionResponse, error) {
	// Parse request
	sessionID := deleteSessionRequest.GetSessionId()

	// Delete session
	err := authenticationController.sessionStore.DeleteSession(ctx, sessionID)
	if err != nil {
		log.Printf("Failed to delete session %s: %v\n", sessionID, err)
		return &authenticator.DeleteSessionResponse{Error: &authenticator.Error{Code: authenticator.ErrorCode_INVALID_SESSION, Message: "failed to delete session"}}, err
	}

	log.Printf("Session %s deleted successfully\n", sessionID)
	return &authenticator.DeleteSessionResponse{Error: nil}, nil
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