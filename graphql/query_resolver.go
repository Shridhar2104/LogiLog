package main

import (
	"context"
	"github.com/Shridhar2104/logilo/graphql/models"
	
)


type queryResolver struct {
	server *Server
}


func (r *queryResolver) Accounts(ctx context.Context, pagination PaginationInput) ([]*models.Account, error) {
	res, err := r.server.accountClient.ListAccounts(ctx, uint64(pagination.Skip), uint64(pagination.Take))
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.Account, len(res))
	for i, account := range res {
		accounts[i] = &models.Account{ID: account.ID, Name: account.Name}
	}
	return accounts, nil
}