// package main

import "fmt"

type rect struct {
	width, height int
}

// area method는 *rect 를 receiver type으로 가짐.
// 메서드는 포인터 또는 값을 리시버 타입으로 하여 정의될 수 있음.
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	// 구조체에 정의된 두 개의 메서드를 호출
	r := rect{width: 10, height: 5}

	// Go는 메서드의 호출에 대해 갑소가 포인터간의 변환을 자동으로 처리.
	// 메서드 호출 시 값 복사를 피하기 위해 포인터 리시버 타입을 사용할 수도 있고
	// 전달되는 구조체를 변경할 수 있도록 할 수도 있음.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
