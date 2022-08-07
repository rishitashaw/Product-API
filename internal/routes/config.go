package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/cors"
)

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Cors(next http.Handler) http.Handler {
	// cors
	return cors.New(cors.Options{
		AllowOrigin: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		ExposeHeaders: []string{"*"},
		AllowCredentials: true,
		MaxAge: 5,
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSecond int) *Config{
	// timeout
	c.timeout=time.Duration(timeInSecond)*time.Second
	return c
}

func (c *Config) GetTimeout() time.Duration {
	// get timeout
	return c.timeout
}