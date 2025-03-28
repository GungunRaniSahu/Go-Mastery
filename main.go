package main

import (
    "fmt"
    "Go/Basics/Loops"
	"Go/Basics/Pointers"
	"Go/Basics/Functions"
	"Go/Basics/Variables"
	"Go/DSA/Graph"
	"Go/DSA/Linked-List"
	"Go/DSA/Searching"
	"Go/DSA/Queue"
	"Go/DSA/Sorting"
	"Go/DSA/Stack"
	"Go/DSA/Trees"
	"Go/Competitive"
)

func main() {
    fmt.Println("Calling loop functions...")
    loops.ForAsWhile()
	loops.For()
	loops.ForInRange()
	//loops.Infinite()
	loops.Nested()
	
	fmt.Println("Calling Pointer functions...")
	pointers.Array()
	pointers.Declare()
	pointers.Nil()
	pointers.Pointers()
	pointers.Ptr()
	pointers.Reference()
	pointers.Value()

	fmt.Println("Calling Functions...")
	functions.Func()
	functions.Lambda()
	functions.Multiple()
	functions.Name()
	functions.Para()
	functions.Recursion()
	functions.Return()
	functions.Variadic()

	fmt.Println("Calling Variables...")
	variables.Const()
	variables.Multiples()
	variables.Var()
	variables.Variables()

	fmt.Println("Calling Graph...")
	graph.BFS()
	graph.Dfs()

	fmt.Println("Calling Linked-list...")
	list.DoublyList()
	list.SingleList()

	fmt.Println("Calling Queue...")
	queue.Queues()

	fmt.Println("Calling Searching...")
	searching.Binary()
	searching.Linear()

	fmt.Println("Calling Sorting...")
	sorting.Bubble()
	sorting.MergeSort()
	sorting.Quick()
	
	fmt.Println("Calling Stack...")
	stack.Stacks()

	fmt.Println("Calling Trees...")
	tree.BinaryTrees()
	tree.BinarySearchTree()	

	fmt.Println("Calling Two Sum...")
	array.Two_sum()
	array.Rotate()

}
