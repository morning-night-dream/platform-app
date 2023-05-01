package model

import (
	"context"
	"crypto/rsa"
	"time"
)

type EMail string

type Password string

type IDToken string // FirebaseのIDTokenではなく独自のトークン

type RefreshToken string // FirebaseTokenとかの命名の方がよいかも

type CodeID string

type Code struct {
	CodeID    CodeID
	SessionID SessionID
}

type Signature string

type Auth struct {
	ID           string         `json:"id"`
	UserID       string         `json:"userId"`
	IDToken      IDToken        `json:"idToken"`
	PublicKey    *rsa.PublicKey `json:"publicKey"`
	RefreshToken RefreshToken   `json:"refreshToken"`
	ExpiresIn    int            `json:"expiresIn"`
	Expires      time.Time      `json:"expires"`
}

func (auth Auth) Verify() error {
	return nil
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
