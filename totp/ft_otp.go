package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func print_usage() {
	fmt.Fprintln(os.Stderr, "Usage of ./ft_otp:")
	flag.PrintDefaults()
	os.Exit(1)
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* Dynamic truncation of hmac hash */
func totp_truncate(mac [20]byte) int32 {
	offs := mac[19] & 0x0f
	var dbc1 int = (int(mac[offs])&0x7f)<<24 |
		(int(mac[offs+1]))<<16 |
		(int(mac[offs+2]))<<8 |
		int(mac[offs+3])

	dbc2 := int32(dbc1) % int32(math.Pow(10, 6))
	return dbc2
}

func totp_timestamp() []byte {
	T := make([]byte, 8)
	timestamp := time.Now().Unix() / 30

	for i := 0; i < 4; i++ {
		T[i+4] = uint8(timestamp >> (24 - (i * 8)))
	}
	return T
}

func totp_new_code(K []byte) (int32, error) {
	T := totp_timestamp()
	mac, err := Hmac(K, T)

	if err != nil {
		return 0, err
	}
	totp_key := totp_truncate(mac)

	return totp_key, nil
}

func generate_code(key_file string) (int32, error) {
	key, err := os.ReadFile(key_file)

	if err != nil {
		return 0, err
	}
	K, err := key_decrypt(key)

	if err != nil {
		return 0, err
	}
	code, err := totp_new_code(K)

	if err != nil {
		return 0, err
	}
	return code, nil
}

func store_new_key(user_input string) error {
	if len(user_input) < 64 || (len(user_input)%2) != 0 {
		return errors.New("error: key must be hex encoded and not lower than 64 bytes")
	}
	plain_key, err := hex.DecodeString(user_input)

	if err != nil {
		return err
	}
	f, err := os.OpenFile("ft_otp.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		return err
	}
	defer f.Close()
	crypt_key, err := key_encrypt(plain_key)

	if err != nil {
		return err
	}
	_, err = f.Write([]byte(crypt_key))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	g := flag.String("g", "", "encrypt and store key given by user")
	k := flag.String("k", "", "generate new temporary password")

	flag.Parse()
	if len(os.Args[1:]) != 2 {
		print_usage()
	}
	if (*g != "" && *k != "") || (*g == "" && *k == "") {
		print_usage()
	}
	if *k != "" {
		code, err := generate_code(*k)
		check_error(err)
		fmt.Printf("%06d\n", code)
	} else {
		err := store_new_key(*g)
		check_error(err)
		fmt.Printf("\033[32m [OK] \033[0m ✨  Stored new key in file \033[32m ft_otp.key \033[0m ✨\n")
	}
}
