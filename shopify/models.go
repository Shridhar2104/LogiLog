package shopify
import "time"

type Order struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ShopName   string    `json:"shop_name"`
	AccountId  string    `json:"account_id"`
	OrderId    string    `json:"order_id"`
	TotalPrice float64   `json:"total_price"`
}