package main

import (
	"log"
)

// nolint: gochecknoglobals
var (
	l *log.Logger

	cmdBin string

	cmdURL string

	cmdFullPage    bool
	cmdWindowWidth int
	cmdWindowHight int

	cmdPaddingBottom int
	cmdPaddingLeft   int
	cmdPaddingRight  int
	cmdPaddingTop    int

	cmdJpeg        bool
	cmdJpegQuality int

	cmdWait     int
	cmdDeadLine int

	cmdRandomProfileDir bool
	cmdProfileDir       string

	cmdAcceptLanguage string
	cmdUserAgent      string
)
