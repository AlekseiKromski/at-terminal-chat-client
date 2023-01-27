package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {
	var url_custom string
	warning := color.New(color.FgHiWhite).Add(color.BgYellow)
	warning.Printf("Enter server ip and address> ")
	fmt.Scan(&url_custom)

	server_url := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(server_url.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			Parse(message)

			fmt.Printf("\n")
			go func() {
				for {
					Scan(c)
				}
			}()
		}
	}()

	//For looping main goroutine
	done := make(chan struct{})
	for {
		select {
		case <-done:
			return
		}
	}
}
