/* 
https://youtu.be/-Zg276gE1ZY
https://golangbot.com/goroutines/#:~:text=Goroutines%20are%20functions%20or%20methods,thousands%20of%20Goroutines%20running%20concurrently

GoRoutines with function and Methods:
	1.
		
*/

package main

import (  
    "fmt"
    "time"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}

func main() {  
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}


