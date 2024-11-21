package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	User string `env:"USER"`
	/*	Files []string `env:"FILES" envSeparator:":"`
		Home  string   `env:"HOME"`
		// required требует, чтобы переменная TASK_DURATION была определена
		TaskDuration time.Duration `env:"TASK_DURATION,required"`*/
}

func main() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current user is %s\n", cfg.User)
	//log.Println(cfg)
}
