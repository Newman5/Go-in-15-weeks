//Exercise3—Bufferedchannel(capacity1)
//Goal:ContrastwithExercise2.
//Changechannelto make(chanint,1)
//Donotreceiveimmediately
//Observe"aftersend"printswithoutwaiting
//Explainwhy(incomments)

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
		fmt.Println("after send")
	}()

	fmt.Println("before receive") // This prints first coz it does not have to wait for the goroutine to finish
	time.Sleep(1 * time.Second)
	fmt.Println(<-ch)
}
