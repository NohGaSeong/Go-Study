package main

import "fmt"

func fact(n int)int {
	if n == 0 {
		return 1
	}
	// 7! : 7 * 6 * 5 * 4 * 3 * 2 * 1
	return n * fact(n-1)
} 

func main(){
	fmt.Println(fact(7))
}