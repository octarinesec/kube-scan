package main

import (
	"flag"
	"github.com/kelseyhightower/envconfig"
	"kube-scan/scanner"
	"log"
)

const (
	AppName = "KUBESCAN"
)

type ServiceConfig struct {
	Port                        int    `default:"8080"`
	RiskConfigFilePath          string `split_words:"true" required:"true"`
	RefreshStateIntervalMinutes int    `split_words:"true" default:"1440"` // 1 day
}

func init() {
	flag.Parse() // This is here for glog
}

func main() {
	var cfg ServiceConfig
	if err := envconfig.Process(AppName, &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := scanner.InitScanner(cfg.RefreshStateIntervalMinutes, cfg.RiskConfigFilePath); err != nil {
		log.Fatalf("Failed to initialize scanner: %v", err)
	}

	if err := scanner.InitApi(cfg.Port); err != nil {
		log.Fatalf("Failed to initialize api: %v", err)
	}
}
