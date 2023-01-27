package main

import (
	"at-terminal-chat-clinet/models"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func serverInfoHandler(info string) {
	infoData := strings.Split(info, "|")
	infoColor := color.New(color.FgHiWhite).Add(color.BgHiCyan)
	serverNameColor := color.New(color.FgHiWhite).Add(color.BgCyan)
	textColor := color.New(color.FgHiWhite).Add(color.BgHiBlack)

	infoColor.Println(" ====INFORMATION==== ")
	serverNameColor.Println("Server name: " + infoData[0])
	textColor.Println("Description: " + infoData[1])
	fmt.Printf("\n")
}

func displayMessageHandler(messageJsonString []byte) {
	fmt.Println("")
	var message models.Message
	err := json.Unmarshal(messageJsonString, &message)
	if err != nil {
		fmt.Printf("Cannot unmarshal message from user: %s\n", err)
	}

	nameColor := color.New(color.FgHiWhite).Add(color.BgBlue)
	textColor := color.New(color.FgHiWhite).Add(color.BgWhite)
	nameColor.Printf("[%s]:", message.From)
	textColor.Printf(" %s ", message.Message)
}
