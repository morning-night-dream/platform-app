package main

import (
	"github.com/morning-night-dream/platform-app/internal/adapter/controller"
	"github.com/morning-night-dream/platform-app/internal/adapter/gateway"
	"github.com/morning-night-dream/platform-app/internal/driver/config"
	"github.com/morning-night-dream/platform-app/internal/driver/database"
	"github.com/morning-night-dream/platform-app/internal/driver/redis"
	"github.com/morning-night-dream/platform-app/internal/driver/server"
	"github.com/morning-night-dream/platform-app/internal/usecase/interactor"
)

var version string

func main() {
	db := database.NewClient(config.Core.DSN)

	cache := redis.NewClient(config.Core.RedisURL)

	articleRepo := gateway.NewArticle(db)

	userRepo := gateway.NewUser(db)

	userSignUp := interactor.NewCoreUserSignUp(userRepo)

	ctl := controller.New(cache)

	ah := controller.NewArticle(ctl, articleRepo)

	uh := controller.NewUser(ctl, userSignUp)

	hh := controller.NewHealth()

	vh := controller.NewVersion(version)

	ch := server.NewConnectHandler(hh, uh, ah, vh)

	srv := server.NewHTTPServer(ch)

	srv.Run()
}
