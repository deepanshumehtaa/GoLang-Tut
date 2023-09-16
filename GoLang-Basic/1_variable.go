/*

"fmt: stands for the Format package."
"This package allows to format basic strings, values, or anything and print them or 
"collect user input from the console, or write into a file using a writer or even print customized fancy error messages. This package is all about formatting input and output"

" There is Nothing called CHAR in GO, but rune which is stored as ASCII in the memory address and not as character "
*/

package main
import (
	"fmt"
	"reflect"
)


var global_variable int = 500

func main() {

//********************************Declaration***********************************

var i = 42
fmt.Print(i)
fmt.Println("") //----------------------------------------

var ii int
ii = 10
fmt.Printf("%v", ii) // same as--> fmt.Printf("%d", ii)

fmt.Println("") //-----------------------------------------

var iii float32 = 101.42912
fmt.Printf("%.1f", iii)
fmt.Println("") //-----------------------------------------

//********************************Short Hand***********************************

a := 50
b := 'Z'
c := 45.65
d := "THE"
e := true
x,y,z := -10, 0, 10

var(	var1 = 50
	var2 = 25.22
	var3 string = "Telefonía"
	var4 bool  // default value is false
)

fmt.Printf("\n The Type is: %T", a)
fmt.Printf("\n The Value is: %v", a)

fmt.Printf("\n The Type is: %T", b)
fmt.Printf("\n The Value is: %v", b)

fmt.Printf("\n The Type is: %T", c)
fmt.Printf("\n The Value is: %v", c)

fmt.Printf("\n The Type is: %T", d)
fmt.Printf("\n The Value is: %v", d)

fmt.Printf("\n The Type is: %T", e)
fmt.Printf("\n The Value is: %v", e)

fmt.Printf("\n The Type is: %T %T %T", x,y,z)
fmt.Printf("\n The Value is: %v %v %v", x,y,z)

fmt.Printf("\n <<<<>>>>> %v %v %v %v \n", var1, var2, var3, var4)


// automatic assigning datatype with var without using int, bool, etc
var au1 = 2*2 == 4 || 10%3 == 0 // bool
var au2 = "The Long String"     // string
var au3 = 'a'                   // rune, 97
var au4 = 0                     // int
var au5 = 0.0                   // float64
var au6 = 5.0 + 50              // float64
var au7 = true                  // bool
var au8 = false && true         // false
var au9 = 5 + 7i                // complex
var au10 = ""                	// string

println(">>>>>>>>>>", au1, au2, au3, au4)
println(au5, au6, au7)
fmt.Println(au8, au9, au10, reflect.TypeOf(au4))

//********************************Local Global***********************************

println(global_variable)
global_variable = global_variable - 100
println(global_variable)


//********************************if-else****************************************

var bool_var bool


if global_variable == 400{
        fmt.Println("Subtraction")
}else{
	fmt.Println("No change")
}



if bool_var{
        fmt.Println("True")
} else {
	fmt.Println("False")
}
	

// the problem with unsigned int is of comparisons, it can't be compared with any other datatype
var ui8 uint8 = 20      // uint8 can only store number b/w 0 to 255
var ui16 uint16 = 1000  // ui16 is 0 to 65535
var ui64 uint64 = 99000 // very large range

println(ui8, ui16, ui64)  // 20 1000 99000

// Error unable to compare the values
// if ui64 == ui8 {
//  println("Hello Daddy!")
// }


// CHAR, rune and string ................................................

var str_var string
var int_var int
var float64_var float64
var float32_var float32
var boool_var bool
var byte_var byte
var rune_var rune  // alias for 32-bit int.
myMap := map[string]float32{}  // {"Golang": 85, "Java": 80, "Python":81}
var myMap1 map[string]float32

fmt.Println(str_var, int_var, float64_var, float32_var, boool_var, byte_var, rune_var, myMap, myMap1)  //  0 0 0 false 0 0 map[] map[]



// maps
myMap1 = map[string]float32{
    "Golang": 85.0,
    "Java":   80.0,
    "Python": 81.0,
}
myMap1["Maths"] = 99.9
fmt.Println(myMap1)  // map[Golang:85 Java:80 Maths:99.9 Python:81]
	
} // end main


// 42
// 10
// 101.4

//  The Type is: int
//  The Value is: 50
//  The Type is: int32
//  The Value is: 90
//  The Type is: float64
//  The Value is: 45.65
//  The Type is: string
//  The Value is: THE
//  The Type is: bool
//  The Value is: true
//  The Type is: int int int
//  The Value is: -10 0 10
//  <<<<>>>>> 50 25.22 Telefonía false 
// true The Long String 115 5
// +5.000000e+000 +5.500000e+001 true
// false (+5.000000e+000+7.000000e+000i)
// 500
// 400
// Subtraction
// False
// 20 1000 99000
// Motto 77 77
// 0 0 0 false 0 0
