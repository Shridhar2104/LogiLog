package account

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Shridhar2104/logilo/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// grpcServer implements the gRPC server and embeds UnimplementedAccountServiceServer
// to ensure forward compatibility with newer versions of the gRPC service definitions.
type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

// NewGRPCServer initializes and starts the gRPC server.
// It takes a Service instance and the server's listening port as parameters.
func NewGRPCServer(service Service, port int) error {
	// Create a TCP listener at the specified port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", port, err)
	}

	// Initialize a new gRPC server instance
	server := grpc.NewServer()

	// Register the gRPC server with the service implementation
	pb.RegisterAccountServiceServer(server, &grpcServer{service: service})

	// Enable gRPC server reflection to make debugging easier
	reflection.Register(server)

	// Start serving gRPC requests
	log.Printf("gRPC server listening on port %d", port)
	return server.Serve(lis)
}

// CreateAccount handles the creation of a new account via gRPC.
// It validates the request, calls the business logic, and sends back a response.
func (s *grpcServer) CreateAccount(ctx context.Context, r *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	// Call the service layer's CreateAccount method
	a, err := s.service.CreateAccount(ctx, r.Name, r.Password, r.Email)
	if err != nil {
		log.Printf("Failed to create account: %v", err)
		// Return a wrapped error with additional context
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	// Construct and return the gRPC response
	return &pb.CreateAccountResponse{
		Account: &pb.Account{
			Id:   a.ID.String(),
			Name: a.Name,
		},
	}, nil
}

// GetAccountByID fetches account details based on the provided ID.
// It queries the service layer and sends the account details back via gRPC.
func (s *grpcServer) GetAccountByEmailAndPassword(ctx context.Context, r *pb.GetAccountByEmailAndPasswordRequest) (*pb.GetAccountByEmailAndPasswordResponse, error) {
	// Query the service layer for account details by email
	a, err := s.service.LoginAccount(ctx, r.Email, r.Password)
	if err != nil {
		log.Printf("Error while authenticating account: %v", err)
		return nil, fmt.Errorf("error while authenticating account: %w", err)
	}

	// Return the account data as part of the gRPC response
	return &pb.GetAccountByEmailAndPasswordResponse{
		Account: &pb.Account{
			Id:   a.ID.String(),
			Name: a.Name,
			Email:a.Email,
		},
	}, nil
}


// ListAccounts retrieves a paginated list of accounts from the database.
// The request contains skip/offset and take/limit parameters for pagination.
func (s *grpcServer) ListAccounts(ctx context.Context, r *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	// Query the service layer for a paginated list of accounts
	accounts, err := s.service.ListAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		log.Printf("Error while listing accounts: %v", err)
		return nil, fmt.Errorf("error while listing accounts: %w", err)
	}

	// Map database account models to gRPC Account representations
	grpcAccounts := []*pb.Account{}
	for _, account := range accounts {
		grpcAccounts = append(grpcAccounts, &pb.Account{
			Id:   account.ID.String(),
			Name: account.Name,
		})
	}

	// Return the mapped list of accounts in the response
	return &pb.ListAccountsResponse{
		Accounts: grpcAccounts,
	}, nil
}
