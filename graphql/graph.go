package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/zjkung/gommerz/account"
	"github.com/zjkung/gommerz/catalog"
	"github.com/zjkung/gommerz/order"
)

type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NewGraphQLServer(accountUrl, categoryUrl, productUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}
	catalogClient, err := catalog.NewClient(categoryUrl)
	if err != nil {
		accountClient.Close()
		return nil, err
	}
	orderClient, err := order.NewClient(productUrl)
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return nil, err
	}
	return &Server{
		accountClient,
		catalogClient,
		orderClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{s}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{s}
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{s}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
