package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// nolint: gochecknoinits
func init() {
	l = log.New(os.Stderr, "", 0)

	flag.StringVar(&cmdURL, "url", "", "URL to Screenshot")
	flag.StringVar(&cmdBin, "bin", "/usr/bin/google-chrome-stable", "Path to Chrome/Chromium binary")

	flag.IntVar(&cmdWait, "time-wait", 5, "Time to wait for Page load")
	flag.IntVar(&cmdDeadLine, "time-deadline", 300, "Deadline in seconds")

	flag.BoolVar(&cmdFullPage, "fullpage", true, "Fullpage or Viewport")
	flag.IntVar(&cmdWindowWidth, "viewport-width", 1920, "Viewport Width")
	flag.IntVar(&cmdWindowHight, "viewport-hight", 1080, "Viewport Hight")

	flag.IntVar(&cmdPaddingTop, "padding-top", 0, "Viewport padding-top")
	flag.IntVar(&cmdPaddingBottom, "padding-bottom", 0, "Viewport padding-bottom")
	flag.IntVar(&cmdPaddingLeft, "padding-left", 0, "Viewport padding-left")
	flag.IntVar(&cmdPaddingRight, "padding-right", 0, "Viewport padding-right")

	flag.BoolVar(&cmdJpeg, "jpeg", true, "Create JPEG instead of PNG")
	flag.IntVar(&cmdJpegQuality, "jpeg-quality", 95, "JPEG quality")

	flag.StringVar(&cmdAcceptLanguage, "accept-language", "uk-UA,uk; ru-RU,ru; en-US,en", "Sets prefered language for page")
	flag.StringVar(&cmdUserAgent, "user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36", "Sets custom UserAgent")

	flag.BoolVar(&cmdRandomProfileDir, "random-profile", true, "Use random (newly created) profile")
	flag.StringVar(&cmdProfileDir, "profile", "", "Path to profile directory")

	flag.Parse()

	if len(cmdURL) == 0 {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()

		os.Exit(1)
	}
}
