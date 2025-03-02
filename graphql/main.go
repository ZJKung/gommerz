package main

import (
	"log"
	"net/http"

	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CategoryUrl string `envconfig:"CATALOG_SERVICE_URL"`
	OrderUrl    string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}
	s, err := NewGraphQLServer(cfg.AccountUrl, cfg.CategoryUrl, cfg.OrderUrl)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", handler.Playground("akhil", "/graphql"))
	// http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
	// http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Println("connect to http://localhost:8080/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
