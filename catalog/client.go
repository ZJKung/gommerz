package catalog

import (
	"context"
	"fmt"

	"github.com/zjkung/gommerz/catalog/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.CatalogServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:    conn,
		service: pb.NewCatalogServiceClient(conn),
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	r, err := c.service.PostProduct(ctx, &pb.PostProductRequest{Name: name, Description: description, Price: price})
	if err != nil {
		return nil, err
	}
	return &Product{ID: r.Product.Id, Name: r.Product.Name, Description: r.Product.Description, Price: fmt.Sprintf("%f", r.Product.Price)}, nil
}

func (c *Client) GetProduct(ctx context.Context, id string) (*Product, error) {
	r, err := c.service.GetProduct(ctx, &pb.GetProductRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &Product{ID: r.Product.Id, Name: r.Product.Name, Description: r.Product.Description, Price: fmt.Sprintf("%f", r.Product.Price)}, nil
}

func (c *Client) GetProducts(ctx context.Context, query string, ids []string, skip, take uint64) ([]*Product, error) {
	var r *pb.GetProductsResponse
	var err error
	if query != "" {
		r, err = c.service.GetProducts(ctx, &pb.GetProductsRequest{Query: query})
	} else if len(ids) > 0 {
		r, err = c.service.GetProducts(ctx, &pb.GetProductsRequest{Ids: ids})
	} else {
		r, err = c.service.GetProducts(ctx, &pb.GetProductsRequest{Skip: skip, Take: take})
	}
	if err != nil {
		return nil, err
	}
	products := make([]*Product, 0)
	for _, p := range r.Products {
		products = append(products, &Product{ID: p.Id, Name: p.Name, Description: p.Description, Price: fmt.Sprintf("%f", p.Price)})
	}
	return products, nil
}
