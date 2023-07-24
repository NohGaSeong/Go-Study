// package main

import "fmt"

// int형 인자를 임의의 갯수만큼 받겠다는거임.
func sum(nums ...int){
	fmt.Print(nums, " ")
	total := 0

	// 임의로 받은 수들을 num에 대입하고 total 에 더한다는 것.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


func main(){
	// 1 2 가 Nums 에 들어가고 이걸 더한 값이 출력됨.
	sum(1,2)
	sum(1, 2, 3)

	// 만약 값들이 이렇게 슬라이스에 들어가있다면 
	nums := []int{1, 2, 3, 4}
	// func(slice...) 를 이용해서 가변 함수의 인자체 적용시키면됨.
	sum(nums...)
}