package repository

import (
	"authentication_service/internal/proto/authenticator"
	"context"
	"database/sql"
	"log"
)

// UserRepo represents the repository to manage users
type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser creates a new user
func (userRepository *UserRepository) CreateUser(ctx context.Context, username, hashSaltPassword string) error {
	// Example SQL statement for creating a user
	query := "INSERT INTO user_tab (username, hash_salt_password) VALUES (?, ?)"

	// Use ExecContext to pass the context, so the query respects timeouts/cancellations
	_, err := userRepository.db.ExecContext(ctx, query, username, hashSaltPassword)
	if err != nil {
		return err
	}

	log.Println("User created successfully")
	return nil
}

// GetUserByUsername retrieves a user by username
func (userRepository *UserRepository) GetUserByUsername(ctx context.Context, username string) (*authenticator.DBUser, error) {
	// Example SQL statement for getting a user by username
	query := "SELECT user_id, username, hash_salt_password FROM user_tab WHERE username = ?"

	// Use QueryRowContext to pass the context, so the query respects timeouts/cancellations
	row := userRepository.db.QueryRowContext(ctx, query, username)

	var user authenticator.DBUser
	err := row.Scan(&user.Id, &user.Username, &user.HashSaltPassword)
	if err != nil {
		log.Printf("Failed to retrieve user: %v\n", err)
		return nil, err
	}
	log.Println("User retrieved successfully")
	return &user, nil
}
