package auth_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
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

		defer func() {
			// secret
			priv, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				panic(err)
			}

			pub := priv.Public()

			log.Printf("pub key: %+v", pub)

			bytes, err := json.Marshal(pub)
			if err != nil {
				panic(err)
			}

			pubstr := base64.StdEncoding.EncodeToString(bytes)

			res, err := client.Client.V1AuthSignIn(context.Background(), openapi.V1AuthSignInJSONRequestBody{
				Email:     types.Email(email),
				Password:  password,
				PublicKey: pubstr,
			})
			if err != nil {
				t.Fatalf("failed to auth sign in: %s", err)
			}

			log.Printf("res: %+v", res)
		}()
	})
}
