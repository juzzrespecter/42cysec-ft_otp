package main

import (
	"testing"
	"crypto/hmac"
	"crypto/sha"
	"totp"
)

key := []byte{...}

func TestHmac(t *testing.T) {
	mac := hmac.New(sha.New, key)
}

func TestTotp(t *testing.T) {
}
