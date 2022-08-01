package main

import (
	"testing"
	"crypto/hmac"
	"crypto/sha"
	"totp"
	"time"
)

K := []byte{...}

func TestHmac(t *testing.T) {
	mac := hmac.New(sha.New, K)
	T := totp_timestamp()

	expected_mac := hmac.Sum(T)
	our_mac, err2 := Hmac(K, T)

	if hmac.Equal(expected_mac, our_mac) || err1 != nil || err2 != nil {
		t.Fatalf(`Mismatch in hmac output. (ours %s, theirs %s)`, our_mac, expected_mac)
	}
}

func TestTotp(t *testing.T) {
	expectedtotp, err1 := totp.GenerateCode(K, time.Now())
	our_totp, err2 := Totp_new_code(K)

	if expected_totp != our_totp || err1 != nil || err2 != nil {
		t.Fatalf(`Mismatch in totp output. (ours %s, theirs %s)`, our_totp, expected_totp)
	}
}
