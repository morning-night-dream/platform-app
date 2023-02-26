package model

import "crypto/rsa"

type SessionID string

type Session struct {
	SessionID SessionID      `json:"id"`
	UserID    string         `json:"user_id"`
	PublicKey *rsa.PublicKey `json:"key"`
}
