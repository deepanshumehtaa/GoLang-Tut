package main
import ("fmt")

func main(){

	fmt.Println("Number is:", 7)

	// array
	// var var_name [size]var_type
	var arr [10]string
 	fmt.Println(arr)
	
	arr[0] = "first"
	arr[1] = "second"
	arr[9] = "last"
	fmt.Println(arr)
	fmt.Println("Length of arr:", len(arr))
	fmt.Printf("\n")
	
	
	// slice
	// var var_name []var_type,  eg: var arr1 [] float64
	
	// -OR-
	elements := []int {3, 5, 2, 6, 2, 100}
	nums:=[]int {5,7,4,8,9}
	str := []string {"Amazing", "Brilliant", "Good One", "Excellent"}
	
	fmt.Printf("%T \n", nums)
	fmt.Println(nums)
	fmt.Println(str)
	fmt.Print("\n")
		
	
	// arr2 := make([]int, 0, 3)
 	
	
	for n := range elements {
		fmt.Println("Elements", n)
	}
	fmt.Printf("\n")
	

// for loop, https://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go/7782507#7782507

for i := 0; i < len(elements); i++ {
        fmt.Print(elements[i] + 1)
        fmt.Print(" ")
    }
    fmt.Println("... DONE!")



// List Slicing 
	// https://deployeveryday.com/2019/10/21/python-idioms-go.html
	words := []string {"guess", "whos", "back"}
	word, words := words[len(words)-1], words[:len(words)-1]
	println(word)
	// backS
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


	Links: https://www.geeksforgeeks.org/class-and-object-in-golang/

} // main ends

