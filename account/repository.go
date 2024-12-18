package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"  // For password hashing and validation

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

// Repository defines the interface for interacting with the accounts database.
type Repository interface {
	Close()                                          // Close the database connection
	PutAccount(ctx context.Context, account Account) error  // Insert a new account
	GetAccountByEmailAndPassword(ctx context.Context, email, password string) (*Account, error) // Retrieve an account by email and password
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) // List accounts with pagination
	Ping() error                                  // Check database connection health
}

// postgresRepository is the PostgreSQL implementation of the Repository interface.
type postgresRepository struct {
	db *sql.DB // SQL database connection
}

// NewPostgresRepository creates and initializes a new postgresRepository instance.
func NewPostgresRepository(url string) (Repository, error) {
	// Open a database connection
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Check the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Configure connection pooling
	db.SetMaxOpenConns(25)          // Maximum number of open connections
	db.SetMaxIdleConns(5)           // Maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Lifetime of a single connection

	return &postgresRepository{db}, nil
}

// Close releases the database connection resources.
func (r *postgresRepository) Close() {
	r.db.Close()
}

// Ping checks the health of the database connection.
func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

// PutAccount inserts a new account into the accounts table with hashed password.
func (r *postgresRepository) PutAccount(ctx context.Context, account Account) error {
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Prepare the SQL query to insert a new account
	query := `
		INSERT INTO accounts (id, name, email, password, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = r.db.ExecContext(ctx, query, account.ID, account.Name, account.Email, string(hashedPassword), account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}
	return nil
}

// GetAccountByEmailAndPassword retrieves an account by its email and validates its password.
func (r *postgresRepository) GetAccountByEmailAndPassword(ctx context.Context, email, password string) (*Account, error) {
	// Check if the email exists in the database
	query := `
		SELECT id, name, email, password, created_at, updated_at 
		FROM accounts 
		WHERE email = $1
	`
	row := r.db.QueryRowContext(ctx, query, email)

	var account Account
	if err := row.Scan(&account.ID, &account.Name, &account.Email, &account.Password, &account.CreatedAt, &account.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("email not found: %w", err)
		}
		return nil, fmt.Errorf("failed to query account: %w", err)
	}

	// Validate the password
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &account, nil
}

// ListAccounts retrieves a paginated list of accounts from the accounts table.
func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	query := `
		SELECT id, name, email 
		FROM accounts 
		ORDER BY id DESC 
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, take, skip)
	if err != nil {
		return nil, fmt.Errorf("failed to query accounts: %w", err)
	}
	defer rows.Close()

	accounts := []Account{}
	for rows.Next() {
		var a Account
		if err := rows.Scan(&a.ID, &a.Name, &a.Email); err != nil {
			return nil, fmt.Errorf("failed to scan account: %w", err)
		}
		accounts = append(accounts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return accounts, nil
}
