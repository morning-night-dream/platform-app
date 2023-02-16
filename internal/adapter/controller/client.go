package controller

import (
	"github.com/morning-night-dream/platform-app/pkg/connect/proto/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/proto/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/proto/health/v1/healthv1connect"
)

type ClientFactory interface {
	Of(string) (*Client, error)
}

type Client struct {
	Article articlev1connect.ArticleServiceClient
	Health  healthv1connect.HealthServiceClient
	Auth    authv1connect.AuthServiceClient
}
