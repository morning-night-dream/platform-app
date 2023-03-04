package helper

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

type User struct {
	EMail    string
	Password string
	Cookies  []*http.Cookie
	Client   *OpenAPIClient
}

func NewUser(
	t *testing.T,
	url string,
) User {
	t.Helper()

	client := NewOpenAPIClient(t, url)

	id := uuid.New().String()
	email := fmt.Sprintf("%s@example.com", id)
	password := id

	if _, err := client.Client.V1AuthSignUp(context.Background(), openapi.V1AuthSignUpJSONRequestBody{
		Email:    types.Email(email),
		Password: password,
	}); err != nil {
		t.Fatalf("failed to auth sign up: %s", err)
	}

	res, err := client.Client.V1AuthSignIn(context.Background(), openapi.V1AuthSignInJSONRequestBody{
		Email:     types.Email(email),
		Password:  password,
		PublicKey: "",
	})
	if err != nil {
		t.Fatalf("failed to auth sign in: %s", err)
	}

	client.Client.Client = &http.Client{
		Transport: NewCookiesTransport(t, res.Cookies()),
	}

	return User{
		EMail:    email,
		Password: password,
		Cookies:  res.Cookies(),
		Client:   client,
	}
}
