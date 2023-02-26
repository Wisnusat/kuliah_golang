package main

import "fmt"

func main() {
	var i, j, n int
	fmt.Scan(&n)

	n += 1
	for i = 0; i < n; i++ {
		for j = 1; j < n; j++ {
			if i == j || i+j == n {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
