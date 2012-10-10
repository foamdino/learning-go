// +build main2

package main

import (
	"fmt"
	"os"
	"io"
	"code.google.com/p/go.net/websocket"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}

	service := os.Args[1]
	con, err := websocket.Dial(service, "", "http://localhost")
	checkErr(err)

	var msg string
	for {
		err := websocket.Message.Receive(con, &msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Couldn't receive message: ", err.Error())
			break
		}

		fmt.Println("Received from server: ", msg)

		err = websocket.Message.Send(con, msg)
		if err != nil {
			fmt.Println("Couldn't return message")
			break
		}
	}
	os.Exit(0)
}