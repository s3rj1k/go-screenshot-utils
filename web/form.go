package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/* # User-Agent MUST be URL Encoded
curl -s -X POST http://zuul.s3rj1k.lt:8888/screenshot \
  -d 'url=https://mirohost.net/' \
  -d 'time-wait=5' \
  -d 'viewport-width=1920' \
  -d 'viewport-hight=1080' \
  -d 'fullpage=true' \
  -d 'user-agent=Mozilla/5.0' \
  -d 'accept-language=ru-RU,ru;'
*/

type screenshotForm struct {
	remoteURL      string
	userAgent      string
	acceptLanguage string
	timeWait       int
	viewportWidth  int
	viewportHight  int
	isFullpage     bool
}

func parseScreenshotForm(c *gin.Context) (screenshotForm, error) {

	var form screenshotForm

	formURL := strings.TrimSpace(c.PostForm("url"))
	formTimeWait := strings.TrimSpace(c.PostForm("time-wait"))
	formViewportWidth := strings.TrimSpace(c.PostForm("viewport-width"))
	formViewportHight := strings.TrimSpace(c.PostForm("viewport-hight"))
	formFullpage := strings.TrimSpace(strings.ToLower(c.PostForm("fullpage")))
	formUserAgent := strings.TrimSpace(c.PostForm("user-agent"))
	formAcceptLanguage := strings.TrimSpace(c.PostForm("accept-language"))

	if len(formURL) == 0 {
		return screenshotForm{}, fmt.Errorf("%s", "empty remote URL")
	}

	_, err := url.ParseRequestURI(formURL)
	if err != nil {
		return screenshotForm{}, fmt.Errorf("%s: %s", "Invalid remote URL", formURL)
	}

	form.remoteURL = formURL

	if len(formTimeWait) != 0 {
		if i, err := strconv.Atoi(formTimeWait); err == nil {
			if i >= 0 && i < maxTimeWait {
				form.timeWait = i
			} else {
				form.timeWait = maxTimeWait
			}
		} else {
			form.timeWait = defaultTimeWait
		}
	} else {
		form.timeWait = defaultTimeWait
	}

	if len(formViewportWidth) != 0 {
		if i, err := strconv.Atoi(formViewportWidth); err == nil {
			if i > 0 && i <= maxViewportWidth {
				form.viewportWidth = i
			} else {
				form.viewportWidth = defaultViewportWidth
			}
		} else {
			form.viewportWidth = defaultViewportWidth
		}
	} else {
		form.viewportWidth = defaultViewportWidth
	}

	if len(formViewportHight) != 0 {
		if i, err := strconv.Atoi(formViewportHight); err == nil {
			if i > 0 && i <= maxViewportHight {
				form.viewportHight = i
			} else {
				form.viewportHight = defaultViewportHight
			}
		} else {
			form.viewportHight = defaultViewportHight
		}
	} else {
		form.viewportHight = defaultViewportHight
	}

	switch formFullpage {
	case "true":
		form.isFullpage = true
	case "false":
		form.isFullpage = false
	default:
		form.isFullpage = defaultIsFullpage
	}

	if len(formUserAgent) != 0 {
		form.userAgent = formUserAgent
	}

	if len(formAcceptLanguage) != 0 {
		form.acceptLanguage = formAcceptLanguage
	}

	return form, nil
}
