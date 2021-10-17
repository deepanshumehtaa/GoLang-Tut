/* Slices https://youtu.be/daJ8ZDzDX4o
   Slice (GO) <--> list slicing (python)
   Slice is a variable-length sequence which stores elements of a similar type,
	not allowed to store different type of elements in the same slice.
	just like an array with index value and length, but the size of the slice is variable and not fixed.
	slices built on the top of array.
	The starting index in a slice is 0 and the last one is (length of slice – 1).
	declared just like an array, just remove the length
	
   Rules:
	1. Declaration: <slice_name> := []<data_type>{....elements...}
	2 length and capacity: capacity is the len of array from where slice is extended and length u know.
	3. arr[1:3] --> 3rd index is excluded.
	4. What will happen to the array if I change the value of an element in the slice?
		- slice is an alias(view) of array in GO
		- so change in slice -> change in array
	5. Create slice with "make()"
		- slc := make([]int, <len>, <capacity>)  --capacity is optional
		- capacity >= len
		- python print -> print() or println() has now power to show element in array or slice, but address
	6. add new element in the slice using "append()" -> add element like push() ie. at the last.
		- slice_name = append(<slice_name>, elements seprated by comma)
		
	7. appending element into the slice never raise error
		- slice capacity gets double itself automatically when u trying to add element more than the capacity
		
	8. Like array, slice can also can be nested ie. matrix
		
		
	
*/

package main
import "fmt"

func main() {


	// array
	arr := [...]int{1,2,3,4,5,}
	fmt.Println(arr)		// [1 2 3 4 5]
	
	// slice
	slc := []int{10,20,30,}		// [10 20 30]
	fmt.Println(slc)
	
	// slice of an array
	slc_arr := arr[1:4]		// [2 3 4]
	fmt.Println(slc_arr)
	
	// Types
	fmt.Printf("Type: %T", arr)  	//Type: [5]int
	fmt.Printf("Type: %T", slc_arr)	//Type: []int
	
	//Rule2
	fmt.Printf("Len: %v - Capacity %v", len(arr), cap(arr))  		// Len: 5 - Capacity 5
	fmt.Printf("Len: %v - Capacity %v", len(slc), cap(slc))			// Len: 3 - Capacity 3
	fmt.Printf("Len: %v - Capacity %v", len(slc_arr), cap(slc_arr))		// Len: 3=(4-1) , Capacity = 4(len of array - 1)
		
	// Rule4
	slc_arr1 := arr[:]		// Len: 5, Capacity 5
	slc_arr1[0] = 1000
	fmt.Println(slc_arr1)		// [1000 2 3 4 5]
	fmt.Println(arr)		// [1000 2 3 4 5]
	
	// Rule5
	slc_make := make([]int, 5)	// len and capacity are 5
	println(slc_make)		// [5/5]0xc000068e68
	fmt.Println(slc_make)		// [0 0 0 0 0]
	
	slc_make_ := make([]int, 2, 7)	// len and capacity= 2 & 7
	fmt.Println(slc_make_)		// [0 0]
		
	
	// Rule6
	slc_make_ = append(slc_make_, 6, 9, 69)		// len and capacity= 5 & 7
	fmt.Println(slc_make_)				// [0 0 6 9 69]
				
	// Rule7**
	slc_make_ = append(slc_make_, 01, 02, 03)		// len and capacity= 8 & 14
	fmt.Println(slc_make_)					// [0 0 6 9 69 1 2 3]
	
	
	// ***********************************************************
	// combining two arrays in GO
	//arr3 = arr1 + arr2				// not allowed
	
	// array
	arr1 := [...]int {1,2,3,}
	arr2 := [...]int {4,5,6, }
	
	// converting arrays into slice	
	slc_arr1, slc_arr2 := arr1[:], arr2[:]
	
	slc_arr3 := make([]int, 0)
	slc_arr3 = append(slc_arr1, slc_arr2...)
	fmt.Println(slc_arr3)				// [1 2 3 4 5 6]

	// Best way
	slc_best := append(arr1[:], arr2[:]...)
	fmt.Println(slc_best)
} // main ends



// O/P

