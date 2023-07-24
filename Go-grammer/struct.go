package main

import "fmt"

type person struct{
	name string
	age int
}

func main(){
	
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name:"Alice", age : 20})

	// 앞에 &를 붙이면 구조체 포인터를 생성할 수 있음.
	fmt.Println(&person{name: "Ann", age : 44})

	// .을 사용해 구조체 필드에 접근함.
	s := person{name: "Sean", age : 50}
	fmt.Println(s.name)

	// 구조체 포인터에도 .를 적용할 수 있음. 이때 포인터는 자동으로 역참조됨.
	sp := &s
	fmt.Println(sp.age)

	// 구조체는 수정이 가능.
	sp.age = 51
	fmt.Println(sp.age)
}