package main

import "fmt"

func main() {

	ss := make([][]int, 2)

	for i := 0; i < len(ss); i++ {
		s := &ss[i]
		*s = append(*s, i)
	}

	fmt.Printf("%v\n", ss)

}
