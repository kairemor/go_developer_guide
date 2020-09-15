package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
*Go routine is our running process execute code syncrhone way
* if there are blocking code lke making a request a program us waiting until it finish
* the go keyword allow to run a function in an other thread go routine
 */
/*
	Go routine theory
	*One CPU Core
		* Go scheduler work with one CPU of our machine
		* only one go routine in been running in a giving time but if it block it's been pause

	*Multiple CPU Core
	* Go scheduler can mamage go routine to be run i a CPU

	* Bu default go try to use one CPU

	*Concurrency is not parallelism
		*Concurrency : We can have multiple threads executing code . If one thread block, another one is picked up and worked on
			Concurrency mean program have the ability to run multiple thing
	*parallelism:
		Multiple threads executed at the same exact time . It required multiple CPU's

	When running a program by default the main go routing is created by default and we create child go routine with
	the go keyword

	** go keywork is only in front of function call **
*/

/*
	the main routine exit automatically when finish without waiting child routine tou finish
	**** Channel are the only way to communicate between routine
	channel are typed  that mean the message send between type are the same and pre-declared

*/
func main() {
	links := []string{
		"http://google.fr",
		"http://amazon.com",
		"http://facebook.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// alternatively use
	// watch for channel when value put in l
	// function literal are anomynous function in js or lambda in python
	for l := range c {
		go func(link string) {
			time.Sleep(4 * time.Second)
			checkLink(link, c)
		}(l)
	}
	// for i := 0; i < len(links); i++ {
	// fmt.Println(<-c)
	// }
}

/*
*blocking channel
* wainting for message from channel is a blocking code when the main receive one after  it exit
 */
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error : ", link, " might be down ")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}

/*
	sending value to a channel
	chanmel <- 5

	wait for a value to be sent into the channel .when arrive get in variable
	myNumber <- channel

	we can also print it directly
	fmt.Println(<-channel )
*/
