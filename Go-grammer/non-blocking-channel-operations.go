// 채널의 송수신은 기본적으로 동기적임. 그러나 select를 default문과 함께 사용하면 비동기 송수신을 구현할 수 있음.
// 비동기 다중 select도 구현 가능.

package main

import "fmt"

func main(){
	messages := make(chan string)
	signals := make(chan bool)

	
	select {
	// messages 에서 값을 사용할 수 있는 경우 select는 <-messages case에서 그 값을 가져옴.
	// 그렇지 않은 경우 default 케이스를 수행함.
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	// 다중 비동기 select를 구현하기 위해 위의 default문에 다중 case를 구현할 수 있음.
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

