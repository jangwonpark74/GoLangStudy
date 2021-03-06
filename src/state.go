package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Running = 1 << iota
	Wait
	Send
	Receive
	Sleep
)

const MAX_STATE = Sleep

func update(state *int, n int) {
	fmt.Printf("update prev =%d shift =%d ", *state, n)
	*state = (*state << n) % (Sleep | Receive | Send | Wait | Running)
	fmt.Printf("next = %d \n", *state)
}

func display(state int) {

	if state&Running == Running {
		fmt.Println("Running")
	}
	if state&Wait == Wait {
		fmt.Println("Wait")
	}

	if state&Send == Send {
		fmt.Println("Send")
	}

	if state&Receive == Receive {
		fmt.Println("Receive")
	}
	if state&Sleep == Sleep {
		fmt.Println("Sleep")
	}

}

func main() {

	state := Running
	for i := 0; i < 20; i++ {
		x := rand.Intn(6)
		display(x)
		time.Sleep(1 * time.Second)
		update(&state, x)
	}
}
