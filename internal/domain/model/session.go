package model

import (
	"crypto/rsa"
	"time"
)

const Age = 2592000 * time.Second // 30 days

type SessionID string

type SessionToken string

type Session struct {
	SessionID SessionID      `json:"id"`
	UserID    UserID         `json:"user_id"`
	PublicKey *rsa.PublicKey `json:"key"`
}
