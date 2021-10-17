"fmt: stands for the Format package."
"This package allows to format basic strings, values, or anything and print them or collect user input from the console, or write into a file using a writer or even print customized fancy error messages. This package is all about formatting input and output"


package main
import "fmt"


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
	var4 bool
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
}else{
	fmt.Println("False")
}


} // end main

