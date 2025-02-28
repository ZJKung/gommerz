package catalog

import (
	"context"
	"fmt"

	"github.com/segmentio/ksuid"
)

type Service interface {
	CreateProduct(ctx context.Context, name, description string, price float64) (*Product, error)
	GetProductById(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProduct(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type catalogService struct {
	repo Repository
}

func NewService(repo Repository) *catalogService {
	return &catalogService{repo}
}

func (s *catalogService) CreateProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	product := &Product{
		ID:          ksuid.New().String(),
		Name:        name,
		Description: description,
		Price:       fmt.Sprintf("%.2f", price),
	}
	err := s.repo.PutProduct(ctx, *product)
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (s *catalogService) GetProductById(ctx context.Context, id string) (*Product, error) {
	product, err := s.repo.GetProductById(ctx, id)
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (s *catalogService) ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	products, err := s.repo.ListProducts(ctx, skip, take)
	if err != nil {
		return []Product{}, err
	}
	return products, nil
}

func (s *catalogService) ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
	products, err := s.repo.ListProductsWithIDs(ctx, ids)
	if err != nil {
		return []Product{}, err
	}
	return products, nil
}

func (s *catalogService) SearchProduct(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	products, err := s.repo.SearchProduct(ctx, query, skip, take)
	if err != nil {
		return []Product{}, err
	}
	return products, nil
}
