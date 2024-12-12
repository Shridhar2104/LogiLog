package account

import (
	"context"
	"log"
	"time"

	"github.com/Shridhar2104/logilo/account/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}


	c:= pb.NewAccountServiceClient(conn)
	return &Client{conn: conn, service: c}, nil


}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) CreateAccount(ctx context.Context, a *Account) (*Account, error) {
	res, err := c.service.CreateAccount(ctx, &pb.CreateAccountRequest{
		Name: a.Name,
	})
	
	if err != nil {
		log.Printf("Error creating account: %v", err)
		return nil, err
	}
	return &Account{
		ID:   uuid.MustParse(res.Account.Id),
		Name: res.Account.Name,
		Password: res.Account.Password,
		Email: res.Account.Email,
		ShopName: res.Account.ShopName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}


func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	res, err := c.service.GetAccountByID(ctx, &pb.GetAccountByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:   uuid.MustParse(res.Account.Id),
		Name: res.Account.Name,
	}, nil
}


func (c *Client) ListAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {
	res, err := c.service.ListAccounts(ctx, &pb.ListAccountsRequest{Skip: skip, Take: take})
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, len(res.Accounts))
	for i, a := range res.Accounts {
		accounts[i] = Account{ID: uuid.MustParse(a.Id), Name: a.Name}
	}
	return accounts, nil
}