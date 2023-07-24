package main

import "fmt"

// int 타입 파라미터를 가지므로 인자는 값으로 전달.
// zeroval은 함수를 호출할 때의 값을 ival에 복사하여 가져옴.
func zeroval(ival int){
	ival = 0
}

// int형 포인터를 받도록 *int 타입을 파라미터로
// 함수 안에서 *iptr은 메모리 주소에서 해당 주소의 현재값으로 포인터를 역참조함.
// 역참조된 포인터에 값을 할당하면 이는 참조된 주소의 값을 바꿈.

func zeroptr(iptr *int){
	*iptr = 0
}

func main(){
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// &i 는 i 의 메모리 주소를 반환함. i의 포인터임.
	fmt.Println("pointer", &i)


}
