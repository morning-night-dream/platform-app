package gateway

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
)

var _ repository.APIAuth = (*APIAuth)(nil)

type APIAuth struct {
	firebase *firebase.Client
}

func (aa *APIAuth) SignUp(ctx context.Context, uid model.UserID, email model.EMail, password model.Password) error {
	if err := aa.firebase.CreateUser(ctx, string(uid), string(email), string(password)); err != nil {
		return err
	}

	return nil
}

func (aa *APIAuth) SignIn(ctx context.Context, email model.EMail, password model.Password) (model.Auth, error) {
	res, err := aa.firebase.Login(ctx, string(email), string(password))
	if err != nil {
		return model.Auth{}, err
	}

	// exp, _ := strconv.Atoi(res.ExpiresIn)

	strs := strings.Split(res.IDToken, ".")

	tmpPayload, err := base64.RawStdEncoding.DecodeString(strs[1])
	if err != nil {
		return model.Auth{}, err
	}

	type Payload struct {
		UserID string `json:"user_id"`
	}

	var payload Payload

	if err := json.Unmarshal(tmpPayload, &payload); err != nil {
		return model.Auth{}, err
	}

	return model.Auth{
		UserID:       model.UserID(payload.UserID),
		IDToken:      model.IDToken(res.IDToken),
		RefreshToken: model.RefreshToken(res.RefreshToken),
	}, nil
}

func (aa *APIAuth) Refresh(context.Context, model.RefreshToken) (model.RefreshToken, error) {
	return model.RefreshToken(""), nil
}

func (aa *APIAuth) Verify(context.Context) error {
	return nil
}

func (aa *APIAuth) Delete(context.Context, model.UserID) error {
	return nil
}

func (aa *APIAuth) Save(context.Context, model.Auth) error {
	return nil
}

func (aa *APIAuth) Find(context.Context, model.UserID) (model.Auth, error) {
	return model.Auth{}, nil
}
