package server

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/morning-night-dream/platform-app/internal/adapter/controller"
	"github.com/morning-night-dream/platform-app/pkg/connect/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/health/v1/healthv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/user/v1/userv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/version/v1/versionv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewConnectHandler(
	health *controller.Health,
	user *controller.User,
	article *controller.Article,
	version *controller.Version,
) http.Handler {
	interceptor := connect.WithInterceptors(NewInterceptor())

	mux := NewRouter(
		NewRoute(healthv1connect.NewHealthServiceHandler(health, interceptor)),
		NewRoute(articlev1connect.NewArticleServiceHandler(article, interceptor)),
		NewRoute(userv1connect.NewUserServiceHandler(user, interceptor)),
		NewRoute(versionv1connect.NewVersionServiceHandler(version, interceptor)),
	).Mux()

	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker()))

	return h2c.NewHandler(mux, &http2.Server{})
}
