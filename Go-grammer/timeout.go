// TimeOut은 외부 리소스를 연결하거나 실행 시간을 제한해야하는 프로그램에서 중요함.

package main

import (
	"time"
	"fmt"
)

func main() {
	
	c1 := make(chan string, 1)
	go func(){
		// 2초후에 c1 채널에 결괏값을 반환하는 외부 호출을 실행한다고 가정함.
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// res := <-c1은 결괏값을 대기하고 <-Time.After는 1초의 타임아웃 후에 전달되는 값을 대기
	// .select는 대기중인 첫번째 리시브로 진행되고 있기에 이 연산이 허용된 1초보다 더 오래 걸릴 경우 타임아웃 발생
	// 3초로 늘리면 c2로부터 수신은 성공하고 결괏값을 출력할것임.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string ,1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
	
	// 실행하게되면 첫 연산은 타임아웃, 두번째 연산은 성공하는 것을 볼 수 있음.
	// select 타임아웃 패턴을 사용하려면 채널을 통해 결괏값을 전달해야함.
}