package main

import (
	"stack"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var stak = new(stack.Stack)


func main() {
	//loop and get tokens from stdin
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range(s) {
			switch {
			case c >= '0' && c <='9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				stak.Push(r)
				token = ""
			case c == '+':
				fmt.Printf("%d\n", stak.Pop()+stak.Pop())
			case c == '*':
				fmt.Printf("%d\n", stak.Pop()*stak.Pop())
			case c == '-':
				p := stak.Pop()
				q := stak.Pop()
				fmt.Printf("%d\n", q-p)
			case c == '/':
				p := stak.Pop()
				q := stak.Pop()
				fmt.Printf("%d\n", q/p)
			case c == 'q':
				return
			default:
				//error
			}
		}
	}
}