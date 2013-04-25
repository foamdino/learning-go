package main

// simple linked-list of ints

import "fmt"

type list struct {
     head *element
     size int
}

type element struct {
    data int
    next *element
}

func (l *list) addToTail(val int) {  
  var current = l.head
  for current.next != nil {
    //loop to end of list
    current = current.next
  }
  elem := &element{val, nil}
  current.next = elem
  l.size++
}

func (l *list) insertAtPosition(val, pos int) {
  //move to correct position
  var current = l.head
  var i = 0
  for i < pos && current.next != nil {
    current = current.next
    i++
  }

  // create a new element, set value and pointer to next
  elem := &element{val, current.next}
  current.next = elem
  l.size++
}

func (l *list) addToHead(val int) {
  elem := new(element)
  elem.data = val

  if l.head != nil {
    elem.next = l.head
  }
  
  l.head = elem
  l.size++
}

func (l *list) len() int {
  return l.size
}

func main() {
     fmt.Println("linked list example")

    l := new(list)

    l.addToHead(3)
    fmt.Printf("size: %d\n", l.size)
    l.addToHead(2)
    l.addToHead(1)
    l.insertAtPosition(4, 3)
    l.addToTail(5)
    fmt.Printf("size: %d\n", l.size)

    var current = l.head
    for current != nil {
      fmt.Printf("val: %d\n", current.data)
      current = current.next
    }
     
}
