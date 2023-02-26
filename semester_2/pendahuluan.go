package main

import (
	"fmt"
	"math"
)

func cek_kelipatan_1301223456() {
	var x int

	fmt.Scanln(&x)
	if (x%2 == 0 && x%3 == 0) || (x%3 == 0 && x%5 == 0) {
		fmt.Println("bilangan kelipatan 2 dan 3, atau kelipatan 3 dan 5")
	} else {
		fmt.Println("BUKAN kelipatan 2 dan 3, juga BUKAN kelipatan 3 dan 5")
	}
}

func hasil_ujian_1301223456() {
	var i, n, n_passed, n_failed int
	var n1, n2, n3, avg float64

	fmt.Print("Berapa jumlah siswa yang nilainya akan diproses?")
	fmt.Scanln(&n)
	n_passed = 0
	n_failed = 0
	for i = 1; i <= n; i++ {
		fmt.Scanln(&n1, &n2, &n3)
		avg = (n1 + n2 + n3) / 3
		if avg > 80.0 {
			fmt.Println("Memenuhi syarat administratif")
			n_passed += 1
		} else {
			fmt.Println("Tidak memenuhi syarat administratif")
			n_failed += 1
		}
	}
	fmt.Println("Jumlah siswa lolos seleksi administratif", n_passed)
	fmt.Println("Jumlah siswa tidak lolos seleksi administratif", n_failed)
}

func persegi_panjang_1301223456() {
	var panjang, lebar, luas, keliling int
	var diagonal float64

	fmt.Scanln(&panjang, &lebar)
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	diagonal = math.Sqrt(float64((panjang * panjang)) + float64((lebar * lebar)))
	fmt.Println("Luas: ", luas)
	fmt.Println("Keliing: ", keliling)
	fmt.Println("Panjang diagonal: ", int(diagonal))
}

func diskon_1301223456() {
	var tahun, total_belanja, jumlah_bayar, diskon, a, c, d int

	fmt.Scanln(&tahun)
	fmt.Scanln(&total_belanja)
	a = (tahun / 1000) % 10
	c = (tahun / 10) % 10
	d = tahun % 10
	diskon = a * ((c * 10) + d)
	jumlah_bayar = total_belanja - ((total_belanja * diskon) / 100)
	fmt.Println("Besar diskon: ", diskon, "%")
	fmt.Println("Jumlah yang dibayar: ", jumlah_bayar)
}

func main() {
	cek_kelipatan_1301223456()
	hasil_ujian_1301223456()
	persegi_panjang_1301223456()
	diskon_1301223456()
}
