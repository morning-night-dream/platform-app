package server

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/internal/adapter/handler"
	"github.com/morning-night-dream/platform-app/pkg/connect/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/health/v1/healthv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/version/v1/versionv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewConnectHandler(
	health *handler.Health,
	article *handler.Article,
	auth *handler.Auth,
	version *handler.Version,
) http.Handler {
	interceptor := connect.WithInterceptors(NewInterceptor())

	mux := NewRouter(
		NewRoute(healthv1connect.NewHealthServiceHandler(health, interceptor)),
		NewRoute(articlev1connect.NewArticleServiceHandler(article, interceptor)),
		NewRoute(authv1connect.NewAuthServiceHandler(auth, interceptor)),
		NewRoute(versionv1connect.NewVersionServiceHandler(version, interceptor)),
	).Mux()

	return h2c.NewHandler(mux, &http2.Server{})
}
