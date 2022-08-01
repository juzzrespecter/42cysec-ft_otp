package main

import (
	"crypto/aes"
	"encoding/base32"
)

/* Utilities for key encryption/decryption */

const secret_key = "abubilla"

func key_decrypt(key string) ([]byte, err) {
	dst := make([]byte, len(key))
	c, err := aes.NewCypher(key)
	if err != nil {
		return []byte{}, err
	}
	raw_key := base32.Decode(key)
	aes.Decrypt(dst, key)
	return dst, nil
}

func key_encrypt(src []byte) string {
	dst := make([]byte, len(toEncrypt))
	c, err := aes.NewCypher(secret_key)

	if err != nil {
		return nil, err
	}
	c.Encrypt(dst, src)
	fmt.Println(dst) /* test */
	return base32.Encode(dst), nil
}
