package server

import (
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/driver/config"
	"github.com/morning-night-dream/platform-app/internal/driver/env"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Router struct {
	newrelic *newrelic.Application
	routes   []Route
}

type Route struct {
	path    string
	handler http.Handler
}

func NewRoute(path string, handler http.Handler) Route {
	return Route{
		path:    path,
		handler: handler,
	}
}

func NewRouter(routes ...Route) *Router {
	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName(config.Core.NewRelicAppName),
		newrelic.ConfigLicense(config.Core.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
		func(c *newrelic.Config) {
			c.Logger = nrzap.Transform(log.Log())
		},
	)

	return &Router{
		newrelic: app,
		routes:   routes,
	}
}

func (r Router) Mux() *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range r.routes {
		path := route.path
		handler := route.handler

		if env.Env.IsProd() {
			path, handler = newrelic.WrapHandle(r.newrelic, route.path, route.handler)
		}

		mux.Handle(path, handler)
	}

	return mux
}
