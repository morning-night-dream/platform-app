package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/e2e/helper"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

func TestE2EAuthSighUp(t *testing.T) {
	t.Parallel()

	url := helper.GetAPIEndpoint(t)

	t.Run("サインアップできる", func(t *testing.T) {
		t.Parallel()

		client := helper.NewOpenAPIClient(t, url)

		id := uuid.New().String()
		email := fmt.Sprintf("%s@example.com", id)
		password := id

		if _, err := client.Client.V1AuthSignUp(context.Background(), openapi.V1AuthSignUpJSONRequestBody{
			Email:    types.Email(email),
			Password: password,
		}); err != nil {
			t.Fatalf("failed to auth sign up: %s", err)
		}

		// defer func() {
		// 	res, err := client.Client.V1AuthSignIn(context.Background(), openapi.V1AuthSignInJSONRequestBody{
		// 		Email:     types.Email(email),
		// 		Password:  password,
		// 		PublicKey: "",
		// 	})
		// 	if err != nil {
		// 		t.Fatalf("failed to auth sign in: %s", err)
		// 	}

		// 	client.Client.Client = &http.Client{
		// 		Transport: helper.NewCookiesTransport(t, res.Cookies()),
		// 	}

		// 	req := &authv1.DeleteRequest{
		// 		Email:    email,
		// 		Password: password,
		// 	}

		// 	if _, err := dclient.Auth.Delete(context.Background(), connect.NewRequest(req)); err != nil {
		// 		t.Fatalf("failed to delete user in: %s", err)
		// 	}
		// }()
	})
}
