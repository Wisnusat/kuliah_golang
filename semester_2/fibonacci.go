package main

import "fmt"

func main() {
	var hasil, i, previous, next, current, n int
	previous = 0
	current = 1

	fmt.Scanln(&n)
	if n < 2 {
		hasil = 0
		fmt.Print(hasil)
	} else {
		for i = 0; i < n; i++ {
			fmt.Printf("%d ", previous)
			hasil += previous
			next = previous + current
			previous = current
			current = next
		}
		fmt.Printf("\n")
		fmt.Println(hasil)
	}
}
