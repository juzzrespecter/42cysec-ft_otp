package main

import (
	"flag"
	"os"
	"fmt"
	"log"
)

func generate_new_password(key_file string) {

}

func store_new_key(user_input string) {
	if len(user_input) < 64 {
		log.Panic("error: key must not be lower than 64 bytes")
	}
	key_file, err := os.OpenFile("key")
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
