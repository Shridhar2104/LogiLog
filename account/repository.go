package account

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	PutAccount(ctx context.Context, account Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
	
}

func NewPostgresRepository(url string) (*postgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()

}

func (r *postgresRepository) PutAccount(ctx context.Context, account Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts (id, name, email, password, shop_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", account.ID, account.Name, account.Email, account.Password, account.ShopName, account.CreatedAt, account.UpdatedAt)
	return err
	
}

func (r *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {

	row := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM accounts WHERE id = $1", id)
	a := &Account{}
	if err := row.Scan(&a.ID, &a.Name, &a.Email); err != nil {
		return nil, err
	}
	return a, nil
}
func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email FROM accounts Order by id DESC LIMIT $1 OFFSET $2", take, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []Account{}
	for rows.Next() {
		var a Account
		if err := rows.Scan(&a.ID, &a.Name, &a.Email); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}
