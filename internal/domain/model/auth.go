package model

import (
	"context"
	"time"
)

type Auth struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	IDToken      string    `json:"idToken"`
	PublicKey    []byte    `json:"publicKey"`
	RefreshToken string    `json:"refreshToken"`
	SessionToken string    `json:"sessionToken"`
	ExpiresIn    int       `json:"expiresIn"`
	Expires      time.Time `json:"expires"`
}

type uidCtxKey struct{}

func SetUIDCtx(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, uidCtxKey{}, uid)
}

func GetUIDCtx(ctx context.Context) string {
	v := ctx.Value(uidCtxKey{})

	id, ok := v.(string)
	if !ok {
		return ""
	}

	return id
}
