/*
https://youtu.be/LGVRPFZr548
Deepanshu Mehta
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



func getUserName(_id int) string {
	time.Sleep(time.Millisecond * 100)
	if _id == 1 {
		return "deepanshumehtaa"
	} else {
		panic("No id found !")
	}
}

func getUserData(userName string, response_channel chan any) {
	time.Sleep(time.Millisecond * 300)
	if userName == "deepanshumehtaa" {
		// data := make(map[any]string)
		date := map[any]any{
			"id":      1,
			"Name":    "John",
			"Age":     "30",
			"Country": "USA",
		}
		response_channel <- date // pushing value into the channel

	} else {
		panic("No id found !")
	}
}

func getUserFrinedsCount(userName string, response_channel chan any) {
	time.Sleep(time.Millisecond * 100)
	if userName == "deepanshumehtaa" {
		response_channel <- 100
	} else {
		panic("No id found !")
	}
}

func main() {
	start := time.Now()

	_id := 1
	var userName string = getUserName(_id)

	response_channel := make(chan any)
	go getUserData(userName, response_channel)
	go getUserFrinedsCount(userName, response_channel)

	// value, ok := <-response_channel
	// if !ok {
	// 	fmt.Println("Channel is closed")
	// } else {
	// 	fmt.Println("Received:", value)
	// }

	// raise, error, bcoz channel dont have any value at that time
	// if panic raised of status code 2:
	// goroutine is waiting to receive a value from a channel, but no value is sent on that channel, and the channel is closed.
	for resp := range response_channel {
		fmt.Println("The Response for the channel is: ", resp)
	}
	close(response_channel)

	// fmt.Printf("User id %d and username %s having data %v and Friend Couts %d", _id, userName, data, friendCount)
	fmt.Println("Program ends in ", time.Since(start))

} // main ends


// 3rd Now its time to wait for the other workers to complete the work and at the end join the work


package main
import (
	"fmt"
	"sync"
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

func getUserData(userName string, response_channel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 900)
	if userName == "deepanshumehtaa" {
		data := map[string]interface{}{
			"id":      1,
			"Name":    "John",
			"Age":     30,
			"Country": "USA",
		}
		response_channel <- data // pushing value into the channel
		wg.Done()
	} else {
		panic("No id found !")
	}
}

func getUserFriendsCount(userName string, response_channel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 900)
	if userName == "deepanshumehtaa" {
		response_channel <- 100
		wg.Done()

	} else {
		panic("No id found !")
	}
}

func main() {
	start := time.Now()

	_id := 1
	var userName string = getUserName(_id)

	response_channel := make(chan any, 2) // channel has buffer size of 2, 2 is the capacity of the channel
	// allowing the channel to hold more than one value at a time before it blocks further sends.
	// advantages are: Avoiding Deadlock
	wg := &sync.WaitGroup{}

	wg.Add(2) // as we have two workers, now each of them needs to tell that they have done working
	go getUserData(userName, response_channel, wg)
	go getUserFriendsCount(userName, response_channel, wg)
	wg.Wait() // wait for workers to finish
    
	close(response_channel) // Close the channel after all values are sent

	// Retrieve and print values from the channel
	for resp := range response_channel {
		fmt.Println("The Response for the channel is: ", resp)
	}

	fmt.Println("Program ends in ", time.Since(start))

} // main ends

