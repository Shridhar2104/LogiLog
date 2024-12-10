package shopify

import (
	"context"
	"time"
)

type Service interface {
	SyncOrders(ctx context.Context, shopName string, sinceId string, limit int, token string) error
	GetOrdersForShopAndAccount(ctx context.Context, shopName string, accountId string) ([]Order, error)
	UpdateOrder(ctx context.Context, order Order, accountId string, shopName string) error
	StoreToken(ctx context.Context, shopName string, accountId string, token string) error
	PutOrder(ctx context.Context, order Order) error
}

type Order struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ShopName   string    `json:"shop_name"`
	AccountId  string    `json:"account_id"`
	OrderId    string    `json:"order_id"`
	TotalPrice float64   `json:"total_price"`
}

type shopifyService struct {
	repo Repository
}

func NewShopifyService(repo Repository) Service {
	return &shopifyService{repo}
}

func (s *shopifyService) PutOrder(ctx context.Context, order Order) error {

	return s.repo.PutOrder(ctx, order)

}

func (s *shopifyService) SyncOrders(ctx context.Context, shopName string, sinceId string, limit int, token string) error {

	_, err := s.repo.SyncOrders(ctx, shopName, sinceId, limit, token)
	if err != nil {
		return err
	}
	return nil

}

func (s *shopifyService) GetOrdersForShopAndAccount(ctx context.Context, shopName string, accountId string) ([]Order, error) {

	return s.repo.GetOrdersForShopAndAccount(ctx, shopName, accountId)

}

func (s *shopifyService) UpdateOrder(ctx context.Context, order Order, accountId string, shopName string) error {

	return s.repo.UpdateOrder(ctx, order, accountId, shopName)

}

func (s *shopifyService) StoreToken(ctx context.Context, shopName string, accountId string, token string) error {

	return s.repo.StoreToken(ctx, shopName, accountId, token)

}
