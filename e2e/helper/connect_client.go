package helper

import (
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/pkg/connect/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform-app/pkg/connect/health/v1/healthv1connect"
)

type ConnectClient struct {
	Article articlev1connect.ArticleServiceClient
	Health  healthv1connect.HealthServiceClient
	Auth    authv1connect.AuthServiceClient
}

func NewConnectClient(t *testing.T, client connect.HTTPClient, url string) *ConnectClient {
	t.Helper()

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

	return &ConnectClient{
		Article: ac,
		Health:  hc,
		Auth:    auc,
	}
}

func NewPlainConnectClient(t *testing.T, url string) *ConnectClient {
	t.Helper()

	return NewConnectClient(t, http.DefaultClient, url)
}

func NewConnectClientWithAPIKey(t *testing.T, key string, url string) *ConnectClient {
	t.Helper()

	client := &http.Client{
		Transport: NewXAPIKeyTransport(t, key),
	}

	return NewConnectClient(t, client, url)
}

func NewConnectClientWithCookie(t *testing.T, cookie string, url string) *ConnectClient {
	t.Helper()

	client := &http.Client{
		Transport: NewCookieTransport(t, cookie),
	}

	return NewConnectClient(t, client, url)
}

type XAPIKeyTransport struct {
	t         *testing.T
	APIKey    string
	Transport http.RoundTripper
}

func NewXAPIKeyTransport(
	t *testing.T,
	key string,
) *XAPIKeyTransport {
	return &XAPIKeyTransport{
		t:         t,
		APIKey:    key,
		Transport: http.DefaultTransport,
	}
}

func (at *XAPIKeyTransport) transport() http.RoundTripper {
	at.t.Helper()

	return at.Transport
}

func (at *XAPIKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	at.t.Helper()

	req.Header.Add("X-API-KEY", at.APIKey)

	resp, err := at.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
