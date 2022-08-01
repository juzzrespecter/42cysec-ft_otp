package main

import (
	"crypto/aes"
	"encoding/base32"
	"fmt"
)

var secret_key = [16]byte{ 0x61, 0x62, 0x75, 0x62, 0x69, 0x6c, 0x6c, 0x61,
	0x61, 0x62, 0x75, 0x62, 0x69, 0x6c, 0x6c, 0x61 } /* 128-bit key */

func key_decrypt(key []byte) ([]byte, error) {
	dst := make([]byte, len(key))
	c, err := aes.NewCypher(secret_key)
	if err != nil {
		return []byte{}, err
	}
	raw_key := base32.Decode(key)
	aes.Decrypt(dst, key)
	return dst, nil
}

func key_encrypt(src []byte) (string, error) {
	dst := make([]byte, len(src))
	c, err := aes.NewCypher(secret_key)

	if err != nil {
		return "", err
	}
	c.Encrypt(dst, src)
	fmt.Println(dst) /* test */
	return base32.Encode(dst), nil
}
