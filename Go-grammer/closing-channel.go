package main

import "fmt"

func main(){

	// main() 고루틴에서 워커 고루틴으로 작업을 전달하기 위해 jobs 채널을 사용함.
	// 워커에서 돌릴 Job이 더 이상 없을 경우, Jobs채널을 Close 함.
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs
			if more{
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++{
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("send all jobs")

	<- done
}