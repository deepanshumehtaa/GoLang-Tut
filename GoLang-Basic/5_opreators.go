/* 3 types of opreator:
	1. Arithmatic
	2. Locgical
	3. Relational

   Rules:
	1. the ans of Arithmatic opreation by default is int, but if any number is float, then the final answer is in float.
	2. GO is smart in Logical cases. (<, >, ==, !=, >=, <=, &&)
	3. 
*/

package main
import "fmt"



func main() {
	println("Rule-1")
	fmt.Printf("div in int: %v  \n", 	  101/2)
	fmt.Printf("div in float: %v  \n", 	101/2.0)
	fmt.Printf("add-sub-mul-remainder: %v \n", 101%2)
	
	println("Logical")
	fmt.Printf("%v  \n", 2 == 2)
	fmt.Printf("%v  \n", 2 == 2.0)
	fmt.Printf("%v  \n", 2 == 2.01)
	fmt.Printf("%v  \n", 2 != -2)
	
	println("Less-than-greater-than")
	fmt.Printf("%v  \n", 2 == 2)


} // main ends

// O/P:

