package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list *List[T]) display() string {
	out := ""
	iter := list
	for iter != nil {
		out += fmt.Sprintf("%v -> ", iter.val)
		iter = iter.next
	}
	out += fmt.Sprintf("nil")
	return out
}

func (list *List[T]) String() string {
	out := ""
	iter := list
	for iter != nil {
		out += fmt.Sprintf("%v -> ", iter.val)
		iter = iter.next
	}
	out += fmt.Sprintf("nil")
	return out
}

func main() {
	list := List[string]{next: nil, val: "first"}
	list.next = &List[string]{next: nil, val: "second"}
	list.next.next = &List[string]{next: nil, val: "third"}
	fmt.Printf("%v", list.display())
	fmt.Printf("%v", list)
}
