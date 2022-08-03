package main

/* hmac_algorithm:
1. append 0 to k < 64 bytes
2. K xor [0x36 * 64]
3. append message to result string
4. apply H (SHA-1)
5. K xor [0x5c * 64]
6. append (5) | (4)
7. apply H

HMAC(K,C) = H(K xor opad | H(K xor ipad | msg))
*/

import (
	"crypto/sha1"
	"errors"
)

const (
	ipad_b = 0x36 /* inner padding byte */
	opad_b = 0x5c /* outer padding byte */
	B      = 64   /* in bytes, size of block data */
	L      = 20   /* in bytes, size of SHA1 output */
)

func hmac_xor(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("xor with different size arrays")
	}
	xor := make([]byte, len(a))

	for i := 0; i < len(a); i++ {
		xor[i] = a[i] ^ b[i]
	}
	return xor, nil
}

func hmac_init_mask(val uint8, len int) []byte {
	mask := make([]byte, len)

	for i := 0; i < len; i++ {
		mask[i] = val
	}
	return mask
}

func Hmac(k []byte, t []byte) ([L]byte, error) {
	ipad := hmac_init_mask(ipad_b, B)
	opad := hmac_init_mask(opad_b, B)

	if len(k) > B {
		hk := sha1.Sum(k)
		k = hk[:]
	}
	if len(k) < B {
		pad := make([]byte, B-len(k))
		k = append(k, pad...)
	}
	k_xor_ipad, err := hmac_xor(k, ipad)
	if err != nil {
		return [20]byte{}, err
	}
	k_h1 := append(k_xor_ipad, t...)
	h1 := sha1.Sum(k_h1)
	k_xor_opad, err := hmac_xor(k, opad)
	if err != nil {
		return [20]byte{}, err
	}
	k_h2 := append(k_xor_opad, h1[:]...)
	return sha1.Sum(k_h2), nil
}
