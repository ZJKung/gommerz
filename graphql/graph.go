package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/zjkung/gommerz/account"
)

type Server struct {
	accountClient *account.Client
	// categoryClient *category.Client
	// orderClient     *order.Client
}

func NewGraphQLServer(accountUrl, categoryUrl, productUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}
	categoryClient, err := category.NewCategoryClient(categoryUrl)
	if err != nil {
		accountClient.Close()
		return nil, err
	}
	orderClient, err := order.NewProductClient(productUrl)
	if err != nil {
		accountClient.Close()
		categoryClient.Close()
		return nil, err
	}
	return &Server{
		accountClient,
		categoryClient,
		orderClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return MutationResolverImp{s}
}

func (s *Server) Query() QueryResolver {
	return QueryResolverImp{s}
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{s}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
