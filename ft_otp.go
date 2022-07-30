package main

import (
	"flag"
	"os"
	"log"
	"encoding/hex"
)

func generate_new_password(key_file string) {

}

func store_new_key(user_input string) {
	if len(user_input) < 64 {
		log.Panic("error: key must not be lower than 64 bytes")
	}
	f, err := os.OpenFile("ft_otp.key", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.close()
	_, err := f.Write(key_encrypt(user_input))
	if err != nil {
		/* f closes when panic ? */
		log.Panic(err)
	}
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
		generate_new_password(*k)
	} else {

		store_new_key(*g)
	}
}
