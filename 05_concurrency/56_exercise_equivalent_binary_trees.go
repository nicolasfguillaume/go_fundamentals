// https://tour.golang.org/concurrency/7
// https://tour.golang.org/concurrency/8

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	)

/*
Binary Tree Data Structure
A tree whose elements have *at most* 2 children is called a binary tree. 
Since each element in a binary tree can have only 2 children, 
we typically name them the left and right child.
A Binary Tree node contains following parts:
- Data
- Pointer to left child
- Pointer to right child
*/

/*
The tree package defines the type:
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

/*
Binary Tree Traversals:
(a) Inorder (Left, Root, Right)
(b) Preorder (Root, Left, Right)
(c) Postorder (Left, Right, Root)
See https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	WalkInorder(t, ch)
	// Closes the channel to indicate that no more values will be sent
	close(ch)
}

// Traverse the tree: Inorder (Left, Root, Right)
func WalkInorder(t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}
	WalkInorder(t.Left, ch)
	ch <- t.Value
	WalkInorder(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	// Implement the Same function using Walk to determine 
	// whether t1 and t2 store the same values
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		// ok is false if there are no more values to receive 
		// and the channel is closed
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		}
		
		if !ok1 {
			return true
		}
	}
}

func main() {
	/* Test the Walk function */

	// Create a new channel ch
	ch := make(chan int)

	// Kick off the walker
	// The function tree.New(k) constructs a randomly-structured 
	// (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k
	go Walk(tree.New(1), ch)

	// Then read and print 10 values from the channel
	// It should be the numbers 1, 2, 3, ..., 10
	// The loop receives values from the channel repeatedly until it is closed
	for i := range ch {
		fmt.Println(i)
	}

	/* Test the Same function */

	fmt.Println(Same(tree.New(1), tree.New(1)))  // should return true
	fmt.Println(Same(tree.New(1), tree.New(2)))  // should return false
}
