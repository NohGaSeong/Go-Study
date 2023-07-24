// package main

import "fmt"

// intSeq() 함수는 내부에 익명으로 정의한 또 다른 함수를 반환
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	// nextInt에 intSeq() 를 할당하고 호출하면 i+=1 이 계속 반영됨.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 다시 할당하면 값들이 초기화되는 것 같음.
	newInts := intSeq()
	fmt.Println(newInts())

}
