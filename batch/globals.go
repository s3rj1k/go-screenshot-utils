package main

import (
	"log"
)

// nolint: gochecknoglobals
var (
	maxWorkers *int

	info    *log.Logger
	fail    *log.Logger
	success *log.Logger
)
