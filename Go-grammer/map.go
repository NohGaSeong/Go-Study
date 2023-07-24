// map = go의 내장 연관 데이터타입. (다른 언어에선 hashed 또는 dicts라고 부름.)

package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 값 자체가 필요없는 경우엔 공백 식별자 _ 를 사용해서 무시할 수 있음.
	// _ 를 사용하지 않았을 경우엔 0이 나오고 _ 를 사용한 경우 false가 나옴
	// ["K1"] 으로 변경해서 진행했을때 _를 사용하지 않은 경우 7이 나오고 _를 사용한 경우엔 true가 나왔다.
	// bool 형식으로 모시꺵할때 사용한다고 생각하면 편할듯.
	
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	
	// 한 줄로 맵 선언 및 초기화 가능.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
