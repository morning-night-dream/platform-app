package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/morning-night-dream/platform-app/internal/adapter/controller"
	"github.com/morning-night-dream/platform-app/internal/driver/client"
	"github.com/morning-night-dream/platform-app/internal/driver/config"
	"github.com/morning-night-dream/platform-app/internal/driver/server"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

func main() {
	c, err := client.New().Of(config.Gateway.AppCoreURL)
	if err != nil {
		panic(err)
	}

	ctr := controller.New(c)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	handler := openapi.HandlerWithOptions(ctr, openapi.ChiServerOptions{
		BaseURL:     "/api",
		BaseRouter:  router,
		Middlewares: []openapi.MiddlewareFunc{server.Middleware},
	})

	srv := server.NewHTTPServer(handler)

	srv.Run()
}
