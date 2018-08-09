package main

import "fmt"
import "time"

func server1(ch chan string) {
	time.Sleep(100)
	ch <- "Server 1"
}

func server2(ch chan string) {
	time.Sleep(50)
	ch <- "Server 2"
}

func main() {

	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output1:
		fmt.Println(s2)
	}
}
