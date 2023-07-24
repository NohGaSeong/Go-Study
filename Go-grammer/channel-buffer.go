// 기본적으로 채널은 버퍼가 없는 상태. 이는 송신된 값을 받으려고 준비하고 있는 리시브가 있는 경우에만 보낼 수 있음.
// 버퍼링된 채널은 대응되는 리시버가 없는 값을 제한된 갯수만큼 허용함.
package main

import "fmt"

func main() {
	// 최대 2개의 값을 버퍼링할 수 있는 string 채널을 생성하였음.
	messages := make(chan string, 2)

	// 채널이 생성되었기에 동시에 실행중인 대응되는 리시버 없이 채널에 값이 전달 가능.
	messages <- "buffered"
	messages <- "channel"


	// 평소와 같은 방법으로 이 두 값을 받을 수 있음.
	fmt.Println(<-messages)

	fmt.Println(<-messages)
}
