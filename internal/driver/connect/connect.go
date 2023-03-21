package connect

import (
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/adapter/external"
	"github.com/morning-night-dream/platform-app/internal/adapter/handler"
	"github.com/morning-night-dream/platform-app/pkg/connect/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/health/v1/healthv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/user/v1/userv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/version/v1/versionv1connect"
)

var _ handler.ClientFactory = (*Client)(nil)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Of(url string) (*handler.Client, error) {
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

	return &handler.Client{
		Article: ac,
		Health:  hc,
		Auth:    auc,
		Version: vc,
	}, nil
}

var _ external.UserFactory = (*Connect)(nil)

type Connect struct {
	client *http.Client
}

func New() *Connect {
	return &Connect{
		client: http.DefaultClient,
	}
}

func (con *Connect) User(url string) (*external.User, error) {
	uc := userv1connect.NewUserServiceClient(
		con.client,
		url,
	)

	return external.NewUser(uc), nil
}
