package main
import "fmt"


func makeChange(s string) string {
	s = s + " Changed"
	return s
}

func makeChangeInplace(s *string) string {
	*s = *s + " Changed"
	return *s
}

func main() {
	var myString string = "Deepanshu"
	
	ans1 := makeChange(myString)
	println("Ans: ", ans1)
	
	fmt.Println(myString)
	
	ans2 := makeChangeInplace(&myString)
	println("Ans: ", ans2)
	
	fmt.Println(myString)
	
	
} // main ends

