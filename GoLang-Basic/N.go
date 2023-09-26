/*
https://youtu.be/LGVRPFZr548
*/

package main

import (
	"fmt"
	"time"
)

func getUserName(_id int) string {
    time.Sleep(time.Millisecond * 100)
	if _id == 1 {
		return "deepanshumehtaa"

	} else {
		panic("No id found !")
	}
}

func getUserData(userName string) map[any]any {
    time.Sleep(time.Millisecond * 300)
	if userName == "deepanshumehtaa" {
		// data := make(map[any]string)
		date := map[any]any{
			"id":      1,
			"Name":    "John",
			"Age":     "30",
			"Country": "USA",
		}
		return date

	} else {
		panic("No id found !")
	}
}

func getUserFrinedsCount(userName string) uint {
    time.Sleep(time.Millisecond * 100)
	if userName == "deepanshumehtaa" {
		return 100
	} else {
		panic("No id found !")
	}
}

func main() {
	start := time.Now()

	_id := 1
	var userName string = getUserName(_id)
	data := getUserData(userName)
	var friendCount uint = getUserFrinedsCount(userName)

	fmt.Printf("User id %d and username %s having data %v and Friend Couts %d", _id, userName, data, friendCount)
	fmt.Println("Program ends in ", time.Since(start))

} // main ends

// above programme took around: 509ms

// With GoRoutines .................................................................................................................




