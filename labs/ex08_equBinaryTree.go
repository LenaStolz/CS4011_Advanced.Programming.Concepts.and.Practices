package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

// Test the walk function
func firstTest() {
	channel := make(chan int)
	go Walk(tree.New(1), channel)
	for i := 1; i < 11; i ++ {
		if (i != <- channel) {
			fmt.Println("Walk not possible")
			break
		}
	}
	fmt.Println("Walk works")
}

func main() {
	firstTest()
}
