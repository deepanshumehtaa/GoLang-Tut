package main

import (
	"fmt"
	"sync"
	"time"
)

func xyz(myChannel chan []int, myGroup *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	arr := []int{10, 20, 30, 40, 50}
	myChannel <- arr
	defer myGroup.Done()
}

func abc(myChannel chan []int, myGroup *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	arr := []int{1, 2, 3, 4, 5}
	myChannel <- arr
	defer myGroup.Done()
}

func dotProduct(arr1 []int, arr2 []int) []int {
	if len(arr1) == len(arr2) {
		ans_arr := make([]int, len(arr1))
		for i := 0; i < len(arr1); i++ {
			ans_arr[i] = arr1[i] * arr2[i]
		}
		return ans_arr
	} else {
		panic("The arrays are not the same length!")
	}
}

func main() {
	start := time.Now()

	// channels are nothing but Queues that follow FIFO
	myChannel := make(chan []int, 2)  // the channel has capacity of 2 and can hold element of type `array of int`
	myGroup := &sync.WaitGroup{}
	myGroup.Add(2)
	go abc(myChannel, myGroup)
	go xyz(myChannel, myGroup)
	myGroup.Wait()

	ans1 := <-myChannel
	ans2 := <-myChannel

	close(myChannel)

	f_ans := dotProduct(ans1, ans2)

	fmt.Println("answer: ", f_ans)

	fmt.Println("The Program terminates in: ", time.Since(start))
}
