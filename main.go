package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type Server struct {
	Name string
	Ip   string
}

func main() {
	warning := color.New(color.FgHiWhite).Add(color.BgYellow)

	servers := []Server{
		Server{
			Name: "Public main server",
			Ip:   "at-terminal-chat-server.alekseikromski.com",
		},
	}
	info := color.New(color.FgHiWhite).Add(color.BgHiMagenta)
	listItem := color.New(color.FgHiWhite).Add(color.BgHiBlack)
	info.Println("There is some public registered servers")

	for i, server := range servers {
		listItem.Printf(" 🧨 %d. %s [%s] \n", i, server.Name, server.Ip)
	}
	listItem.Printf(" 👋 %d. %s \n", 99, "enter by a hand")

	var choose int
	warning.Printf("Your choose> ")
	fmt.Scan(&choose)

	var url_custom string
	if choose == 99 {
		warning.Printf("Enter server ip and address> ")
		fmt.Scan(&url_custom)
	} else {
		url_custom = servers[choose].Ip
	}

	server_url := url.URL{Scheme: "ws", Host: url_custom, Path: "/"}
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
