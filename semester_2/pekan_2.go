package main

import "fmt"

func segitiga(a, b, c int) string {
	if a < b && a < c {
		return "bukan segitiga"
	} else {
		return "segitiga"
	}
}

func diskon(total_belanja int, isMember bool) int {
	var hasil int
	if total_belanja > 100000 && isMember {
		hasil = total_belanja - ((total_belanja * 10) / 100)
	} else if total_belanja > 100000 && !isMember {
		hasil = total_belanja - ((total_belanja * 5) / 100)
	} else {
		hasil = total_belanja
	}
	return hasil
}

func fibonacci(n int) int {
	var hasil, i, previous, next, current int
	previous = 0
	current = 1

	if n < 2 {
		hasil = 0
		return hasil
	} else {
		for i = 0; i < n; i++ {
			hasil += previous
			next = previous + current
			previous = current
			current = next
		}
		return hasil
	}
}

func faktorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * faktorial(num-1)
}

func findPermutation(x, y int) int {
	var permutation int
	if x > y {
		permutation = faktorial(x) / faktorial(x-y)
	} else if x < y {
		permutation = faktorial(y) / faktorial(y-x)
	} else {
		permutation = faktorial(x)
	}
	return permutation
}

func main() {
	// manggil function segitiga
	// var a, b, c int
	// fmt.Scanln(&a, &b, &c)
	// fmt.Println(segitiga(a, b, c))
	// -------------------------

	// manggil function diskon
	var total_belanja int
	var isMember bool

	fmt.Scanln(&total_belanja, &isMember)
	fmt.Println(diskon(total_belanja, isMember))
	// -------------------------

	// manggil function fibonacci
	// var n int

	// fmt.Scanln(&n)
	// fmt.Println(fibonacci(n))
	// -------------------------

	//manggil soal permutation
	// var x, y int

	// fmt.Scanln(&x, &y)
	// fmt.Println(faktorial(x), faktorial(y), findPermutation(x, y))
	// -------------------------
}
