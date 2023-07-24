// 채널을 함수의 인자로 사용할 때, 채널이 수신용인지 송신용인지를 정할 수 있습니다.
package main

import "fmt"

// 송신용 채널만을 받음. 이 채널에 수신용 채널을 넣으면 컴파일시 에러가 발생함.
func ping(pings chan<- string, msg string){
	pings <- msg
}

// 수신용 채널을 하나 받고 송신용인 pongs를 두번째 인자로 받음.
func pong(pings <-chan string, pongs chan<- string){
	msg := <-pings
	pongs <- msg
}

func main(){
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}