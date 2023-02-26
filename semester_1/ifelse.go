package main

import "fmt"

func konsonan() {
	var huruf string
	fmt.Scanln(&huruf)
	if huruf == "A" || huruf == "I" || huruf == "E" || huruf == "O" || huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Println("bukan konsonan")
	} else {
		fmt.Println("konsonan")
	}
}

// func kelipatan() {
// 	var bil int
// 	fmt.Scanln(&bil)
// 	if bil%3 == 0 && bil%5 == 0 {
// 		fmt.Println("Kelipatan 3")
// 		fmt.Println("Kelipatan 5")
// 	} else if bil%3 == 0 {
// 		fmt.Println("Kelipatan 3")
// 	} else if bil%5 == 0 {
// 		fmt.Println("Kelipatan 5")
// 	}
// }

func segitiga() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == c && a != b && b != c {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}

func mutlak() {
	var x int
	fmt.Scanln(&x)
	if x > 0 {
		fmt.Println(x)
	} else if x < 0 {
		x = x * -1
		fmt.Println(x)
	} else {
		fmt.Println(0)
	}
}

func temperatur() {
	var a1, a2, a3, a4, a5 float64
	fmt.Scanln(&a1, &a2, &a3, &a4, &a5)
	if a1 > a2 && a2 > a3 && a3 > a4 && a4 > a5 {
		fmt.Println("Stabil turun")
	} else if a1 < a2 && a2 < a3 && a3 < a4 && a4 < a5 {
		fmt.Println("Stabil naik")
	} else {
		fmt.Println("Tidak stabil")
	}
}

func profit() {
	var first, second float64
	fmt.Scanln(&first, &second)
	if first < second {
		fmt.Println("Peningkatan sebesar", second-first)
	} else if first > second {
		fmt.Println("Penurunan sebesar", first-second)
	} else {
		fmt.Println("Tetap")
	}
}

func sepak_bola() {
	var g1, g2, g3, g4, terbesar, terkecil int
	fmt.Scanln(&g1, &g2, &g3, &g4)
	if g1 > g2 && g1 > g3 && g1 > g4 {
		terbesar = g1
	} else if g2 > g1 && g2 > g3 && g2 > g4 {
		terbesar = g2
	} else if g3 > g1 && g3 > g2 && g3 > g4 {
		terbesar = g3
	} else if g4 > g1 && g4 > g2 && g4 > g3 {
		terbesar = g4
	} else {
		terbesar = g1
	}

	if g1 < g2 && g1 < g3 && g1 < g4 {
		terkecil = g1
	} else if g2 < g1 && g2 < g3 && g2 < g4 {
		terkecil = g2
	} else if g3 < g1 && g3 < g2 && g3 < g4 {
		terkecil = g3
	} else if g4 < g1 && g4 < g2 && g4 < g3 {
		terkecil = g4
	} else {
		terkecil = g1
	}
	fmt.Println(terbesar, terkecil)
}

func parkir() {
	var h1, m1, h2, m2, h3, m3 int
	fmt.Scanln(&h1, &m1, &h2, &m2)
	if h2 <= 5 {
		h2 += 12
	}
	if m2 > m1 {
		m3 = m2 - m1
	} else {
		m3 = m1 - m2
	}
	if h2 > h1 {
		h3 = h2 - h1
	} else {
		h3 = h1 - h2
	}
	fmt.Println(h3, "jam", m3, "menit")
}

func akhir_tahun() {
	var total, diskon int
	var isCard bool
	fmt.Scanln(&total, &isCard)
	if total >= 200000 && isCard {
		diskon = total * 10 / 100
		fmt.Println("Kartu?", isCard)
		fmt.Println("Diskon?", true)
		fmt.Println("Cashback?", true)
		fmt.Println("Rp.", total-diskon-75000)
	} else if total >= 100000 {
		diskon = total * 10 / 100
		fmt.Println("Kartu?", isCard)
		fmt.Println("Diskon?", true)
		fmt.Println("Cashback?", false)
		fmt.Println("Rp.", total-diskon)
	} else {
		fmt.Println("Kartu?", isCard)
		fmt.Println("Diskon?", false)
		fmt.Println("Cashback?", false)
		fmt.Println("Rp.", total)
	}
}

func avatar() {
	var p, appaBesar, appaKecil, dewasa, kecil, sisa int
	fmt.Scanln(&p)
	appaBesar = 3
	appaKecil = 5
	for p > 0 {
		if appaBesar != 0 {
			p -= 5
			dewasa += 1
			appaBesar -= 1
		} else if appaKecil != 0 {
			p -= 2
			kecil += 1
			appaKecil -= 1
		} else {
			sisa = p
			p = 0
		}
	}
	if kecil == 0 && sisa == 0 {
		fmt.Println("Dewasa", dewasa)
	} else if kecil != 0 && sisa == 0 {
		fmt.Println("Dewasa", dewasa, ",", "Kecil", kecil)
	} else {
		fmt.Println("Dewasa", dewasa, ",", "Kecil", kecil, ",", "dan", sisa, "tak berangkat")
	}
}

// Tugas rumah 3 percabangan
func bil_terkecil() {
	var a, b, c, small int
	fmt.Scanln(&a, &b, &c)
	if a < b && a < c {
		small = a
	} else if b < a && b < c {
		small = b
	} else if c < a && c < b {
		small = c
	}
	fmt.Println(small)
}

func jalan_bojongsoang() {
	var hujan string
	fmt.Scanln(&hujan)
	if hujan == "tinggi" {
		fmt.Println("macet")
	} else {
		fmt.Println("tidak macet")
	}
}

func denda_buku() {
	var hari int
	fmt.Scanln(&hari)
	if hari <= 10 {
		fmt.Println(hari * 2500)
	} else {
		fmt.Println(hari * 5000)
	}
}

func urutan_bil() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	if a < b && a < c {
		if b < c {
			fmt.Println(a, b, c)
		} else {
			fmt.Println(a, c, b)
		}
	} else if b < a && b < c {
		if a < c {
			fmt.Println(b, a, c)
		} else {
			fmt.Println(b, c, a)
		}
	} else {
		if b < a {
			fmt.Println(c, b, a)
		} else {
			fmt.Println(c, a, b)
		}
	}
}

func peta_tetangga() {
	var one, two string
	fmt.Scanln(&one, &two)
	if one == "A" && two == "B" || one == "B" && two == "A" {
		fmt.Println(true)
	} else if one == "B" && two == "D" || one == "D" && two == "B" {
		fmt.Println(true)
	} else if one == "B" && two == "C" || one == "C" && two == "B" {
		fmt.Println(true)
	} else if one == "E" && two == "C" || one == "C" && two == "E" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func titik_potong() {
	var a, b, c, diskriminan int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	diskriminan = (b * b) - (4 * a * c)
	if diskriminan >= 0 {
		fmt.Println("memiliki titik potong dengan sumbu-x")
	} else {
		fmt.Println("tidak memiliki titik potong dengan sumbu-x")
	}
}

// -----------------------------------

// tugas sekolah 3
func sertifikat() {
	var p, v, ga, fs, ci, tc, jumlah float64
	fmt.Scanln(&p, &v, &ga, &fs, &ci, &tc)
	jumlah = p + v + ga + fs + ci + tc
	if jumlah/6.0 <= 3.33 {
		fmt.Println("tidak mendapat sertifikat")
	} else {
		fmt.Println("mendapat sertifikat")
	}
}

func kelipatan() {
	var a, b int
	fmt.Scanln(&a, &b)
	if a%b == 0 {
		fmt.Println("Ya,", a, "kelipatan", b)
	} else {
		fmt.Println("Tidak,", a, "bukan kelipatan", b)
	}
}

func swalayan() {
	var a, b, c, d, e, jumlahc, jumlahd, jumlahe int
	fmt.Scanln(&a, &b, &c, &d, &e)
	jumlahc = a + b + c
	jumlahd = c + d + e
	jumlahe = b + c + d
	if jumlahc == jumlahd && jumlahe%(e+a) == 0 {
		fmt.Println("cashback")
		fmt.Println("diskon")
	} else if jumlahc == jumlahd {
		fmt.Println("cashback")
	} else if jumlahe%(e+a) == 0 {
		fmt.Println("diskon")
	}
}

func kapal() {
	var kapal string
	var x, y float64
	fmt.Scanln(&kapal)
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	if kapal == "tempur" && x+y >= 5.0 {
		fmt.Println("tembak")
	} else {
		fmt.Println("tidak tembak")
	}
}

func bola() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if a > b {
		fmt.Println("menang")
	} else if b > a {
		fmt.Println("kalah")
	} else {
		fmt.Println("draw")
	}
}

// -------------------------

func parcel() {
	var berat, biaya, tambahan int
	biaya = 10000 //per kilo
	fmt.Scanln(&berat)
	if berat >= 10000 { //gram
		berat = berat / 1000
		fmt.Println(berat * biaya)
	} else {
		tambahan = berat % 1000
		if tambahan >= 500 {
			berat = berat / 1000
			fmt.Println((berat * biaya) + (tambahan * 5))
		} else {
			berat = berat / 1000
			fmt.Println((berat * biaya) + (tambahan * 15))
		}
	}
}

func gaji() {
	// kamus
	var jabatan string
	var masaKerja, jumlahAnak, gaji, tunjangan int

	// input data
	fmt.Print("Masukkan Jabatan: ")
	fmt.Scanln(&jabatan)
	fmt.Print("Masukkan Masa Kerja: ")
	fmt.Scanln(&masaKerja)
	fmt.Print("Masukkan Jumlah Anak: ")
	fmt.Scanln(&jumlahAnak)

	// percabangan
	if jabatan == "staf" || jabatan == "Staf" {
		if masaKerja < 5 {
			gaji = 4000
			tunjangan = 0
			fmt.Println(gaji + tunjangan)
		} else if masaKerja > 10 {
			gaji = 5000
			tunjangan = 100
			if jumlahAnak <= 3 {
				fmt.Println(gaji + (tunjangan * jumlahAnak))
			} else {
				fmt.Println(gaji + (tunjangan * 3))
			}
		} else if masaKerja >= 5 && masaKerja <= 10 {
			gaji = 4000
			tunjangan = 100
			if jumlahAnak <= 3 {
				fmt.Println(gaji + (tunjangan * jumlahAnak))
			} else {
				fmt.Println(gaji + (tunjangan * 3))
			}
		}
	} else if jabatan == "manager" || jabatan == "Manager" {
		if masaKerja > 10 {
			gaji = 10000
			tunjangan = 300
			if jumlahAnak <= 3 {
				fmt.Println(gaji + (tunjangan * jumlahAnak))
			} else {
				fmt.Println(gaji + (tunjangan * 3))
			}
		} else if masaKerja <= 10 {
			gaji = 8500
			tunjangan = 300
			if jumlahAnak <= 3 {
				fmt.Println(gaji + (tunjangan * jumlahAnak))
			} else {
				fmt.Println(gaji + (tunjangan * 3))
			}
		}
	} else if jabatan == "direktur" || jabatan == "Direktur" {
		gaji = 20000
		tunjangan = 500
		if jumlahAnak <= 3 {
			fmt.Println(gaji + (tunjangan * jumlahAnak))
		} else {
			fmt.Println(gaji + (tunjangan * 3))
		}
	}
}

func motor_kopling() {
	var gigi int
	var kopling, gas bool
	fmt.Scanln(&gigi, &kopling, &gas)
	if gigi == 0 {
		fmt.Println("Mesin menyala dan motor tidak berjalan")
	} else if gigi > 0 && !kopling && !gas {
		fmt.Println("Mesin mati")
	} else {
		if !kopling && gas {
			fmt.Println("Motor berjalan")
		} else if kopling && !gas {
			fmt.Println("Mesin menyala dan motor tidak berjalan")
		}
	}
}

func bil_belakang() {
	var angka = 9599
	fmt.Println(angka % 10)
}

func polisi() {
	var nim, plat, angkaBelakang int
	fmt.Scanln(&nim, &plat)
	angkaBelakang = plat % 10
	if nim%2 != 0 && angkaBelakang%2 != 0 {
		fmt.Println("NIM ganjil")
		fmt.Println("Plat nomor ganjil")
		fmt.Println("Silahkan lewat")
	} else if nim%2 != 0 && angkaBelakang%2 == 0 {
		fmt.Println("NIM ganjil")
		fmt.Println("Plat nomor genap")
		fmt.Println("Tidak diperbolehkan lewat")
	} else if nim%2 == 0 && angkaBelakang%2 != 0 {
		fmt.Println("NIM genap")
		fmt.Println("Plat nomor ganjil")
		fmt.Println("Tidak diperbolehkan lewat")
	} else if nim%2 == 0 && angkaBelakang%2 == 0 {
		fmt.Println("NIM genap")
		fmt.Println("Plat nomor genap")
		fmt.Println("Silahkan lewat")
	}
}

func bmi() {
	var tinggi, berat, hasil float64
	fmt.Scanln(&tinggi, &berat)
	hasil = berat / (tinggi * tinggi)
	if hasil >= 18.5 && hasil <= 22.9 {
		fmt.Println("normal")
	} else if hasil < 18.5 {
		fmt.Println("kurang")
	} else if hasil >= 23 {
		fmt.Println("lebih")
	}
}

func skye() {
	var hp int
	var isSwarm bool
	fmt.Scanln(&hp)
	fmt.Scanln(&isSwarm)
	if hp >= 0 && hp <= 100 {
		if hp < 50 && hp > 0 {
			if isSwarm {
				fmt.Println("Membuat dinding")
				fmt.Println("Tambah HP")
			} else {
				fmt.Println("Tambah HP")
			}
		} else if hp == 0 {
			fmt.Println("Revive")
			fmt.Println("Membuat dinding")
		} else if isSwarm {
			fmt.Println("Membuat dinding")
		} else {
			fmt.Println("Tidak perlu mengeluarkan skill")
		}
	} else {
		fmt.Println("Range HP melebihi batas")
	}
}

func fotocopy() {
	var isStudent, isSkripsi bool
	var berlangganan, priority int
	var total float64
	fmt.Scanln(&isStudent, &isSkripsi, &berlangganan, &total)
	if isStudent && berlangganan > 12 {
		priority = 1
		if isSkripsi {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 35 / 100))
		} else {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 30 / 100))
		}
	} else if isStudent && berlangganan > 8 && berlangganan <= 12 {
		priority = 2
		if isSkripsi {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 20 / 100))
		} else {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 15 / 100))
		}
	} else if isStudent {
		fmt.Println(total - (total * 10 / 100))
	} else {
		fmt.Println(total)
	}
}

func cekKapital() {
	var x byte
	fmt.Scanf("%c", &x)
	if x > 64 && x <= 90 {
		fmt.Println(string(x), "adalah huruf kapital")
	} else if x > 96 && x <= 122 {
		fmt.Println(string(x), "bukan huruf kapital")
	} else {
		fmt.Println("Inputan bukan huruf!")
	}
}

func menentukanSegitiga() {
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

func jenisSoal() {
	var jenis string
	fmt.Scanln(&jenis)
	if jenis == "mudah" {
		fmt.Println("Keren! Good job! :D")
	} else if jenis == "sedang" {
		fmt.Println("Nice try :D")
	} else if jenis == "sulit" {
		fmt.Println("yu bisa yu")
	} else {
		fmt.Println("Inputan tidak valid")
	}
}

func main() {
	// konsonan()
	// kelipatan()
	// segitiga()
	// mutlak()
	// temperatur()
	// profit()
	// sepak_bola()
	// parkir()
	// akhir_tahun()
	// avatar()
	// urutan_bil()
	// sertifikat()
	// kelipatan()
	// kapal()
	// bola()
	// parcel()
	// gaji()
	// motor_kopling()
	// bil_belakang()
	// polisi()
	// bmi()
	// skye()
	// fotocopy()
	// cekKapital()
	menentukanSegitiga()
	// jenisSoal()
}
