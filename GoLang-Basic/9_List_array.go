/* Arrays https://medium.com/@marty.stepien/arrays-vs-slices-bonanza-in-golang-fa8d32cd2b7c
	https://www.educative.io/edpresso/what-is-the-for-range-loop-in-golang?aid=5082902844932096&utm_source=google&utm_medium=cpc&utm_campaign=edpresso-dynamic&gclid=CjwKCAjwg4-EBhBwEiwAzYAlssYqBcD0eFelouYnzPJBRBQ0xdPfy5GlfdyjhGsLlSzNCNYDp26xtBoCyssQAvD_BwE

   Rules:
	1. Declaration: <array_name> := [size]<data_type>{....elements...}
	2. Declaration also done with three dots called "elipses", three dots allow any size of array --> vector.
	3. we can access the array by index, and start with 0
	4. we can append variable into the array.
	5. Empty array declaration with no use raise error.
	6. you can't declear matrix with elipses.
	7. arr1 := arr2, this will create a new copy and not alising like python.
	8. The for statement supports one additional form that uses the keyword 
	   "range" to iterate over an expression that evaluates to an array, slice, map, string, or channel
*/

package main
import "fmt"

func main() {


	// Rule1
	prices1 := [3]int{1,2,3}
	fmt.Println(prices1)
	
	// Rule2 
	prices2 := [...]int{10,20,30,40,50}
	fmt.Println(prices2)
	
	fmt.Println(prices1[0] + prices2[1])
	fmt.Println("Length of prices2:", len(prices2))
	
	// Rule3 
	some_variable := 100
	arr := [...]int{10,20,30,some_variable}
	fmt.Println(arr)
	
	// Best way to delclear array
	// Rule5
	// var empty_arr [3]string
	
	var arr_1 [3]string
	arr_1[0] = "a"
	fmt.Println(arr_1)
	
	
	//..............MATRIX..................
	
	var matrix [3][3]string
	fmt.Println(matrix)
	
	var int_matrix [3][3]int
	int_matrix[2][2] = 10
	int_matrix[0] = [3]int{1,2,3}  	//filling full row
					// here the int_matrix[0] acting as individual array
	
	fmt.Println(int_matrix)
	
	
	// with loops.............................................
	
	
	for i := 0; i < len(int_matrix); i++ {
		fmt.Println(int_matrix[i])
	}
	
	println("with double for loop")
	for p:= 0; p<2; p++{
	for q:= 0; q<2; q++{
		fmt.Println(int_matrix[p][q])
  
	}}
	
	// printing list in dictionary style
	
	// printing only indexes
	for v := range arr{
		print(v)
	}
	
	println()
	// printing index and values
	for i, v := range arr{
		println(i,"values: ",v)
	}
	
	
	j := 0
	for range arr {
		fmt.Println(arr[j])
		j++
	}
	
	a := [3][4]int{
    		{0, 1, 2, 3} , // initializers for row indexed by 0
    		{4, 5, 6, 7} , // initializers for row indexed by 1
    		{8, 9, 10, 11} // initializers for row indexed by 2
	}
	
	// ****itreating over a string*****
	
	for _, ch:= range "Deepanshu Mehta" {
		println(ch)
		// j++
	}
	
	for _, ch:= range "Deepanshu Mehta" {
		fmt.Println("%U ", ch)
		// j++
	}

} // main ends



// O/P

[1 2 3]
[10 20 30 40 50]
21

Length of prices2: 5

[10 20 30 100]
[a  ]
[[  ] [  ] [  ]]
[[1 2 3] [0 0 0] [0 0 10]]

[1 2 3]
[0 0 0]
[0 0 10]

with double for loop
1
2
0
0

0123
0 values:  10
1 values:  20
2 values:  30
3 values:  100

10
20
30
100

68
101
101
112
97
110
115
104
117
32
77
101
104
116
97

%U  68
%U  101
%U  101
%U  112
%U  97
%U  110
%U  115
%U  104
%U  117
%U  32
%U  77
%U  101
%U  104
%U  116
%U  97
