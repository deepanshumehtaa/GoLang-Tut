/* constant: variable which do not change it's value through the course of proramme
   
   Rules:
	1. Acc. to naming convention of GO, if anything in uppercase ==> an exported entity  or can be use across the whole project.
	2. camel case or pascal case allowed for variables.
	3. Go will not care about the usage of constants in the programme.
	4. tring to change the value of constant results in error.
	5. Constant can only be initialize on compile time and not on the runtime.
	6. we can create a constant without specifing the type
*/



package main
import "fmt"
import "math"

const goCon1 int = 6		// available only for main
const GoCon2 float32 = 12.5	// available ouside the file
const goCon3 float64 = 12.5*6   // get initialized at compile time
const goCon4 float64 = math.sqrt(100)   // error, cannot initialized on runtime

const goCon5 = 2000   // without datatype **BEST**

func main() {
	println("Constants")
	fmt.Printf("Value: %v and Type: %T\n", goCon1, goCon1)
	fmt.Printf("Value: %v and Type: %T\n", GoCon2, GoCon2)
	fmt.Printf("Value: %v and Type: %T\n", goCon3, goCon3)
	
	const goCon3 int = 100	
	fmt.Printf("Value: %v and Type: %T\n", goCon3, goCon3)
	
	goCon3 = 100 // cause error cannot assign new value to constants
	
	
	//**********constant without datatype can be use with anyone************
	
	x := goCon5 + 100
	println(x)
	
	y := goCon5 + 10.5
	println(y)
	
	z := goCon5 + "xxx"  // error because goCon5 belongs to digit class and not string
	println(z)


} // main ends



O/P:

Constants
Value: 6 and Type: int
Value: 12.5 and Type: float32
Value: 75 and Type: float64
Value: 100 and Type: int

2100
+2.010500e+003

