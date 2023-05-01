package external

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/rpc"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

var _ rpc.Auth = (*Auth)(nil)

type AuthFactory interface {
	Auth(secret string, endpoint string, apiKey string) (*Auth, error)
}

type Auth struct {
	Endpoint     string
	APIKey       string
	HTTPClient   *http.Client
	FirebaseAuth *auth.Client
}

func (fb *Auth) SignUp(ctx context.Context, uid string, email model.EMail, password model.Password) error {
	params := (&auth.UserToCreate{}).
		UID(string(uid)).
		Email(string(email)).
		EmailVerified(false).
		Password(string(password)).
		Disabled(false)

	if _, err := fb.FirebaseAuth.CreateUser(ctx, params); err != nil {
		log.GetLogCtx(ctx).Warn("faile to create user", log.ErrorField(err))

		return err
	}
	return nil
}

type SignInRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignInResponse struct {
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

func (fb *Auth) SignIn(ctx context.Context, email model.EMail, password model.Password) (*model.User, error) {
	// https://firebase.google.com/docs/reference/rest/auth#section-sign-in-email-password
	url := fmt.Sprintf("%s/v1/accounts:signInWithPassword?key=%s", fb.Endpoint, fb.APIKey)

	req := SignInRequest{
		Email:             string(email),
		Password:          string(password),
		ReturnSecureToken: true,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(req)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to encode json", log.ErrorField(err))

		return nil, err
	}

	res, err := fb.HTTPClient.Post(url, "application/json", &buf)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to post "+url, log.ErrorField(err))

		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		message, err := io.ReadAll(res.Body)
		if err != nil {
			log.GetLogCtx(ctx).Warn("faile to read body", log.ErrorField(err))

			message = []byte(fmt.Sprintf("could not laod message caused by %v", err))
		}

		return nil, fmt.Errorf("firebase error. status code is %d, message is %v", res.StatusCode, string(message))
	}

	var resp SignInResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode json", log.ErrorField(err))

		return nil, err
	}

	// exp, _ := strconv.Atoi(res.ExpiresIn)

	strs := strings.Split(resp.IDToken, ".")

	tmpPayload, err := base64.RawStdEncoding.DecodeString(strs[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	type Payload struct {
		UserID string `json:"user_id"`
	}

	var payload Payload

	if err := json.Unmarshal(tmpPayload, &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	return &model.User{
		UserId: payload.UserID,
	}, nil
}
