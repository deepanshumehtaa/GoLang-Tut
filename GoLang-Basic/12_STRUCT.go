/* Struct https://youtu.be/A7lSaLWTNkM
	https://www.geeksforgeeks.org/structures-in-golang/
   Struct is a collection type, can contain various datatype undera same name, where as other store data of single types.
   
   
   Rules:
	1. Defining:
		type <name> struct {....}
		while defining you can give a new line or semicolon
		
	2. calling
		1. using var keyword
			var obj <name>
			var obj = <name>{"RAM", 25, 121004, 10000.0}
			
			Best Practices
			var obj = <name>{"RAM", 25, 121004, 10000.0}
	
	3. if the first letter of struct is:
		1. capital than it is visible all outside the package
		2. small letter, private to that package
		- similarly for the variables in the struct.
	
	4. Annonymous struct - which created and intialised during the runtime without the name
	5. s1 := s2  ==> pass by value, change in one will not reflect in another.
		
*/

package main
import "fmt"

type Person struct {
	Name string
	Age, pincode int
	Cash float64
}

func main() {
	
	s1 := Person{Name: "RAM", Age: 15, pincode: 123546, Cash: 10000.0,}
	
	// println(s1) 					// illigal for struct
	fmt.Println(s1)					// {RAM 15 123546 10000}
	fmt.Printf("%T", s1)				// main.Person{SHAAM 19 123546 11000}
					
	
	s2 := Person{"SHAAM",19, 123546, 11000.0,}	
	fmt.Println(s2)					// {SHAAM 19 123546 11000}
	fmt.Println(s2.Name)				// SHAAM
	fmt.Println(s2.Age)				// 19


	// Rule 3
	type inside_main struct{
		name string
		age, pincode int
		cash float64
	}
	s3 := inside_main{"aam",19, 123500, 110.0,}
	fmt.Println(s3)					// {aam 19 123500 110}
	fmt.Printf("%T", s3)				// main.inside_main{Deepanshu 7}


	// Rule 4
	as1 := struct{ 	name string; rollno int }{ name : "Deepanshu", rollno: 7 }
	fmt.Println(as1)				// {Deepanshu 7}
	fmt.Printf("%T", as1)				// struct { name string; rollno int }
	
	as1.name = "Mehta"
	fmt.Println(as1)				// {Mehta 7}
	
} // main ends

