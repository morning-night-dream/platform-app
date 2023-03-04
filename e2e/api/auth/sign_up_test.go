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

		// Userのライフサイクルも未定のため削除は未実施
	})
}
