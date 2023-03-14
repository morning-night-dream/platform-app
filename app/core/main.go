package main

import (
	"github.com/morning-night-dream/platform-app/internal/adapter/controller"
	"github.com/morning-night-dream/platform-app/internal/adapter/gateway"
	"github.com/morning-night-dream/platform-app/internal/driver/config"
	"github.com/morning-night-dream/platform-app/internal/driver/database"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/redis"
	"github.com/morning-night-dream/platform-app/internal/driver/server"
)

var version string

func main() {
	db := database.NewClient(config.Core.DSN)

	cache := redis.NewClient(config.Core.RedisURL)

	da := gateway.NewArticle(db)

	fb := firebase.NewClient(config.Core.FirebaseSecret, config.Core.FirebaseAPIEndpoint, config.Core.FirebaseAPIKey)

	handle := controller.NewHandle(fb, cache)

	ah := controller.NewArticle(da, handle)

	hh := controller.NewHealth()

	auh := controller.NewAuth(handle)

	vh := controller.NewVersion(version)

	ch := server.NewConnectHandler(hh, ah, auh, vh)

	srv := server.NewHTTPServer(ch)

	srv.Run()
}