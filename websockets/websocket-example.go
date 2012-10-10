// +build main1

package main

import ( 
	"fmt"
	"os"
	"net/http"
	"code.google.com/p/go.net/websocket"	
)

func echo(ws *websocket.Conn) {
	fmt.Printf("echo %#v\n", ws.Config())
	for n := 0; n < 10; n++ {
		msg := "Hello " + string(n+48)
		fmt.Println("Sending to client: " + msg)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("Cannot send")
			break
		}

		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Cannot receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}

	fmt.Println("echo finished")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	http.Handle("/", websocket.Handler(echo))
	err := http.ListenAndServe(":5678",nil)
	checkError(err)
}