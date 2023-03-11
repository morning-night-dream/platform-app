package auth_test

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/e2e/helper"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

func TestE2EAuthRefresh(t *testing.T) {
	t.Parallel()

	url := helper.GetAPIEndpoint(t)

	t.Run("リフレッシュできる", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		client := helper.NewOpenAPIClient(t, url)

		id := uuid.New().String()

		email := fmt.Sprintf("%s@example.com", id)

		password := id

		if res, err := client.Client.V1AuthSignUp(context.Background(), openapi.V1AuthSignUpJSONRequestBody{
			Email:    types.Email(email),
			Password: password,
		}); err != nil || res.StatusCode != http.StatusOK {
			t.Fatalf("failed to auth sign up: %s", err)
		}

		prv, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatal(err)
		}

		res, err := client.Client.V1AuthSignIn(ctx, openapi.V1AuthSignInJSONRequestBody{
			Email:     types.Email(email),
			Password:  password,
			PublicKey: helper.Public(t, prv),
		})
		if err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("failed to auth sign in: %d", res.StatusCode)
		}

		client.Client.Client = &http.Client{
			Transport: helper.NewOnlySIDCookieTransport(t, res.Cookies()),
		}

		var sid *http.Cookie
		for _, c := range res.Cookies() {
			if c.Name == "SID" {
				sid = c
			}
		}

		res, err = client.Client.V1AuthVerify(ctx)
		if err != nil {
			t.Fatalf("failed to verify in: %s", err)
		}

		if res.StatusCode == http.StatusOK {
			t.Fatalf("success to verify in: %d", res.StatusCode)
		}

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("failed to verify in: %d", res.StatusCode)
		}

		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		var unauthorized openapi.V1UnauthorizedResponse
		if err := json.Unmarshal(body, &unauthorized); err != nil {
			t.Fatalf("failed marshal response: %s caused by %s", body, err)
			return
		}

		if _, err := uuid.Parse(unauthorized.Code.String()); err != nil {
			t.Errorf("failed to parse code: %s caused by %s", unauthorized.Code.String(), err)
		}

		h := crypto.Hash.New(crypto.SHA256)

		h.Write([]byte(unauthorized.Code.String()))

		hashed := h.Sum(nil)

		signed, err := rsa.SignPSS(rand.Reader, prv, crypto.SHA256, hashed, nil)
		if err != nil {
			panic(err)
		}

		signature := base64.StdEncoding.EncodeToString(signed)

		params := &openapi.V1AuthRefreshParams{
			Code:      unauthorized.Code.String(),
			Signature: signature,
		}

		res, err = client.Client.V1AuthRefresh(ctx, params)
		if err != nil {
			t.Fatalf("failed to refresh in: %s", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("failed to refresh in: %d", res.StatusCode)
		}

		cookies := res.Cookies()

		cookies = append(cookies, sid)

		client.Client.Client = &http.Client{
			Transport: helper.NewCookiesTransport(t, cookies),
		}

		res, err = client.Client.V1AuthVerify(ctx)
		if err != nil {
			t.Fatalf("failed to verify in: %s", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("failed to verify in: %d", res.StatusCode)
		}
	})
}
