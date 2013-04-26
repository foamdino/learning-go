package main

import "fmt"

func main() {
     fmt.Println("reverse string")

     var sample = "wibble"
     fmt.Println(sample)
     runes := []rune(sample)

     var i = 0
     var j = len(runes)-1
     for i < j {
       runes[i], runes[j] = runes[j], runes[i]
       i++
       j--
     }
     fmt.Println(string(runes))
}
