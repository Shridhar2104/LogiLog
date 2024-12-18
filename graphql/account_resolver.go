package main

import (
	"context"
	"github.com/Shridhar2104/logilo/graphql/models"
)

// Resolver implementation
type accountResolver struct {
	server *Server
}

// Orders fetch orders for all shopnames for the given account.
func (r *accountResolver) Orders(ctx context.Context, obj *models.Account) ([]*Order, error) {
	// Handle if there are no shops in the account
	if len(obj.ShopNames) == 0 {
		return nil, nil
	}

	// Fetch orders for each ShopName
	var allOrders []*Order
	for _, shopName := range obj.ShopNames {
		// Fetch orders for each shop/account pair
		res, err := r.server.shopifyClient.GetOrdersForShopAndAccount(ctx, shopName.Shopname, obj.ID)
		if err != nil {
			// Log error and return it
			return nil, err
		}

		// Map each Shopify response to our Order model
		for _, order := range res {
			allOrders = append(allOrders, &Order{
				ID:          order.ID,
				Amount:      order.TotalPrice,
				AccountID:    obj.ID,
			})
		}
	}

	return allOrders, nil
}

// Shopnames returns the shopnames related to the account
func (r *accountResolver) Shopnames(ctx context.Context, obj *models.Account) ([]*ShopName, error) {
	// Simply return the shop names in models.Account
	var shopNames []*ShopName
	for _, shopName := range obj.ShopNames {
		shopNames = append(shopNames, &ShopName{
			Shopname: shopName.Shopname,
		})
	}
	return shopNames, nil
}
