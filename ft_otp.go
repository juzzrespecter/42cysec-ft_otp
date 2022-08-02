package main

import (
	"flag"
	"os"
	"log"
	"encoding/hex"
	"time"
	"math"
	"errors"
	"fmt"
)

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* Dynamic truncation of hmac hash */
func totp_truncate(mac [20]byte) int32 {
	offs := mac[19] & 0x0f
	dbc1 := (mac[offs] & 0x7f) << 24 |
		(mac[offs+1]) << 16 |
		(mac[offs+2]) << 8 |
		 mac[offs+3]

	dbc2 := int32(dbc1) % int32(math.Pow(10, 6))
	return dbc2
}

func totp_timestamp() []byte {	
	T := make([]byte, 8)
	timestamp := time.Now().Unix()
	for i := 0; i < 4; i++ {
		T[i + 4] = uint8(timestamp >> (24 - (i*8)))
	}
	return T
}

/* Totp code generator */
func totp_new_code(K []byte) (int32, error) {
	T := totp_timestamp()
	mac, err := Hmac(K, T)
	if err != nil {
		return 0, err
	}
	totp_key := totp_truncate(mac)
	return totp_key, nil
}

func store_new_key(user_input string) error {
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
	defer f.Close()
	encrypted_key, err := key_encrypt(plain_key)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(encrypted_key))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	g := flag.String("g", "" ,"encrypt and store key given by user")
	k := flag.String("k", "", "generate new temporary password")

	flag.Parse()
	if len(os.Args[1:]) != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *g != "" && *k != "" {
		fmt.Fprintln(os.Stderr, "Usage of ./ft_otp:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *k != "" {
		key, err := os.ReadFile(*k)
		check_error(err)
		
		K, err := key_decrypt(key)
		check_error(err)
		
		code, err := totp_new_code(K)
		check_error(err)

		fmt.Println(code)
	} else {
		err := store_new_key(*g)
		check_error(err)
	}
}
