package auth_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/e2e/helper"
	authv1 "github.com/morning-night-dream/platform-app/pkg/connect/proto/auth/v1"
)

func TestE2EAuthScenario(t *testing.T) {
	t.Parallel()

	url := helper.GetCoreEndpoint(t)

	t.Run("サインアップ~サインイン~パスワード変更~サインアウト~サインイン~ユーザー削除ができる", func(t *testing.T) {
		t.Parallel()

		// サインアップ

		client := helper.NewConnectClient(t, http.DefaultClient, url)

		email := fmt.Sprintf("%s@example.com", uuid.NewString())

		password := uuid.NewString()

		sureq := &authv1.SignUpRequest{
			Email:    email,
			Password: password,
		}

		if _, err := client.Auth.SignUp(context.Background(), connect.NewRequest(sureq)); err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		// サインイン

		sireq := &authv1.SignInRequest{
			Email:    email,
			Password: password,
		}

		sires, err := client.Auth.SignIn(context.Background(), connect.NewRequest(sireq))
		if err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		if reflect.DeepEqual(len(sires.Header().Get("Set-Cookie")), 0) {
			t.Errorf("cookie is empty")
		}

		client = helper.NewConnectClientWithCookie(t, sires.Header().Get("Set-Cookie"), url)

		// パスワード変更

		np := uuid.NewString()

		cpreq := &authv1.ChangePasswordRequest{
			Email:       email,
			OldPassword: password,
			NewPassword: np,
		}

		if _, err := client.Auth.ChangePassword(context.Background(), connect.NewRequest(cpreq)); err != nil {
			t.Fatalf("failed to auth change password: %s", err)
		}

		// サインアウト

		soreq := &authv1.SignOutRequest{}

		sores, err := client.Auth.SignOut(context.Background(), connect.NewRequest(soreq))
		if err != nil {
			t.Fatalf("failed to auth sign out: %s", err)
		}

		if !reflect.DeepEqual(sores.Header().Get("Set-Cookie"), "token=; Max-Age=0") {
			t.Errorf("cookie = %v, want %v", sores.Header().Get("Set-Cookie"), "token=; Max-Age=0")
		}

		// 再ログイン

		client = helper.NewConnectClient(t, http.DefaultClient, url)

		sireq = &authv1.SignInRequest{
			Email:    email,
			Password: np,
		}

		sires, err = client.Auth.SignIn(context.Background(), connect.NewRequest(sireq))
		if err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		if reflect.DeepEqual(len(sires.Header().Get("Set-Cookie")), 0) {
			t.Errorf("cookie is empty")
		}

		// ユーザー削除

		client = helper.NewConnectClientWithCookie(t, sires.Header().Get("Set-Cookie"), url)

		dreq := &authv1.DeleteRequest{
			Email:    email,
			Password: np,
		}

		if _, err := client.Auth.Delete(context.Background(), connect.NewRequest(dreq)); err != nil {
			t.Errorf("failed to auth delete: %s", err)
		}
	})
}
