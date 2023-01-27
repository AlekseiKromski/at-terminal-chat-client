package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlekseiKromski/at-socket-server/client"
)

func Parse(serverMessage []byte) {
	//unmarshal server message
	var decodedServerResponse client.AtServerResponse
	err := json.Unmarshal(serverMessage, &decodedServerResponse)
	if err != nil {
		fmt.Printf("cannot parse server message: %s", err)
		return
	}

	switch {
	case decodedServerResponse.ClientActionType == "server_info":
		serverInfoHandler(decodedServerResponse.Data)
	case decodedServerResponse.ClientActionType == "get_message":
		displayMessageHandler([]byte(decodedServerResponse.Data))
	default:
		fmt.Printf("cannot define handler")
	}
}
