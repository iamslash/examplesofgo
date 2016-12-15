package main

import "fmt"

func test(value int) {
	switch value {
	case 1:
		fmt.Println("One")
		fallthrough
	case 0:
		fmt.Println("Zero")
		break
	}
}

func main() {
	test(1)
}
