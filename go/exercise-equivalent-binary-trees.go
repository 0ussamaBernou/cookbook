package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan<- int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Walking(t *tree.Tree, ch chan<- int) {
	defer close(ch)
	Walk(t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	go Walking(t1, ch1)

	ch2 := make(chan int, 10)
	go Walking(t2, ch2)

	for i1 := range ch1 {
		i2 := <-ch2
		if i1 != i2 {
			return false
		}
	}
	return true
}

func main() {
	// ch := make(chan int, 10)
	// go Walk(tree.New(1), ch)
	// for i := range ch {
	// 	fmt.Println(i)
	// }
	same := Same(tree.New(1), tree.New(1)) //output: true
	fmt.Println(same)
}
