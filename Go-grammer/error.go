// Go에선 명시적으로 별도의 반환값을 통해 에러를 전달하는게 일반적이고 관요엊ㄱ.
// 흔히 사용하는 try cache(except) 같은 느낌이 아니다.

// package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New = 전달된 에러 메시지로 기본적인 error 값을 생성함.
		return -1, errors.New("can't work with 42")
	}

	// 에러 위치의 nil은 에러가 없다는 것을 나타냄.
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// Error 메서드를 구현하여 커스텀 error 타입을 사용하는 것도 가능함.
// 다음은 인자 에러를 명시적으로 나타내기 위해 커스텀 타입을 사용한 위의 변형된 예제임.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 커스텀 에러 사용.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 7 , 42 이렇게 2번을 도는데
	// 7은 nil 이라 else로 가서 +3 하고 일할 수 있다, 10이 출력
	// 42는 -1 이라 if 에 걸려서 Failed하고 42가 출력
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		// 커스텀 에러에서 프로그래밍적으로 데이터를 사용하려면, 타입 단언을 통해
		// 커스텀 에러의 인스턴스로 에러를 받아와야함.
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
