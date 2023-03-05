package client

import (
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/adapter/api"
	"github.com/morning-night-dream/platform-app/pkg/connect/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/health/v1/healthv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/version/v1/versionv1connect"
)

var _ api.ClientFactory = (*Client)(nil)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Of(url string) (*api.Client, error) {
	client := http.DefaultClient

	ac := articlev1connect.NewArticleServiceClient(
		client,
		url,
	)

	hc := healthv1connect.NewHealthServiceClient(
		client,
		url,
	)

	auc := authv1connect.NewAuthServiceClient(
		client,
		url,
	)

	vc := versionv1connect.NewVersionServiceClient(
		client,
		url,
	)

	return &api.Client{
		Article: ac,
		Health:  hc,
		Auth:    auc,
		Version: vc,
	}, nil
}
