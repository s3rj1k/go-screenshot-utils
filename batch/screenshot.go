package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	screenshot "github.com/s3rj1k/go-screenshots"
)

func wrapScreenshot(c *screenshot.Config, worker int, path string) {

	var err error

	c.Port, err = screenshot.GetFreePort()
	if err != nil {
		fail.Printf("#%d: [%s] {%s}\n", worker, c.URL, err.Error())

		return
	}

	image, err := c.Screenshot()
	if err != nil {
		fail.Printf("#%d: [%s] {%s}\n", worker, c.URL, err.Error())

		return
	}

	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		fail.Printf("#%d: [%s] {%s}\n", worker, c.URL, err.Error())

		return
	}

	err = ioutil.WriteFile(path, image, 0644)
	if err != nil {
		fail.Printf("#%d: [%s] {%s}\n", worker, c.URL, err.Error())

		return
	}

	success.Printf("#%d: [%s] -> [%s]\n", worker, c.URL, path)
}
