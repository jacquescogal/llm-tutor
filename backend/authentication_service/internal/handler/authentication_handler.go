package handler

import (
	"authentication_service/internal/controller"
	"authentication_service/internal/proto/authenticator"
	"context"
	pb "authentication_service/internal/proto/authenticator"
)

// AuthenticationHandler represents the handler for managing authentication
// Holds a reference to the authentication controller
type AuthenticationHandler struct {
	// UserRepo
	authenticationController *controller.AuthenticationController
	pb.UnimplementedUserServiceServer
}

// mustEmbedUnimplementedUserServiceServer implements authenticator.UserServiceServer.
func (authenticationHandler *AuthenticationHandler) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}

// NewAuthenticationHandler creates a new AuthenticationHandler
func NewAuthenticationHandler(authenticationController *controller.AuthenticationController) *AuthenticationHandler {
	return &AuthenticationHandler{authenticationController: authenticationController}
}

// RegisterUser registers a new user
func (authenticationHandler *AuthenticationHandler) CreateUser(ctx context.Context, createUserRequest *authenticator.CreateUserRequest) (*authenticator.CreateUserResponse, error) {
	return authenticationHandler.authenticationController.RegisterUser(ctx, createUserRequest)
}

// CreateSession creates a new session with a unique session ID and stores the UserSession in Redis
func (authenticationHandler *AuthenticationHandler) CreateSession(ctx context.Context, createUserSessionRequest *authenticator.CreateSessionRequest) (*authenticator.CreateSessionResponse, error) {
	return authenticationHandler.authenticationController.CreateSession(ctx, createUserSessionRequest)
}

// AuthenticateSession authenticates a user session
func (authenticationHandler *AuthenticationHandler) AuthenticateSession(ctx context.Context, authenticateUserRequest *authenticator.AuthenticateSessionRequest) (*authenticator.AuthenticateSessionResponse, error) {
	return authenticationHandler.authenticationController.AuthenticateSession(ctx, authenticateUserRequest)
}

// DeleteSession deletes a user session
func (authenticationHandler *AuthenticationHandler) DeleteSession(ctx context.Context, deleteUserSessionRequest *authenticator.DeleteSessionRequest) (*authenticator.DeleteSessionResponse, error) {
	return authenticationHandler.authenticationController.DeleteSession(ctx, deleteUserSessionRequest)
}
