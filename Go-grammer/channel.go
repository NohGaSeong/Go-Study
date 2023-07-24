// channels는 동시에 실행되고 있는 고루틴을 연결해주는 파이프임.

package main

import "fmt"

func main() {
	// string 타입을 가지는 messages 라는 채널 생성
	messages := make(chan string)

	// channel <- 구문을 사용하여 값을 채널에 전달.
	// 여기선 "ping" 이라는 문자열을 보내고 있음.
	// 프로그램을 실행하면 "Ping" 메시지가 채널을 통해 한 고루틴에서 다른 고루틴으로 성공적으로 전달됨.
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}