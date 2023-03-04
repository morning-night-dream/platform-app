package article_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/e2e/helper"
	articlev1 "github.com/morning-night-dream/platform-app/pkg/connect/article/v1"
	authv1 "github.com/morning-night-dream/platform-app/pkg/connect/auth/v1"
)

func TestE2EArticleRead(t *testing.T) {
	t.Parallel()

	size := uint32(10)

	url := helper.GetCoreEndpoint(t)

	t.Run("記事が既読できる", func(t *testing.T) {
		t.Parallel()

		dsn := helper.GetDSN(t)

		adb := helper.NewArticleDB(t, dsn)
		defer adb.Close()

		ids := helper.GenerateIDs(t, 10)

		adb.BulkInsert(ids)

		user := helper.NewCoreUser(t, url)

		defer func() {
			user.Delete(t)
		}()

		ac := helper.NewConnectClient(t, http.DefaultClient, url)

		sreq := &authv1.SignInRequest{
			Email:    user.EMail,
			Password: user.Password,
		}

		sres, _ := ac.Auth.SignIn(context.Background(), connect.NewRequest(sreq))

		hc := &http.Client{
			Transport: helper.NewCookieTransport(t, sres.Header().Get("Set-Cookie")),
		}

		client := helper.NewConnectClient(t, hc, url)

		articles, err := client.Article.List(context.Background(), connect.NewRequest(&articlev1.ListRequest{
			MaxPageSize: size,
		}))
		if err != nil {
			t.Fatalf("failed to article share: %s", err)
		}

		req := &articlev1.ReadRequest{
			Id: articles.Msg.Articles[0].Id,
		}

		_, err = client.Article.Read(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to article share: %s", err)
		}
	})
}
