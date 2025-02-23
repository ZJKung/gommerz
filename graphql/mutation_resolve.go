package main

import (
	"context"

	"github.com/zjkung/gommerz/account"
)

type MutationResolverImp struct {
	server *Server
}

func (r *MutationResolverImp) CreateAccount(ctx context.Context, input AccountInput) (*account.Account, error) {
	return r.server.accountClient.PostAccount(ctx, input.Name)
}

func (r *MutationResolverImp) CreateCategory(ctx context.Context, input CategoryInput) (*Category, error) {
	return r.server.categoryClient.CreateCategory(ctx, input)
}
func (r *MutationResolverImp) CreateOrder(ctx context.Context, input OrderInput) (*Order, error) {
	return r.server.orderClient.CreateOrder(ctx, input)
}
