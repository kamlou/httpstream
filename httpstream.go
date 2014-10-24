package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}

	url := os.Getenv("URL")
	if url == "" {
		url = "http://localhost:3000/"
	}

	resp, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Proto, resp.Status)

	for k, v := range resp.Header {
		for _, v2 := range v {
			fmt.Printf("%s: %s\n", k, v2)
		}
	}

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		fmt.Println("--> ", string(line))
	}
}
