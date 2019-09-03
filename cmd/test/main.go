package main

import (
	"fmt"
	"time"
)

func Run(task_id, sleeptime, timeout int) {
	ch_run := make(chan string)
	go run(task_id, sleeptime, ch_run)
	for {
		select {
		case <-ch_run:
			fmt.Sprintf("task id %d , over", task_id)
		case <-time.After(time.Duration(timeout) * time.Second):
			fmt.Sprintf("task id %d , timeout", task_id)

		}
	}

}

func run(task_id, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	return
}

func main() {
	input := []int{3, 2, 1}
	timeout := 2
	//chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		//chs[i] = make(chan string)
		go Run(i, sleeptime, timeout)
	}

	//for _, ch := range chs {
	//	fmt.Println(<-ch)
	//}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
	time.Sleep(100 * time.Second)

}
