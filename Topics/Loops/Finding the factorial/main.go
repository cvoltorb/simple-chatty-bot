package main

import "fmt"

func main() {
	var num, res int
	fmt.Scan(&num)
	res = 1

	for i := 1; i <= num; i++ {
		res *= i
	}

	fmt.Println(res)
}
