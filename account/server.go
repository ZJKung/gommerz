package account

import (
	"context"
	"fmt"
	"net"

	"github.com/zjkung/gommerz/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func NewService(r Repository) Service {
	return &accountService{r}
}
func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	reflection.Register(serv)
	return serv.Serve(lis)
}
func (s *grpcServer) PostAccount(ctx context.Context, req *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	account, err := s.service.PostAccount(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.service.GetAccount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, req *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	accounts, err := s.service.GetAccounts(ctx, req.Skip, req.Take)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetAccountsResponse{Accounts: make([]*pb.Account, 0)}
	for _, account := range accounts {
		resp.Accounts = append(resp.Accounts, &pb.Account{
			Id:   account.ID,
			Name: account.Name,
		})
	}

	return resp, nil
}
