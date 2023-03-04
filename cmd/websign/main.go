package main

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	b := new(bytes.Buffer)

	bt, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)

	pem.Encode(b, &pem.Block{
		Bytes: bt,
	})

	pems := strings.Split(b.String(), "\n")

	pems = remove(pems, len(pems)-1)

	pems = remove(pems, len(pems)-1)

	pems = remove(pems, 0)

	// signin
	res, err := client.V1AuthSignIn(ctx, openapi.V1AuthSignInJSONRequestBody{
		Email:     email,
		Password:  password,
		PublicKey: strings.Join(pems, ""),
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

	signed, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hashed, nil)
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

	log.Printf("%+v", res.StatusCode)
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

func remove(arr []string, i int) []string {
	return append(arr[:i], arr[i+1:]...)
}
