package main

import (
	"crypto/aes"
	"encoding/base32"
	"fmt"
)

var secret_key = []byte{ 0x61, 0x62, 0x75, 0x62, 0x69, 0x6c, 0x6c, 0x61,
	0x75, 0x6c, 0x74, 0x72, 0x61, 0x6d, 0x61, 0x72 } /* 128-bit key */

func key_decode_base32(k []byte) ([]byte, error) {
	dec_k := make([]byte, base32.StdEncoding.DecodedLen(len(k)))
	_, err := base32.StdEncoding.Decode(dec_k, k)
	if err != nil {
		return []byte{}, err
	}
	fmt.Printf("%q\n", dec_k) /* debug, to remove */
	return dec_k, nil
}

func key_decrypt(key []byte) ([]byte, error) {
	dst := make([]byte, len(key))
	c, err := aes.NewCipher(secret_key)
	if err != nil {
		return []byte{}, err
	}
	raw_key, err := key_decode_base32(key)
	if err != nil {
		return []byte{}, err
	}
	c.Decrypt(dst, raw_key)
	return dst, nil
}

func key_encode_base32(k []byte) string {
	enc_k := make([]byte, base32.StdEncoding.EncodedLen(len(k)))
	base32.StdEncoding.Encode(enc_k, k)
	fmt.Printf("%q\n", enc_k) /* debug, to remove */
	return string(enc_k);
}

func key_encrypt(src []byte) (string, error) {
	dst := make([]byte, len(src))
	c, err := aes.NewCipher(secret_key)

	if err != nil {
		return "", err
	}
	c.Encrypt(dst, src)
	fmt.Println(dst) /* test */
	return key_encode_base32(dst), nil
}
