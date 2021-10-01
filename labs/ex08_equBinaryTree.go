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
func Same(t1, t2 *tree.Tree) bool{
	channel1 := make(chan int)
	channel2 := make(chan int)
	go Walk(t1, channel1)
	go Walk(t2, channel2)

	for i := 0; i < cap(channel1); i ++ {
		if(<-channel1 != <-channel2){
			return false
		}
	}
	return true
}

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
	fmt.Println("Walk possible")
}

// Test the 'same' function
func secondTest() {
	if (Same(tree.New(1), tree.New(1)) && Same(tree.New(1), tree.New(2))) {
		fmt.Println("'Same'-function does work")
	} else {
		fmt.Println("'Same'-function does not work")
	}

}

func main() {
	firstTest()
	secondTest()
}
