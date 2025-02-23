package main

import (
	"context"

	"github.com/zjkung/gommerz/account"
)

type QueryResolverImp struct {
	server *Server
}

func (r *QueryResolverImp) Accounts(ctx context.Context, paganation *PaginationInput, id *string) ([]*account.Account, error) {
	return r.server.accountClient.GetAccounts(ctx, uint64(*paganation.Skip), uint64(*paganation.Take))
}

func (r *QueryResolverImp) Products(ctx context.Context, paganation *PaginationInput, id *string) ([]*Product, error) {
	return r.server.productClient.Products(ctx)
}
