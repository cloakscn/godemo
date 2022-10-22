package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	year, m, day := time.Time.Date(time.Now())
	cfg.OutputPaths = []string {
		fmt.Sprintf("./log/%v-%v-%v.log", year, m, day),
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
  
	url := "http://example.org/api"
	logger.Info("failed to fetch URL",
	  zap.String("url", url),
	  zap.Int("attempt", 3),
	  zap.Duration("backoff", time.Second),
	)
  
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
	  "url", url,
	  "attempt", 3,
	  "backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
  }