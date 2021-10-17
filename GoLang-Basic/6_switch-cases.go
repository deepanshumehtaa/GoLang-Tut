/* Switch Case

   Rules:
	1. parathesis are optional for if else
*/

package main

func main() {
	var x int = 100
	var y string = "aa"
	
	switch x{
	case 10:
		println("It is 10")
	case 50:
		println("It is 50")
	case 100:
		println("It is 100")
	case 200:
		println("It is 200")
	
	default:
		println("the default case")

	}
	
	
	
	switch y{
	case "aa":
		println("aa")
	case "bb":
		println("bb")
	default:
		println("the default case")

	}

} // main ends


