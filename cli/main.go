package main

import (
	"encoding/base64"
	"fmt"
	"time"

	screenshot "github.com/s3rj1k/go-screenshots"
)

func main() {

	c := screenshot.DefaultConfig()

	c.CMD = cmdBin
	c.URL = cmdURL
	c.Wait = time.Duration(cmdWait) * time.Second
	c.Deadline = time.Duration(cmdDeadLine) * time.Second
	c.WindowWidth = cmdWindowWidth
	c.WindowHight = cmdWindowHight
	c.FullPage = cmdFullPage
	c.IsJPEG = cmdJpeg
	c.JpegQuality = cmdJpegQuality
	c.AcceptLanguage = cmdAcceptLanguage
	c.UserAgent = cmdUserAgent

	c.PaddingBottom = cmdPaddingBottom
	c.PaddingLeft = cmdPaddingLeft
	c.PaddingRight = cmdPaddingRight
	c.PaddingTop = cmdPaddingTop

	c.RandomProfileDir = cmdRandomProfileDir
	c.ProfileDir = cmdProfileDir

	if image, err := c.Screenshot(); err == nil {
		fmt.Println(base64.StdEncoding.EncodeToString(image))
	}
}
