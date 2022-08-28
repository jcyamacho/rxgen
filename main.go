package main

import (
	"github.com/caarlos0/log"
	"github.com/jcyamacho/rxgen/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.WithError(err).Fatal("unexpected error")
	}
}
