package main

import (
	"context"
	"github.com/Shridhar2104/logilo/graphql/models"
	"log"
	
)


type queryResolver struct {
	server *Server
}


func (r *queryResolver) GetAccountByID(ctx context.Context, email string, password string) (*models.Account, error) {
	// Fetch account data from the account service
	accountResp, err := r.server.accountClient.LoginAndGetAccount(ctx, email, password)
	if err != nil {
		log.Printf("Error fetching account by ID: %v", err)
		return nil, err
	}

	// Map account service response to the GraphQL model
	return &models.Account{
		ID:        accountResp.ID.String(),
		Name:      accountResp.Name,
		Password:  accountResp.Password,
		Email:     accountResp.Email,
	}, nil
}


func (r *queryResolver) Accounts(ctx context.Context, pagination PaginationInput) ([]*models.Account, error) {
	res, err := r.server.accountClient.ListAccounts(ctx, uint64(pagination.Skip), uint64(pagination.Take))
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.Account, len(res))
	for i, account := range res {
		accounts[i] = &models.Account{ID: account.ID.String(), Name: account.Name}
	}
	return accounts, nil
}