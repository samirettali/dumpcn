package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func worker(jobChan <-chan string, resChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	var client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	for job := range jobChan {
		if !strings.HasPrefix(job, "https://") {
			job = "https://" + job
		} else if strings.HasPrefix(job, "http://") {
			job = strings.Replace(job, "http://", "https://", 1)
		}

		req, reqErr := http.NewRequest("HEAD", job, nil)
		if reqErr != nil {
			continue
		}

		resp, clientErr := client.Do(req)
		if clientErr != nil {
			continue
		}

		if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
			resChan <- resp.TLS.PeerCertificates[0].Subject.CommonName
		}
	}

}
func main() {
	workers := flag.Int("t", 32, "numbers of threads")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	jobChan := make(chan string)
	resChan := make(chan string)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(*workers)

	go func() {
		wg.Wait()
		close(done)
	}()

	for i := 0; i < *workers; i++ {
		go worker(jobChan, resChan, &wg)
	}

	go func() {
		for scanner.Scan() {
			jobChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
		close(jobChan)
	}()

	for {
		select {
		case <-done:
			return
		case res := <-resChan:
			fmt.Println(res)
		}
	}
}
