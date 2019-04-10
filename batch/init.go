package main

import (
	"log"
	"os"
)

// nolint: gochecknoinits
func init() {
	info = log.New(os.Stderr, "[?] ", log.LstdFlags)
	fail = log.New(os.Stderr, "[-] ", log.LstdFlags)
	success = log.New(os.Stderr, "[+] ", log.LstdFlags)
}
