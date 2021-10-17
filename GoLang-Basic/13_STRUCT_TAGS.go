/* Struct Tags 	https://youtu.be/I-yrPSj_UE0
		https://www.geeksforgeeks.org/structures-in-golang/
   Struct Tags: are the use to add meta data or extra info to structs.
		
   
   
   Rules:
	1. Defining:
		type User struct {
			Username string `tag: Tag1`
			Password int `tag: Tag2`
		}
	
	2. Any error in the tags leads to Warning and not error.
*/

package main
import "fmt"

type User struct {
	Username string `required max_len: "7"`
	Password int `valid min_len:"7"`
}

func main() {
	
	s1 := User{Username: "RAM", Password: 123546, }
	
	// println(s1) 					// illigal for struct
	fmt.Println(s1)					// {RAM 15 123546 10000}
	fmt.Printf("%T", s1)				// main.Person{SHAAM 19 123546 11000}
					
	
	s2 := User{Username: "RAM", Password: 123546, }	
	fmt.Println(s2)					// {SHAAM 19 123546 11000}


} // main ends

