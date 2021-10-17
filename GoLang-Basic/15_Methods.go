/* Methods - 2
	Methods are just like functions, but there is an object associate with them
   
   
   Rules:
	1. Method for strct as call by valuse, chanhe to "by refrence" by using *.
	2. Methods cannot be called without obj -> obj.method()
	3. we can also pass argument to the Methods.
		
*/

package main
import (
    "strconv"  // use to convert the int -> string
)

type Person struct{
	name string
	age int
}
var obj Person;


// Function Decleared
func function_ex(p Person)(string){
	return (p.name + " " + strconv.Itoa(p.age))
}


// Method Decleared
func (p Person) method_ex()(string){
	return (p.name + " " + strconv.Itoa(p.age))
}

func (p *Person) edit_detail(i int)(string){
	p.age = i
	return (p.name + " " + strconv.Itoa(p.age))
}

func main() {
	obj = Person{
		name: "Deepanshu",
		age: 23,
		}
		
	// just with regular argument
	ans1 := function_ex(obj)
	println(ans1)
	
	
	// method called with an object
	ans2 := obj.method_ex()
	println(ans2)
	
	
	// editing
	ans3 := obj.edit_detail(50)
	println(ans3)
	
	println(obj.name)			// Deepanshu
	println(obj.age)			// 50
	println(obj.method_ex)			// 0xc000068f58
	println(obj.method_ex())		// Deepanshu 50
}// main ends


