package services

import (
	"bff/internal/proto/authenticator"
	"context"
)

type AuthenticationService struct {
    client authenticator.UserServiceClient
}

func NewAuthenticationService(client authenticator.UserServiceClient) *AuthenticationService {
    return &AuthenticationService{client: client}
}


func (s *AuthenticationService) CreateUser(username, password string) error {
    req := &authenticator.CreateUserRequest{Username: username, Password: password}
    _, err := s.client.CreateUser(context.Background(), req)
    if err != nil {
        return HandleGRPCError(err)
    }
    return nil
}

func (s *AuthenticationService) CreateSession(username, password string) (string, error) {
    req := &authenticator.CreateSessionRequest{Username: username, Password: password}
    resp, err := s.client.CreateSession(context.Background(), req)
    if err != nil {
        return "", HandleGRPCError(err)
    }
    return resp.SessionId, nil
}

func (s *AuthenticationService) DeleteSession(sessionID string) error {
    req := &authenticator.DeleteSessionRequest{SessionId: sessionID}
    _, err := s.client.DeleteSession(context.Background(), req)
    if err != nil  {
        return HandleGRPCError(err)
    }
    return nil
}

func (s *AuthenticationService) AuthenticateSession(sessionID string) (*authenticator.UserSession, error) {
    req := &authenticator.AuthenticateSessionRequest{SessionId: sessionID}
    resp, err := s.client.AuthenticateSession(context.Background(), req)
    if err != nil {
        return nil, HandleGRPCError(err)
    }
    return resp.UserSession, nil
}
