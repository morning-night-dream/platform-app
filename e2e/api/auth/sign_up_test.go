package auth_test

import (
	"context"
	"fmt"
	"net/http"
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

		client.Client.Client = &http.Client{
			Transport: helper.NewAPIKeyTransport(t, helper.GetAPIKey(t)),
		}

		id := uuid.New().String()

		email := fmt.Sprintf("%s@example.com", id)

		password := id

		res, err := client.Client.V1AuthSignUp(context.Background(), openapi.V1AuthSignUpJSONRequestBody{
			Email:    types.Email(email),
			Password: password,
		})
		if err != nil {
			t.Fatalf("failed to auth sign up: %s", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("failed to auth sign up: %d", res.StatusCode)
		}
	})

	t.Run("Api-Keyがなくサインアップできない", func(t *testing.T) {
		t.Parallel()

		client := helper.NewOpenAPIClient(t, url)

		id := uuid.New().String()

		email := fmt.Sprintf("%s@example.com", id)

		password := id

		res, err := client.Client.V1AuthSignUp(context.Background(), openapi.V1AuthSignUpJSONRequestBody{
			Email:    types.Email(email),
			Password: password,
		})
		if err != nil {
			t.Fatalf("failed to auth sign up: %s", err)
		}

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("failed to auth sign up: %d", res.StatusCode)
		}
	})
}
