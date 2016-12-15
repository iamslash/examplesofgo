package main

import "fmt"

func main() {
	i := 0
	// infinite loop
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i is: ", i)
		i++
	}
	fmt.Println("A statement just after for loop")
}
