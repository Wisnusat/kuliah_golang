package main

import "fmt"

func main() {
	var s1, s2, s3 int
	fmt.Scanln(&s1, &s2, &s3)
	if (s1 <= 0 || s2 <= 0 || s3 <= 0) || (s1 >= s2+s3 || s2 >= s1+s3 || s3 >= s1+s2) {
		fmt.Println("Tidak dapat membuat segitiga")
	} else if (s1==s2 && s2!=s3) || (s1==s3 && s3!=s2) || (s2==s3 && s3!=s1) {
		fmt.Println("Segitiga sama kaki")
	} else if s1==s2 && s2==s3 {
		fmt.Println("Segitiga sama sisi")
	} else if (s1*s1 == s2*s2+s3*s3) || (s2*s2==s1*s1+s3*s3) || (s3*s3==s1*s1+s2*s2) {
		fmt.Println("Segitiga siku siku")
	} else if (s1<=s2+s3) || (s2<=s1+s3) || (s3<=s1+s2) {
		fmt.Println("Segitiga bebas")
	} else {
		fmt.Println("Tidak dapat membuat segitiga")
	}
}