package model

import "crypto/rsa"

type SessionID string

type SessionToken string

type Session struct {
	SessionID SessionID      `json:"id"`
	UserID    UserID         `json:"user_id"`
	PublicKey *rsa.PublicKey `json:"key"`
}
