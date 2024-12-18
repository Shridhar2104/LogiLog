package account

import (
	"context"
	"time"

	"github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt" // Import for password hashing
)

// Service interface defines the operations provided by the Account service.
type Service interface {
	CreateAccount(ctx context.Context, name string, password string, email string) (*Account, error) // Create a new account
	LoginAccount(ctx context.Context, email string, password string) (*Account, error)               // Login and validate account
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)                  // List accounts with pagination
}

// Account struct represents the Account entity in the system.
type Account struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the account
	Name      string    `json:"name"`       // Account holder's name
	Password  string    `json:"password"`   // Account password (hashed ideally)
	Email     string    `json:"email"`      // Account holder's email
	CreatedAt time.Time `json:"created_at"` // Timestamp when the account was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the account was last updated
}

// accountService is a concrete implementation of the Service interface.
type accountService struct {
	repo Repository // Dependency on the Repository interface for database operations
}

// NewAccountService is a constructor for accountService, returning a Service implementation.
func NewAccountService(repo Repository) Service {
	return &accountService{repo}
}

// CreateAccount creates a new account with the provided details and saves it in the database.
func (s *accountService) CreateAccount(ctx context.Context, name string, password string, email string) (*Account, error) {
	a := &Account{
		ID:        uuid.New(),     // Generate a unique identifier using the UUID package
		Name:      name,           // Assign the name provided
		Password:  password,       // Directly pass the plain password for hashing in repository
		Email:     email,          // Assign the email provided
		CreatedAt: time.Now(),     // Set the current timestamp as the creation time
		UpdatedAt: time.Now(),     // Set the current timestamp as the update time
	}

	// Use the repository's PutAccount method to save the account in the database.
	if err := s.repo.PutAccount(ctx, *a); err != nil {
		return nil, err // Return an error if the database operation fails
	}

	return a, nil // Return the created account on success
}

// LoginAccount validates the email and password and returns the account if valid.
func (s *accountService) LoginAccount(ctx context.Context, email string, password string) (*Account, error) {
	// Use the repository to fetch the account by email and validate the password
	account, err := s.repo.GetAccountByEmailAndPassword(ctx, email, password)
	if err != nil {
		return nil, err // Return an error if validation fails
	}

	return account, nil // Return the account on successful validation
}

// ListAccounts retrieves a list of accounts with support for pagination.
func (s *accountService) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	// Apply a limit of 100 accounts per request if the take value exceeds 100 or is not specified.
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	// Delegate the task to the repository's ListAccounts method.
	return s.repo.ListAccounts(ctx, skip, take)
}
