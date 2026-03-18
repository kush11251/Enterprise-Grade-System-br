package main

import (
	"example.com/enterprise-grade-system/pkg/config"
	"example.com/enterprise-grade-system/pkg/http"
	"example.com/enterprise-grade-system/pkg/logger"
	"flag"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		logger.Fatal(err)
	}
		http.StartServer(cfg)
}