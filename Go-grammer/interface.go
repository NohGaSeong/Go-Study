package main

import (
	"fmt"
	"math"
)

// 인터페이스는 명명된 메서드 시그니쳐들의 모음임.
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 만약 변수가 인터페이스 타입을 가지면 명명된 인터페이스에 있는 메서드를 호출할 수 있음.
// 여기에 있는 measure 함수는 모든 geometry 위에서 동작할 수 있다는 장점이 있음.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() { 
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	
	// circle과 rect 구조체 타입은 모두 geometry 인터페이스를 구현했기 때문에
	// 이 구조체들의 인스턴스를 measure의 인자로 사용할 수 있음.
	measure(r)
	measure(c)
}
