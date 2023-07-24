// goroutine은 경량의 실행 스레이드이다.

// package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	// 블로킹된 호출(동기적으로 호출된 함수)의 결과가 먼저 출력된 후에, 인터리브된 
	// 두 개의 고루틴의 결과값이 출력됨.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
