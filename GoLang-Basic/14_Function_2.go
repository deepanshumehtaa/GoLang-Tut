/* Functions - 2 	https://youtu.be/vdm04bVzkLg
			https://www.youtube.com/redirect?event=video_description&redir_token=QUFFLUhqbml0V0tkSVdoUVRHNTM4YTFEWHhITHQzbzQ1d3xBQ3Jtc0ttekFVNkhXLVBkU0VHZzg0a1dpcFk1X0hCbUt0UWVqOFRGYnVNSkNIRVE4RU5hOXByREdiYjB3bnZWdHk1bU4wVVYtVnFfVjlXeklpX0hUY2pRenkyUEJINUFtYnFmZ2haMU9fZUxDVVcxWkFOOWlTbw&q=https%3A%2F%2Fgithub.com%2Faedorado%2Flearning-go
   Anonymous Functions: function which doesn’t contain any name. It is useful when you want to create an inline function.
		
   
   
   Rules:
	1. function as a variable:
		- assigned as refrence, without round brackets "()"
	
	2. we can not use := outside the main function, otherwise exception
		"non-declaration statement outside function body",
		else you need to use var keyword
	
	3. we can pass and return a function also
	4. exception: "func1(i) used as value", means this func1 should return something, but return noting
				
*/

package main
import (
    "math"
)

func test(i int) (int){
	return 10*i
}


var t func (i int) (int)  		// declearing a variable with type funct with return of int and accept as int
					// t can directly take as x does but without := now use =



func nested_function1(func1 func(i int)(int)) (int){
	// this function is accepting func1
	// to call this function u should have another function ready to pass
	return func1(10)
}


func nested_function2(i int) (func()) {
	// this function is accepting int and returning the empty function
	return func() { math.Abs(float64(i)) }
}

func main() {
	
	x := test
	ans := x(10)			// assigned as refrence, without round brackets "()"
	println(ans)			// calling
	
	
	y := func () {
		println("working")
	}
	
	y()					// calling, working :: y shold be the annonymous function
				
	// calling the outer function
	// printn( t(100) )
	
	ans1 := nested_function1(test)		// test is func, passed to "nested_function"
	println(ans1)
	
	ans2 := nested_function2(-10000)
	println(ans2)
}// main ends


