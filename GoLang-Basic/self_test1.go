package main

import (
	"fmt"
)

func dee(str *string) string {
	return *str + "!!"
}

func main() {
	var (
		ans  string = ""
		name string = "Deepanshu"
	)


	var isPrime = func(x int) bool {
		if x%2 != 0 && x%3 != 0 && x%5 != 0 {
			return true
		}
		return false
	}

	println(isPrime(100))

	ans = dee(&name)
	fmt.Println(ans, name)

}
