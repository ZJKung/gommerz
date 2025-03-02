package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/zjkung/gommerz/account"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
	GRPCPort    int    `envconfig:"GRPC_PORT" default:"8080"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port", cfg.GRPCPort)
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, cfg.GRPCPort))

}
