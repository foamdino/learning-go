package main

import "fmt"

type stak struct {
	i int
	data []int
}


func push(v int, ps *stak) {
	if ps.i == len(ps.data) {
		panic("Stackoverflow")
	} else {
		ps.data[ps.i] = v
		ps.i++
	}
}

func pop(ps *stak) (v int) {
	if ps.i == 0 {
		panic("Stackunderflow")
	} else {
		ps.i--
		v = ps.data[ps.i]
		ps.data[ps.i] = 0
	}
	return
}

func (s stak) String() (out string) {
	out = ""
	for pos, val := range(s.data) {
		out = out + fmt.Sprintf("%d => [%d] ", pos, val)
	}
	out = out +"\n"
	return
}

func main() {
	fmt.Println("Stack test")
	test_stack := &stak{0, []int{0,0,0,0,0,0,0,0}}
	fmt.Printf("stack: %v", test_stack)
	fmt.Println("Pushing 2 onto stack")
	push(2, test_stack)
	fmt.Printf("stack: %v", test_stack)
	fmt.Printf("num ints=%d\n", test_stack.i)
	fmt.Println("Popping 2 off stack")
	fmt.Printf("v=%d\n", pop(test_stack))
	fmt.Printf("stack: %v", test_stack)
	fmt.Printf("num ints=%d\n", test_stack.i)
}