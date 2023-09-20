package main

import (
	"fmt"
)

func luasTrapesium(a, b, t float64) float64 {
	return (a + b) * t / 2
}

func isPrime(n int) bool {
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isMultipleOfSeven(n int) bool {
	if n%7 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Latihan")

	fmt.Println("1. Luas Trapesium")
	fmt.Println("2. Bilangan Prima")
	fmt.Println("3. Kelipatan 7")
	fmt.Println("Pilih fungsi yang akan dijalankan: ")

	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var a, b, t float64
		fmt.Scan(&a, &b, &t)
		fmt.Println(luasTrapesium(a, b, t))
	case 2:
		var angka int
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&angka)
		fmt.Println(isPrime(angka))
	case 3:
		var angka int
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&angka)
		fmt.Println(isMultipleOfSeven(angka))
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}
