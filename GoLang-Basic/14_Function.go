/* Functions 	https://youtu.be/HvsMGNCgfdo
   code that will run any number of time when call by it's name.

   Anonymous Functions: function which doesn’t contain any name. It is useful when you want to create an inline function.
		
   
   
   Rules:
	1. Defining:
		func <func name>{
		......
		}
		
	2. you can declear the function in or outside the main()
	
	3. Visibility, when front letter is:
		capital: it will become global
		small caps: will restricted to that package.
	
	4. in the function arguments the datatype is always on the right, not like python or c on left
	5. functions by default make CallByVale cal & with (*) it call by refrence.
		- Also, slices and maps are always PassByRefrence
		- array and struct are PassByValue
		
	6. PassByRefrence is fast as there is no copy of the variable.
	7. func avg(arr ...int)(<return type>){
		.....
		}
	8. In Go, func can return more than one value
		func avg(arr ...int)(<return type1>, <return type1>){
			.....
		}
	9. Anonymous Functions, Defining:
		hust remove the name after the func ... ()
		
		value := func(){
      			fmt.Println("...")
		}
	
	10. we can not have unused anony func., we can declear and call the function same time
		func(){
      			fmt.Println("...")
		}() <- the empty bracket calling it self 
		
		with Argument:
		func(i int){
      			fmt.Println("...", i)
		}(10) <- calling function with arguments 
		
	11. there is no way to provide default value to function argument, but need to use if-else
	
	12. we can assign a function to a variable, it shows functions are data type by itself.
				
*/

package main
import (
	"fmt"
)


func main() {
	
	s1 := "Deepanshu"
	s2 := "Mehta"
	
	printing("This is My first Func....", 7)		// calling printing
	get_strings(s1, s2)					// Deepanshu Mehta
	fmt.Println(s1)						// Deepanshu
				
	//...........pass by refrence.............
	get_strings_refrence(&s1, s2)				// callByRefrence with "&" sign
	fmt.Println(s1)						// DeepanshuSharma
						
	//***********Variable parameters**********
	ans := avg(1,2,3,4,5,6,7,8)
	println(ans)						// +4.500000e+000
						
						
	// Anonymous Func..................
	
	func(i int){
      	 		println("...", i)
	}(10) 							// calling function with arguments, 10
	
	f := func(s string){
		if s == "" {
			s = "Default-value"
		}
      		println(s, "hal chal.??")
	}
	
	f("kya")						// Calling the function, kya hal chal.??
	
	
}// main ends


func printing(s string, num int ){				// defining printing
	fmt.Println(s,num)
}


func get_strings(s1, s2 string){				// defining get_strings
	if s1 == "" {						// Rule 11
    		s1 = "default-value"
  	}
	s1 += s2 
	fmt.Println(s1)
}


func get_strings_refrence(s1 *string, s2 string){		// defining get_strings_refrence
	fmt.Println(s1, s2)					// 0xc00009e220 Mehta
	fmt.Println(*s1, s2)					// Deepanshu Mehta
	// fmt.Println(*s1, *s2)				// type error		
	*s1 = "Deepanshu" + "Sharma"				// this will change the original value
}



func avg(arr ...int)(float64){
	
	fmt.Printf("%T", arr)					// []int, arr is trated as Slice
	
	var sum float64 = 0					// defining avg
					
	for i :=arr[0]; i<=len(arr); i++ {
		sum += float64(i)
	}
	ans := sum/float64(len(arr))
	return ans
}


