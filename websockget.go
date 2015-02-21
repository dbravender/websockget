package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/textproto"
	"os"
	"strings"
)

func main() {
	var origin string
	var uri string
	var headers string

	flag.StringVar(&origin, "origin", "http://localhost/", "Origin")
	flag.StringVar(&headers, "headers", "", "A string of HTTP headers")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Print("Please specify a ws:// URI")
		os.Exit(2)
	}

	uri = args[0]

	config, _ := websocket.NewConfig(uri, origin)
	tp := textproto.NewReader(bufio.NewReader(strings.NewReader(headers + "\r\n\r\n")))

	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range map[string][]string(mimeHeader) {
		config.Header.Add(key, value[0])
	}

	ws, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		var msg = make([]byte, 1024)
		var n int
		for {
			if n, err = ws.Read(msg); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", msg[:n])
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if _, err := ws.Write([]byte(text)); err != nil {
			log.Fatal(err)
		}
	}
}
