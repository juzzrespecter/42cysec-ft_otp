package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"github.com/pquerna/otp/totp"
	"testing"
	"time"
)

var K = []byte{'t', 'h', 'i', 's', 'i', 's', 'a', 't', 'e', 's', 't', '1', '2', '3', '4'}

func TestHmac(t *testing.T) {
	mac := hmac.New(sha1.New, K)
	T := totp_timestamp()

	their_mac := mac.Sum(T)
	our_mac, err2 := Hmac(K, T)

	if hmac.Equal(their_mac, our_mac[:]) || err2 != nil {
		t.Fatalf(`Mismatch in hmac output. (ours %s, theirs %s)`, our_mac, their_mac)
	}
}

func TestTotp(t *testing.T) {
	their_K := make([]byte, base32.StdEncoding.EncodedLen(len(K)))

	base32.StdEncoding.Encode(their_K, K)
	their_totp, err1 := totp.GenerateCode(string(their_K), time.Now())
	our_totp_n, err2 := totp_new_code(K)
	our_totp := fmt.Sprintf("%06d", our_totp_n)

	if err1 != nil {
		t.Fatalf("An exception occured... %s", err1)
	}
	if their_totp != our_totp || err1 != nil || err2 != nil {
		t.Fatalf(`Mismatch in totp output. (ours %q, theirs %q)`, our_totp, their_totp)
	}
}
