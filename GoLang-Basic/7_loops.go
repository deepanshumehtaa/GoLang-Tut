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

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		print(i)
		println(" in the loop")
	}

	println("\n --While Loop-- \n")

	var i = 0
	for i < 10 {
		if i%2 == 0 {
			fmt.Print("while loop! ", i, "\t")
		}
		i += 1
	}

	i = 0
	for i <= 10 {
		print("while loop again...")
		i++
	}

	// the perfect loop
	for i := 1; i <= 3; i++ {
		print(i)
		println(" in the loop")
	}

	// the infinite Loop without the break
	for {
		println("The infinite Loop")
		break
	}

	// printing all the even number
	for i := 0; i < 100; i++ {
		if i%2 != 0 { continue }
		print(i, "\t")
	}

} // main ends

