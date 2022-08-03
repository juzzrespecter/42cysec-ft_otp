package main

import (
	"crypto/aes"
	"encoding/base32"
	"bytes"
)

var secret_key = []byte{0x61, 0x62, 0x75, 0x62, 0x69, 0x6c, 0x6c, 0x61,
	0x75, 0x6c, 0x74, 0x72, 0x61, 0x6d, 0x61, 0x72} /* 128-bit key */

func key_decode_base32(k []byte) ([]byte, error) {
	dec_k := make([]byte, base32.StdEncoding.DecodedLen(len(k)))
	_, err := base32.StdEncoding.Decode(dec_k, k)

	if err != nil {
		return []byte{}, err
	}
	dec_k = bytes.Trim(dec_k, "\x00")
	return dec_k, nil
}

func key_decrypt(key []byte) ([]byte, error) {
	src, err := key_decode_base32(key)

	if err != nil {
		return []byte{}, err
	}
	dst := make([]byte, len(src))
	c, err := aes.NewCipher(secret_key)

	if err != nil {
		return []byte{}, err
	}
	r := len(src) / 16
	
	for i := 0; i < r; i++ {
		idx := i * 16
		end := idx + 16

		c.Decrypt(dst[idx:end], src[idx:end])
	}
	dst = bytes.Trim(dst, "\x00")
	return dst, nil
}

func key_encode_base32(k []byte) string {
	enc_k := make([]byte, base32.StdEncoding.EncodedLen(len(k)))

	base32.StdEncoding.Encode(enc_k, k)
	return string(enc_k)
}

func key_encrypt(src []byte) (string, error) {
	c, err := aes.NewCipher(secret_key)

	if err != nil {
		return "", err
	}

	pad := len(src) % c.BlockSize()
	
	if pad != 0 {
		src = append(src, make([]byte, c.BlockSize() - pad)...)
	}
	r := len(src) / 16
	dst := make([]byte, len(src))
	
	for i := 0; i < r; i++ {
		idx := i * 16
		end := idx + 16

		c.Encrypt(dst[idx:end], src[idx:end])
	}
	return key_encode_base32(dst), nil
}
