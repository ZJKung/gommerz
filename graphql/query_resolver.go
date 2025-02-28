package main

import (
	"context"

	"github.com/zjkung/gommerz/account"
	"github.com/zjkung/gommerz/catalog"
)

type QueryResolverImp struct {
	server *Server
}

func (r *QueryResolverImp) Accounts(ctx context.Context, paganation *PaginationInput, id *string) ([]*account.Account, error) {
	return r.server.accountClient.GetAccounts(ctx, uint64(*paganation.Skip), uint64(*paganation.Take))
}

func (r *QueryResolverImp) Products(ctx context.Context, paganation *PaginationInput, id *string) ([]*catalog.Product, error) {
	return r.server.catalogClient.GetProducts(ctx, "", nil, uint64(*paganation.Skip), uint64(*paganation.Take))
}
