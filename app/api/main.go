package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/morning-night-dream/platform-app/internal/adapter/api"
	"github.com/morning-night-dream/platform-app/internal/adapter/gateway"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/driver/client"
	"github.com/morning-night-dream/platform-app/internal/driver/config"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/public"
	"github.com/morning-night-dream/platform-app/internal/driver/redis"
	"github.com/morning-night-dream/platform-app/internal/driver/server"
	"github.com/morning-night-dream/platform-app/internal/driver/store"
	"github.com/morning-night-dream/platform-app/internal/driver/user"
	"github.com/morning-night-dream/platform-app/internal/usecase/interactor"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

var version string

func main() {
	c, err := client.New().Of(config.Gateway.AppCoreURL)
	if err != nil {
		panic(err)
	}

	fb := firebase.NewClient(config.Core.FirebaseSecret, config.Core.FirebaseAPIEndpoint, config.Core.FirebaseAPIKey)

	rds := redis.NewRedis(config.Core.RedisURL)

	authCache, err := redis.New[model.Auth]().Of(rds)
	if err != nil {
		panic(err)
	}

	sessionCache, err := redis.New[model.Session]().Of(rds)
	if err != nil {
		panic(err)
	}

	authRepo := gateway.NewAPIAuth(fb)

	codeRepo := gateway.NewAPICode()

	auth := api.NewAuth(
		interactor.NewAPIAuthSignIn(authRepo, authCache, sessionCache),
		interactor.NewAPIAuthSignOut(authCache, sessionCache),
		interactor.NewAPIAuthSignUp(authRepo),
		interactor.NewAPIAuthVerify(authCache),
		interactor.NewAPIAuthRefresh(sessionCache, codeRepo),
		interactor.NewAPIAuthGenerateCode(codeRepo),
	)

	ap := api.New(version, auth, c, store.New(), fb, public.New(), user.New())

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	handler := openapi.HandlerWithOptions(ap, openapi.ChiServerOptions{
		BaseURL:     "/api",
		BaseRouter:  router,
		Middlewares: []openapi.MiddlewareFunc{server.Middleware},
	})

	srv := server.NewHTTPServer(handler)

	srv.Run()
}
