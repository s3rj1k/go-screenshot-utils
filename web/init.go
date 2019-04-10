package main

import (
	"flag"
)

// nolint: gochecknoinits
func init() {
	flag.StringVar(&cmdBin, "cdp-bin", "/usr/bin/google-chrome-stable", "path to Chrome/Chromium binary")
	flag.IntVar(&cmdDeadLine, "time-deadline", 300, "deadline in seconds")
	flag.IntVar(&cmdPort, "listen-port", 8888, "tcp port for web server")

	flag.Parse()
}
