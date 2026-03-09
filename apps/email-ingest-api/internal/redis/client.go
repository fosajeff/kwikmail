package redisclient

import (
	"context"
	"crypto/tls"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func New() *redis.Client {

	redisURL := os.Getenv("REDIS_URL")
	password := os.Getenv("REDIS_PASSWORD")
	useTLS := os.Getenv("REDIS_TLS")

	var opt *redis.Options

	if strings.HasPrefix(redisURL, "redis://") || strings.HasPrefix(redisURL, "rediss://") {
		var err error
		opt, err = redis.ParseURL(redisURL)
		if err != nil {
			panic("invalid REDIS_URL: " + err.Error())
		}
		if password != "" && opt.Password == "" {
			opt.Password = password
		}
	} else {
		addr := redisURL
		if addr == "" {
			addr = "localhost:6379"
		}
		opt = &redis.Options{
			Addr:     addr,
			Password: password,
		}

		if strings.ToLower(useTLS) == "true" {
			opt.TLSConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}
	}

	return redis.NewClient(opt)
}
