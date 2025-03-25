package main

import (
    "fmt"
    "Go/Basics/Loops"
	"Go/Basics/Pointers"
	"Go/Basics/Functions"
	"Go/Basics/Variables"
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
}
