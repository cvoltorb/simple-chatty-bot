package main

import "fmt"

func main() {
	var num, rev int
	fmt.Scan(&num)

	for num != 0 {
		rev = rev*10 + num%10
		num /= 10
	}

	fmt.Println(rev)
}
