package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/zjkung/gommerz/order"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
	AccountURL  string `envconfig:"ACCOUNT_URL" required:"true"`
	CatalogURL  string `envconfig:"CATALOG_URL" required:"true"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r order.Repository
	retry.ForeverSleep(2, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println(("listening on :8080"))
	s := order.NewService(r)
	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, ":8080"))
}
