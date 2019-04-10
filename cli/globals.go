package main

import (
	"log"
)

// nolint: gochecknoglobals
var (
	l *log.Logger

	cmdBin string

	cmdURL string

	cmdFullPage      bool
	cmdWindowWidth   int
	cmdWindowHight   int
	cmdBottomPadding int

	cmdJpeg        bool
	cmdJpegQuality int

	cmdWait     int
	cmdDeadLine int

	cmdRandomProfileDir bool
	cmdProfileDir       string

	cmdAcceptLanguage string
	cmdUserAgent      string
)
