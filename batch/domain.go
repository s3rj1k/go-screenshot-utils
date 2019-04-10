package main

import (
	"strings"
)

func getDomainFromURL(url string) string {

	r := strings.NewReplacer("https://", "", "http://", "")
	noScheme := r.Replace(url)

	if strings.Contains(noScheme, "/") {
		slice := strings.Split(noScheme, "/")

		return slice[0]
	}

	return noScheme
}

func getSubDomainPrefix(domain, gtldPlusOne string) string {

	prefix := strings.TrimSuffix(domain, gtldPlusOne)

	if len(prefix) > 0 && strings.HasSuffix(prefix, ".") && prefix != "." {
		return strings.TrimSuffix(prefix, ".")
	}

	return "_"
}
