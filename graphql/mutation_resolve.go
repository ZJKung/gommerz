package main

import (
	"context"

	"github.com/zjkung/gommerz/account"
	"github.com/zjkung/gommerz/catalog"
)

type MutationResolverImp struct {
	server *Server
}

func (r *MutationResolverImp) CreateAccount(ctx context.Context, input AccountInput) (*account.Account, error) {
	return r.server.accountClient.PostAccount(ctx, input.Name)
}

func (r *MutationResolverImp) CreateCategory(ctx context.Context, input ProductInput) (*catalog.Product, error) {
	return r.server.catalogClient.PostProduct(ctx, input.Name, input.Description, input.Price)
}
func (r *MutationResolverImp) CreateOrder(ctx context.Context, input OrderInput) (*Order, error) {
	return r.server.orderClient.CreateOrder(ctx, input)
}
