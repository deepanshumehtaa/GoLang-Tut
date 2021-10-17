/* ForLoop

   Rules:
	1. parathesis in for loops are not allowed, for(i :=1; i<=3; i++)
	2. NO while loop or do while loop.
	3. While loop => 1. remove the incrment from the for and paste it in the end of the body.
			 2. remove the innitial from the for and paste it in the start & outside of the for.
			 3. and no semicolon in the for defination
	4. The infine loop is
			for {
			....
			}
*/

package main

func main() {

	for i :=1; i<=3; i++ {
	print(i)
	println(" in the loop")
	}
	
	// List Slicing 
	// https://deployeveryday.com/2019/10/21/python-idioms-go.html
	words := []string {"guess", "whos", "back"}
	word, words := words[len(words)-1], words[:len(words)-1]
	println(word)
	// back
	println(words)
	
	words = append(words, "the sweet life")
	
	println("New List", words)
	// println cannot print the list
	
	
	arr := []int {1, 2, 3, 4}
	var double []int
	
	for _, number := range arr {
	double = append(double, arr*2)
	}
	println(double)
	
	
	//**********************************************************************
	// Doubling each value of the dictionary
	
	dict := map[string]int{"a": 1, "b": 2, "c": 3}
	double := make(map[string]int)

	for key, value := range dict {
    	double[key] = value * 2
	}

fmt.Println(double)

} // main ends


