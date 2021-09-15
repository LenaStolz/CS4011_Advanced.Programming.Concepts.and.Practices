package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2000

	p := &i				// point to i
	fmt.Println(*p)		// read i through the pointer
	*p = 21				// set i through the pointer
	fmt.Println(i)		// see the new value of i

	p = &j				// point to j
	*p = *p / 4			// divide j through the pointer
	fmt.Println(j)		// see the new value of j
}

//Output:
//42
//21
//500
