package catalog

import (
	"context"
	"fmt"
	"net"

	"github.com/zjkung/gommerz/catalog/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterCatalogServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostProduct(ctx context.Context, req *pb.PostProductRequest) (*pb.PostProductResponse, error) {
	p, err := s.service.CreateProduct(ctx, req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.PostProductResponse{Product: &pb.Product{Id: p.ID, Name: p.Name, Description: p.Description, Price: p.Price}}, nil
}

func (s *grpcServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	p, err := s.service.GetProductById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductResponse{Product: &pb.Product{Id: p.ID, Name: p.Name, Description: p.Description, Price: p.Price}}, nil
}
func (s *grpcServer) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {

	var products []Product
	var err error
	if req.Query != "" {
		products, err = s.service.SearchProduct(ctx, req.Query, req.Skip, req.Take)
	} else if len(req.Ids) > 0 {
		products, err = s.service.ListProductsWithIDs(ctx, req.Ids)
	} else {
		products, err = s.service.ListProducts(ctx, req.Skip, req.Take)
	}
	if err != nil {
		return nil, err
	}
	resp := &pb.GetProductsResponse{Products: make([]*pb.Product, 0)}
	for _, p := range products {
		resp.Products = append(resp.Products, &pb.Product{Id: p.ID, Name: p.Name, Description: p.Description, Price: p.Price})
	}
	return resp, nil
}
