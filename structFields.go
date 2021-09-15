package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v:= Vertex{2, 4}
	v.X = 8
	fmt.Println(v.X)
}

//Output:
//8
