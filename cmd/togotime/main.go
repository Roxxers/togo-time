package main

import (
	"fmt"
	"log"
	"togotime"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	APIToken string `env:"API_TOKEN,required"`
}

func main() {
	// Load config into env, then load those variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	h, err := togotime.NewAPIClient(cfg.APIToken)
	if err != nil {
		fmt.Println(fmt.Errorf("%s", err))
	} else {
		fmt.Printf("%+v\n", h.Workspaces[0].Projects)
	}
}
