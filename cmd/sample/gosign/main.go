package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

func main() {
	url := "http://localhost:8082"

	client, err := openapi.NewClient(url + "/api")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	email := types.Email(uuid.NewString() + "@example.com")

	password := uuid.NewString()

	// signup
	if _, err := client.V1AuthSignUp(ctx, openapi.V1AuthSignUpJSONRequestBody{
		Email:    email,
		Password: password,
	}); err != nil {
		panic(err)
	}

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

	log.Printf("pubstr: %+v", pubstr)

	// signin
	res, err := client.V1AuthSignIn(ctx, openapi.V1AuthSignInJSONRequestBody{
		Email:     email,
		Password:  password,
		PublicKey: pubstr,
	})
	if err != nil {
		panic(err)
	}

	client.Client = &http.Client{
		Transport: NewCookieTransport(res.Cookies()),
	}

	// sign
	code := "reichankawaii"

	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(code))
	hashed := h.Sum(nil)

	signed, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	if err != nil {
		panic(err)
	}

	signature := base64.StdEncoding.EncodeToString(signed)

	res, err = client.V1Sign(ctx, &openapi.V1SignParams{
		Code:      code,
		Signature: signature,
	})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", res)
}

type CookieTransport struct {
	Cookies   []*http.Cookie
	Transport http.RoundTripper
}

func NewCookieTransport(
	cookies []*http.Cookie,
) *CookieTransport {
	return &CookieTransport{
		Cookies:   cookies,
		Transport: http.DefaultTransport,
	}
}

func (ct *CookieTransport) transport() http.RoundTripper {
	return ct.Transport
}

func (ct *CookieTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, cookie := range ct.Cookies {
		fmt.Printf("cookie: %+v\n", cookie)
		req.AddCookie(cookie)
	}

	resp, err := ct.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
