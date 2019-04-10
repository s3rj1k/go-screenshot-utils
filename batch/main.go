package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"golang.org/x/net/publicsuffix"

	screenshot "github.com/s3rj1k/go-screenshots"
)

func main() {

	in := flag.String("in", "", "file with list of URLs")
	bin := flag.String("bin", "/usr/bin/google-chrome-stable", "path to Chrome/Chromium binary")
	wait := flag.Int("wait", 5, "seconds to wait for animation render")
	deadline := flag.Int("deadline", 300, "deadline in seconds")
	maxWorkers = flag.Int("max-workers", 5, "number of concurrent threads")
	screenshotsDir := flag.String("dir", ".", "path prefix to store screenshots")

	flag.Parse()

	if len(*in) == 0 {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()

		os.Exit(1)
	}

	urls, err := readURLsFromFile(*in)
	if err != nil {
		fail.Fatalf("no input file found: %s\n", err.Error())
	}

	pool := make(chan int, *maxWorkers)
	for w := 1; w <= cap(pool); w++ {
		pool <- w
	}

	jobs := make(chan string, len(urls))
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg := sync.WaitGroup{}

	// https://golang.org/doc/faq#closures_and_goroutines
	for j := range jobs {

		w := <-pool
		wg.Add(1)

		go func(job string) {

			defer func(w int, pool chan int) {
				pool <- w

				wg.Done()
			}(w, pool)

			currentTime := time.Now().Local()

			domain := getDomainFromURL(job)

			gtldPlusOne, err := publicsuffix.EffectiveTLDPlusOne(domain)
			if err != nil {
				fail.Printf("#%d: [%s] {%s}\n", w, job, err.Error())

				return
			}

			firstLetter := string([]rune(gtldPlusOne)[0])
			subdomain := getSubDomainPrefix(domain, gtldPlusOne)
			hash := sha256.Sum256([]byte(job))

			out := filepath.Clean(fmt.Sprintf("%s/%s/%s/%s/%d-%x.jpg", *screenshotsDir, firstLetter, gtldPlusOne, subdomain, currentTime.Unix(), hash))

			c := screenshot.DefaultConfig()
			c.CMD = *bin
			c.URL = job
			c.Wait = time.Duration(*wait) * time.Second
			c.Deadline = time.Duration(*deadline) * time.Second
			c.WindowWidth = 1920
			c.WindowHight = 1080
			c.FullPage = true
			c.IsJPEG = true
			c.JpegQuality = 90
			c.RandomProfileDir = false

			wrapScreenshot(&c, w, out)

		}(j)

		time.Sleep(1 * time.Nanosecond)
	}

	wg.Wait()
}
