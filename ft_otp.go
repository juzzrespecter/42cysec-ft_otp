package main

import (
	"flag"
	"os"
	"log"
	"encoding/hex"
	"io/ioutils"
	"time"
	"math"
)

func totp_truncate([20]byte mac) {
	/* Dynamic truncation */
	offs := mac[19] & 0x0f
	dbc1 := (mac[offs] & 0x7f) << 24
	| (mac[offs+1]) << 16
	| (mac[offs+2]) << 8
	| mac[offs+3]

	dbc2 := dbc1 % math.Pow(10, 6)
	return dbc2
}

func generate_new_password(key_file string) err {
	key, err := ioutils.ReadFile(key_file)
	if err != nil {
		return err
	}
	K, err := key_decrypt(key) /* 1st param. of HMAC-SHA-1 */
	if err != nil {
		return err
	}
	T := make([]byte, 8)
	timestamp := time.Now().Unix()
	for i := 0; i < 4; i++ {
		T[i + 4] = uint8(timestamp >> (24 - (i*8)))
	} /* 2nd param. of HMAC-SHA-1 */
	mac, err := Hmac(K, T)
	if err != nil {
		return err
	}
	totp := hmac_truncate(mac)
	fmt.Println(totp)
}

func store_new_key(user_input string) err {
	if len(user_input) < 64 {
		return errors.New("error: key must not be lower than 64 bytes")
	}
	plain_key, err := hex.DecodeString(user_input)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("ft_otp.key", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.close()
	encrypted_key, err := key_encrypt(plain_key)
	if err != nil {
		return err
	}
	_, err := f.Write(encrypted_key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args[1:]) > 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	g := flag.String("g", "" ,"encrypt and store key given by user")
	k := flag.String("k", "", "generate new temporary password")

	flag.Parse()
	if *g != "" && *k != "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *k != "" {
		err = generate_new_password(*k)
		if err != nil {
			fmt.Fprintf(os.Stderr, err)
		}
	} else {
		err = store_new_key(*g)
		if err != nil {
			fmt.Fprintf(os.Stderr, err)
		}
	}
}
