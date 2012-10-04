package main

import (
	"fmt"
	)

//empty type
type e interface{}

func mapper(f func(e) e, l []e) []e {
	for i:=0; i<len(l); i++ {
		l[i] = f(l[i])
	}
	return l
}

func inc(n e) e {
	switch n.(type) {
	case int:
		return n.(int) + 1
	case string:
		return n.(string) + n.(string)
	case rune:
		return n.(rune) + 1
	}
	return n
}

func main() {
	test_ints := []e{1,2,3,4,5}
	fmt.Printf("before map: %v\n", test_ints)
	mapper(inc, test_ints)
	fmt.Printf("after map: %v\n", test_ints)
	test_strings := []e{"a","b","c"}
	fmt.Printf("before map: %v\n", test_strings)
	mapper(inc, test_strings)
	fmt.Printf("after map: %v\n", test_strings)
	test_runes := []e{'a','b','c'}
	fmt.Printf("before map: %v\n", test_runes)
	mapper(inc, test_runes)
	fmt.Printf("after map: %v\n", test_runes)
}