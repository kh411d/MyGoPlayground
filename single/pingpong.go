package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	//go player("peng", table)

	table <- new(Ball) // game on; toss the ball

	time.Sleep(1 * time.Second)
	fmt.Println(<-table) // game over; grab the ball

	//panic("show me the stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
