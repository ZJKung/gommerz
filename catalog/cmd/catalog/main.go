package main

import (
	"log"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/zjkung/gommerz/catalog"
)

type Config struct {
	DatabaseUrl string `envconfig:"DATABASE_URL" required:"true"`
	Port        int    `envconfig:"PORT" default:"8080"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		r, err = catalog.NewElasticRepository(cfg.DatabaseUrl)
		if err != nil {
			log.Fatal(err)
		}
		return err
	})
	defer r.Close()
	log.Println("Listening on port " + strconv.Itoa(cfg.Port))
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, cfg.Port))
}
