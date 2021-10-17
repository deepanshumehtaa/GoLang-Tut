	/* Pointers 	https://youtu.be/HvsMGNCgfdo
			https://stackoverflow.com/questions/9320862/why-would-i-make-or-new
   store the address of the other function
		
   
   
   Rules:
	1. Defining:
		1.  
		num1 := 10
		var p1 *int = &num1
		
		-OR JUST-
		2.
		num1 := 10
		p2 := &num1
	
	2. pointers are so basic that they can be printed with just "print()"
	3. Defrencing the pointer => extracting the value of the refrence stored in that pointer.
		print(*p1)
	4. Changing the value using de-refrencing.
	5. Declearing the pointer with null aka nil btw "The default value or zero-value of a pointer is always nil"
	6. pointers as a struct type.
	7. "new", keyword
	8. GO is against pointers arthematics as it reduce the readability and cause confusions.
	9. Double Pointers: situation in which one pointer store the address of another.
		 - chain of pointers.


*/

package main
import (
	"fmt"
	"reflect"
)

type MyStruct struct{
	num int
}

func main() {
	
	num1 := 10
	var num2 int = 20
	s1 := "Deepanshu"
	
	var p1 *int = &num1
	var p2 *int = &num2
	
	
	println(&num1)				// 0xc000018050
	println(p1)				// 0xc000018050
	println(p2)				// 0xc000018058
	println(*p1)				// 10
	println(*p2)				// 20
	
	fmt.Println(p1)				// 0xc000018050
	fmt.Println(p2)				// 0xc000018058
	fmt.Println(*p1)			// 10
	fmt.Println(*p2)			// 20
			
	*p1 = 100
	println(num1)				// 100
	println(*p1)				// 100
	println(reflect.ValueOf(num1).Kind())	// 2
	
	var p3 *string = &s1			// for string
	*p3 = "Mehta"
	
	// Rule 5 
	var p4 *float64
	println(p4)				// 0x0
	// println(*p4)				// warning, invalid pointer derefrence
				
	if p4 != nil{
		println(*p4)
	} else {
		println("In Valid DeRefrence")
	}
	

	// Rule 6
	p5 := &num1
	println(*p5)
	
	obj := MyStruct{ num1 }
	p6 := &obj
	
	println(obj.num)			// 100
	println(p6.num)				// 100
	println((*p6).num)			// 100
			
			
	// NEW
	
	
	// Double Pointer
	var var1 int= 50
	var pt1 *int = &var1
	var pt2 **int = &pt1
	
	println(**pt2)				// 50
	println(*pt2)				// 0xc00009aef0, p1's valuse which is inturn address of var1
	println(pt2)				// 0xc00009af10, p1's address
	println(&pt2)				// 0xc000068f08, p2's address
	
	
	
	
}// main ends

