package main

import "context"

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context, acct *Account) ([]*Order, error) {
	return r.server.orderClient.Orders(ctx)
}
