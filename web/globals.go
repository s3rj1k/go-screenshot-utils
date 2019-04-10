package main

const (
	defaultDeadline      = 15 * 60
	defaultTimeWait      = 15
	maxTimeWait          = defaultDeadline
	defaultViewportWidth = 1920
	maxViewportWidth     = 4096
	defaultViewportHight = 1080
	maxViewportHight     = 2160
	defaultIsFullpage    = true
)

// nolint: gochecknoglobals
var (
	cmdBin      string
	cmdDeadLine int
	cmdPort     int
)
