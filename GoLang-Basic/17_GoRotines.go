/* 
https://youtu.be/-Zg276gE1ZY
Concurency:
	ability of progremme, algorithm or problem divided into n number and then executed irrespective of their order or in partial
	order, without affecting the final outcome.
	Benifit:
		can achive parallelism: parallel execution of concurrent units.
		improve overall speed on multi-processor or multi-core system.
	
	Concurency achived through threads (OS Threads).
	Threads are resourse hungary.
	
GoRoutines:
	Green Threads:
		Goroutines can be thought of as light weight threads few kb, in stack size and the stack can grow and shrink
		according to needs of the application whereas in the case of threads the stack size has to be 
		specified and is fixed.
		- are not managed by OS, but by user library.
		
		- If any Goroutine in that thread blocks say waiting for user input, then another OS thread is created and the
		remaining Goroutines are moved to the new OS thread. 
		All these are taken care by the runtime and we as programmers are abstracted from these intricate details and
		are given a clean API to work with concurrency.

		- Goroutines communicate using channels. 
		Channels by design prevent race conditions from happening when accessing shared memory using Goroutines. 
		Channels can be thought of as a pipe using which Goroutines communicates.
 
		"Race Condition" : this condition occurs when a software program depends on the timing of one or 
		more processes to function correctly. If a thread runs or finishes at an unexpected time, 
		it may cause unpredictable behavior, such as incorrect output or a program deadlock.
		
		in ECE, AND has 2 i/ps if one i/p is applied and one instantly checks the o/p, without appling 2nd i/p,
		the answer will be wrong though no problem in logic but timing.
		
		in CSE, when two or more threads can access shared data and they try to change it at the same time. 
		Because the thread scheduling algorithm can swap between threads at any time, we don't know 
		the order in which the threads will attempt to access the shared data. 
		Therefore, the result of the change in data is dependent on the thread scheduling algorithm, 
		i.e. both threads are "racing" to access/change the data. 
		
		"multithreaded" : meaning they can process several threads at once.
		
		- The main function runs in its own Goroutine and its called the main Goroutine.
		
		SUPRPISES:
		
		1. When a new Goroutine is started, the goroutine call returns immediately. 
		Unlike functions, the control does not wait for the Goroutine to finish executing. 
		The control returns immediately to the next line of code after the Goroutine call 
		and any return values from the Goroutine are ignored.
		
		2. For any other Goroutines to run, the "main" Goroutine need to be running. 
		If the main Goroutine terminates then the program will be terminated and no other Goroutine has to terminate.


Threads in OS:
Process is running programme and a thread is also known as lightweight process.
lighweight =>  as this runs in a shared memory space with other threads like
No thread can exist outside a process.

we can use multiple threads for a single process and thus can use parrlelly => "Multithreading"
** Threads are executed one after another quickly switched, but gives the illusion as if they are executing in parallel.
** Threads are not independent, but processes are.

Similarity in Threads and Processes:
- Only one thread or process is active at a time
- Within process both execute sequentially
- Both can create children
	child process by fork()
	child thread by clone()

Q. Adv of Threads over Process
>  Faster communication:
	Communication between multiple threads is easier, as they shares common memory.

>  More throughput:
	Dividing into threads and each thread function is considered as one job, 
	then the number of jobs completed per unit of time is increased.

> Responsiveness: If one thread completes its execution, then its output can be immediately returned.

		
Types of Threads
- User Level Thread:
	
- Kernel Level Thread:
	
		
*/
package main

import (  
    "fmt"
)

func hello() {
	fmt.Println("Hello GOA")
}

func main() {
	go hello()
	fmt.Println("main function")
} // main ends


The o/p:
main function

Program exited.

Why??, as hello is the go function so it is a go routine and it is running concurrently/parllely with the main function, 
when the main function ends, the Goroutine running, parllely also ends.


// *************************************************************************

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
    time.Sleep(time.Millisecond * 10)  // wait for 10 milli seconds
    fmt.Println("main function")
}


The Coolest thing about is u don't need to import any MultiThread or THread Class, by by just 
using "go" keyword u can achive concurrency.

also, the goroutine execute parllely so diffrent print statements will get printed.

