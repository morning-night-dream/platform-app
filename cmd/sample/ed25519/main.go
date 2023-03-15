package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	pub, prv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(prv)
	if err != nil {
		panic(err)
	}

	prvstr := base64.StdEncoding.EncodeToString(bytes)

	fmt.Printf("%s\n", prvstr)

	decoded, err := base64.StdEncoding.DecodeString(prvstr)
	if err != nil {
		panic(err)
	}

	var key ed25519.PrivateKey
	if err := json.Unmarshal(decoded, &key); err != nil {
		panic(err)
	}

	message := "test"

	hashed := sha256.Sum256([]byte(message))

	sig, err := key.Sign(rand.Reader, hashed[:], crypto.SHA256)
	if err != nil {
		panic(err)
	}

	ok := ed25519.Verify(pub, hashed[:], sig)

	log.Printf("%t\n", ok)
}
