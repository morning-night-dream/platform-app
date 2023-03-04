package helper

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	authv1 "github.com/morning-night-dream/platform-app/pkg/connect/auth/v1"
)

type CoreUser struct {
	EMail    string
	Password string
	Cookie   string
	Client   *ConnectClient
}

func NewCoreUser(
	t *testing.T,
	url string,
) CoreUser {
	t.Helper()

	ctx := context.Background()

	client := NewConnectClient(t, http.DefaultClient, url)

	email := fmt.Sprintf("%s@example.com", uuid.NewString())

	password := uuid.NewString()

	sureq := &authv1.SignUpRequest{
		Email:    email,
		Password: password,
	}

	if _, err := client.Auth.SignUp(ctx, connect.NewRequest(sureq)); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	sireq := &authv1.SignInRequest{
		Email:    email,
		Password: password,
	}

	res, err := client.Auth.SignIn(ctx, connect.NewRequest(sireq))
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	cookie := res.Header().Get("Set-Cookie")

	return CoreUser{
		EMail:    email,
		Password: password,
		Cookie:   cookie,
		Client:   NewConnectClientWithCookie(t, cookie, url),
	}
}

func (u CoreUser) ChangePassword(t *testing.T, password string) CoreUser {
	t.Helper()

	req := &authv1.ChangePasswordRequest{
		Email:       u.EMail,
		OldPassword: u.Password,
		NewPassword: password,
	}

	if _, err := u.Client.Auth.ChangePassword(context.Background(), connect.NewRequest(req)); err != nil {
		t.Fatalf("failed to change password: %v", err)
	}

	return CoreUser{
		EMail:    u.EMail,
		Password: password,
		Cookie:   u.Cookie,
		Client:   u.Client,
	}
}

func (u CoreUser) Delete(t *testing.T) {
	t.Helper()

	req := &authv1.DeleteRequest{
		Email:    u.EMail,
		Password: u.Password,
	}

	if _, err := u.Client.Auth.Delete(context.Background(), connect.NewRequest(req)); err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}
}
