package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Printf("a")
	fmt.Print("b")
	fmt.Println("c") 		// next line will be as new line
	fmt.Println("a", 007 , "z")	// support all datatypes
	fmt.Println("Deepan" + "shu" + "_07")

	fmt.Printf("1") // doesn't support direct print other than string
	fmt.Print(2)
	fmt.Println(3.155)
	
	cati := "Hello, " + "I am Deepanshu"
	fmt.Println(cati)

	//---------------------without FMT------------------------------------------

	println("without FMT")
	print(10000)

	k := 600
	println(k)

	// you cant use printf without fmt.

	// with variables
	items := []int {10, 20, 50, 100}

	println(items)       // print the address
	fmt.Println(items)   // print the values
	
	
	//---------------------Sprintf and Sprintln------------------------------------------
	
	value0 := 2;
	value1 := "trees";
	value2 := "the park";
	
	result := fmt.Sprintf("I saw %v %v in %v.\n", value0, value1, value2);
	println(result)		// print in RED colour
	fmt.Println(result)
	
	
	result0 := fmt.Sprintln("How many", 100, "or", value0, true)  // return the concated string
	fmt.Println("[" + result0 + "]")

	fmt.Print(goCon1, goCon2, reflect.TypeOf(goCon2))
	
	/*
	
	1. Print and Println are the same the only diffrence is of new line.
	2. if you are printing without fmt, then it shows in RED colour in terminal.
	3. Sprintln is the most beautiful function of fmt it cannot print anything by itself, but return thr processed value,
	with anything as input.
	
	*/

} // main ends

// O/P:
// without FMT
// 10000600
// [4/4]0xc0000a2000
// [10 20 50 100]
// I saw 2 trees in the park.

// I saw 2 trees in the park.

// [How many 100 or 2 true
// ]
