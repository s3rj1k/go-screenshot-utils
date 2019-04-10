package main

import (
	"io/ioutil"
	"strings"
)

func readURLsFromFile(file string) ([]string, error) {

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(bytes), "\n"), nil
}
