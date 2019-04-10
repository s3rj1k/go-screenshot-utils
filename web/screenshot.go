package main

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	screenshot "github.com/s3rj1k/go-screenshots"
)

func createScreenshotFromWeb(engine *gin.Engine) {
	engine.POST("/screenshot", func(c *gin.Context) {

		form, err := parseScreenshotForm(c)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}

		cdp := screenshot.DefaultConfig()
		cdp.CMD = cmdBin
		cdp.URL = form.remoteURL
		cdp.Wait = time.Duration(form.timeWait) * time.Second
		cdp.Deadline = time.Duration(cmdDeadLine) * time.Second
		cdp.WindowWidth = form.viewportWidth
		cdp.WindowHight = form.viewportHight
		cdp.FullPage = form.isFullpage
		cdp.IsJPEG = false
		cdp.UserAgent = form.userAgent
		cdp.AcceptLanguage = form.acceptLanguage

		image, err := cdp.Screenshot()
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.String(http.StatusOK, base64.StdEncoding.EncodeToString(image))
	})
}
