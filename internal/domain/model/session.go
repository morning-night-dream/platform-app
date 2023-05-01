package model

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"time"
)

const Age = 2592000 * time.Second // 30 days

type SessionID string

type SessionToken string

func (sss *Session) RSAPublicKey() (*rsa.PublicKey, error) {
	var p *rsa.PublicKey

	if err := json.Unmarshal([]byte(sss.PublicKey), p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	return p, nil
}

func PublicKeyToString(key *rsa.PublicKey) (string, error) {
	bytes, err := json.Marshal(key)
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}

	return string(bytes), nil
}
