package auth_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/e2e/helper"
	authv1 "github.com/morning-night-dream/platform-app/pkg/connect/auth/v1"
)

func TestE2EAuthSignOut(t *testing.T) {
	t.Parallel()

	url := helper.GetCoreEndpoint(t)

	t.Run("サインアウトできる", func(t *testing.T) {
		t.Parallel()

		user := helper.NewUser(t, url)

		defer func() {
			user.Delete(t)
		}()

		hc := helper.NewConnectClient(t, http.DefaultClient, url)

		sreq := &authv1.SignInRequest{
			Email:    user.EMail,
			Password: user.Password,
		}

		sres, err := hc.Auth.SignIn(context.Background(), connect.NewRequest(sreq))
		if err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		c := &http.Client{
			Transport: helper.NewCookieTransport(t, sres.Header().Get("Set-Cookie")),
		}

		client := helper.NewConnectClient(t, c, url)

		req := &authv1.SignOutRequest{}

		res, err := client.Auth.SignOut(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to auth sign out: %s", err)
		}

		if !reflect.DeepEqual(res.Header().Get("Set-Cookie"), "token=; Max-Age=0") {
			t.Errorf("cookie = %v, want %v", res.Header().Get("Set-Cookie"), "token=; Max-Age=0")
		}
	})
}
