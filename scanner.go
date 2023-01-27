package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/AlekseiKromski/at-socket-server/core"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"os"
	"strings"
)

func Scan(conn *websocket.Conn) {
	system := color.New(color.FgHiWhite).Add(color.BgGreen)
	system.Printf("[Enter your message] => ")

	inputReader := bufio.NewReader(os.Stdin)
	message, _ := inputReader.ReadString('\n')

	clearedMessage := strings.Replace(message, string('\n'), "", 1)
	action := core.Action{
		ActionType: "send_message",
		Data:       clearedMessage,
	}
	encoded, err := json.Marshal(action)
	if err != nil {
		fmt.Println("cannot marshal action from sending message")
		return
	}
	err = conn.WriteMessage(1, encoded)
	if err != nil {
		fmt.Println("cannot send message")
		return
	}

}
