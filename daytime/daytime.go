package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}

func main() {

	// localhost:1200
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("ip4", service)
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		con, err := listener.Accept()	
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		fmt.Println("Time requested: ", daytime)
		con.Write([]byte(daytime))
		con.Close()
	}
}