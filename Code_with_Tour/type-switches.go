package main

import "fmt"

func doSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't kown about type %T!\n", v)
	}
}

func main() {
	doSomething(21)
	doSomething("hello")
	doSomething(true)
}
