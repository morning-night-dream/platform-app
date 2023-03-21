package main

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/morning-night-dream/platform-app/internal/adapter/handler"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/driver/client"
	"github.com/morning-night-dream/platform-app/internal/driver/config"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/public"
	"github.com/morning-night-dream/platform-app/internal/driver/redis"
	"github.com/morning-night-dream/platform-app/internal/driver/server"
	"github.com/morning-night-dream/platform-app/internal/usecase/interactor"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

var version string

func main() {
	c, err := client.New().Of(config.API.AppCoreURL)
	if err != nil {
		panic(err)
	}

	authRPC, err := firebase.New().Of(config.Core.FirebaseSecret, config.Core.FirebaseAPIEndpoint, config.Core.FirebaseAPIKey)
	if err != nil {
		panic(err)
	}

	rds := redis.NewRedis(config.Core.RedisURL)

	authCache, err := redis.New[model.Auth]().Of(rds)
	if err != nil {
		panic(err)
	}

	sessionCache, err := redis.New[model.Session]().Of(rds)
	if err != nil {
		panic(err)
	}

	codeCache, err := redis.New[model.Code]().Of(rds)
	if err != nil {
		panic(err)
	}

	auth := handler.NewAuth(
		interactor.NewAPIAuthSignIn(authRPC, authCache, sessionCache),
		interactor.NewAPIAuthSignOut(authCache, sessionCache),
		interactor.NewAPIAuthSignUp(authRPC),
		interactor.NewAPIAuthVerify(authCache),
		interactor.NewAPIAuthRefresh(sessionCache, codeCache),
		interactor.NewAPIAuthGenerateCode(codeCache),
	)

	hdl := handler.New(version, config.API.APIKey, auth, c, public.New())

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(config.API.CorsAllowOrigins, ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	handler := openapi.HandlerWithOptions(hdl, openapi.ChiServerOptions{
		BaseURL:     "/api",
		BaseRouter:  router,
		Middlewares: []openapi.MiddlewareFunc{server.Middleware},
	})

	srv := server.NewHTTPServer(handler)

	srv.Run()
}
