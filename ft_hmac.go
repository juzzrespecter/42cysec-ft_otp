package main

// block_size == 64 bytes

/* hmac_algo:
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
)

const ipad = 0x36
const opad = 0x5c

func Hmac(k string, t string) string {
	if k > sha1.BlockSize {
		k = 
	}
}