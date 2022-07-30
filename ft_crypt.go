package main

import (
	"crypto/aes"
	"encoding/base32"
)

/* Utilities for key encryption/decryption */

const secret_key = "abubilla"

func key_decrypt(key ) {}

func key_encrypt(src string) string{
	dst := make([]byte, len(toEncrypt))
	c, err := aes.NewCypher(secret_key)
	
	c.Encrypt(dst, []byte(src))
	fmt.Println(dst)
	return base32.Encode(dst)
}
