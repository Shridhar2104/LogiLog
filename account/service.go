package account

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateAccount(ctx context.Context, name string, password string, email string, shopName string) (*Account, error)
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID        uuid.UUID `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	ShopName  string `json:"shopName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type accountService struct {
	repo Repository
	
}

func NewAccountService(repo Repository) Service {
	return &accountService{repo}
}

func (s *accountService) CreateAccount(ctx context.Context, name string, password string, email string, shopName string) (*Account, error) {
	a := &Account{
		ID:        uuid.New(),
		Name:      name,
		Password:  password,
		Email:     email,
		ShopName:  shopName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	return s.repo.GetAccountByID(ctx, id)
}

func (s *accountService) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {

	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repo.ListAccounts(ctx, skip, take)
}

