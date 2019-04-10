package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		log := []string{}

		log = append(log, fmt.Sprintf("%v", time.Now().Format("2006/01/02 - 15:04:05")))
		log = append(log, fmt.Sprintf("status: %d", c.Writer.Status()))
		log = append(log, fmt.Sprintf("remote_ip: %s", c.ClientIP()))
		log = append(log, fmt.Sprintf("user_agent: %s", c.Request.UserAgent()))
		log = append(log, fmt.Sprintf("method: %s", c.Request.Method))

		if len(query) != 0 {
			log = append(log, fmt.Sprintf("path: %s?%s", path, query))
		} else {
			log = append(log, fmt.Sprintf("path: %s", path))
		}

		if path == "/screenshot" {
			form, err := parseScreenshotForm(c)
			if err == nil {
				log = append(log, fmt.Sprintf("remote_url: %s", form.remoteURL))
				log = append(log, fmt.Sprintf("wait_time: %d sec", form.timeWait))
				log = append(log, fmt.Sprintf("fullscreen: %t", form.isFullpage))
				log = append(log, fmt.Sprintf("size: %dx%d", form.viewportWidth, form.viewportHight))
				log = append(log, fmt.Sprintf("emulated_user_agent: %s", form.userAgent))
				log = append(log, fmt.Sprintf("accept_language: %s", form.acceptLanguage))
			}
		}

		if len(c.Errors) > 0 {
			log = append(log, strings.Replace(c.Errors.String(), "\n", "; ", -1))
		}

		fmt.Fprintf(gin.DefaultWriter, "%s\n", strings.Join(log, " | "))
	}
}
