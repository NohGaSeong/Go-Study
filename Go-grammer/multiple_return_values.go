package main

import "fmt"

// 이 함수가 2개의 Int를 반환한다는 것임.
func vals() (int, int){
	return 3, 7
}

func main(){
	// 다중반환을 하는 함수를 받기 위해선 다중할당으로 받아야함.
	// 다중 할당이 아니라면 cannot initialize 1 variables with 2 values 식으로 에러가 나게됨.
	a, b := vals()
	// 할당된 첫번째 순서
	fmt.Println(a)
	// 할당된 두번째 순서
	fmt.Println(b)

	// 이렇게 하나의 값만 사용할 수 있음.
	// 앞의 값을 사용하고 싶다면 c,_ 로 하면되겠죠?
	_,c := vals()
	fmt.Println(c)

}