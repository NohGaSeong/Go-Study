package main

import "fmt"

func main(){
	// nums 라는 배열에 2,3,4 를 담아줌
	nums := []int{2, 3, 4}
	sum := 0

	// num이라는 변수는 nums 안에 있는 값이 되고 for 돌면서 sum에 더해줌
	// 그러면 Num은 2,3,4가 되고, 이걸 각기 sum에 더해주고 println 해주면 2+3+4 = 9가 되는거임.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)
	
	// 이것도 Nums 에 있는 2,3,4 를 Num에 할당해주고
	// 1바퀴 : i는 0이고 num은 2이기 때문에 반복문 진행
	// 2바퀴 : i는 1이고 num은 3이기 때문에 if문에 걸려서 i가 출력됨(1)
	for i, num := range nums{
		if num == 3{
			fmt.Println("index:", i)
		}
	}

	// kvs라는 String dict를 생성해주고 변수를 {"a": "apple", "b":"banana"} 라는 값으로 초기화
	kvs := map[string]string{"a": "apple", "b":"banana"}

	// kvs를 돌게되는데 이렇게 2개를 두고 dict를 돌게 되면 key, value가 각 변수에 담기나보다
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// for k := range kvs 로 두게되면 Key들만 출력된다.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 문자열에서 Range를 돌게되면 유니코드 코드 포인트들을 순회한다.
	for i, c := range "go"{
		fmt.Println(i,c)
	}
}