package main

import "fmt"

func main() {
	// Mendeklarasikan array dengan tipe data integer
	var angka [5]int
	// Menginisialisasi nilai elemen-elemen array
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	angka[3] = 40
	angka[4] = 50
	// Mengakses dan mencetak nilai elemen-elemen array
	fmt.Println(angka[0]) // Output: 10
	fmt.Println(angka[1]) // Output: 20
	fmt.Println(angka[2]) // Output: 30
	fmt.Println(angka[3]) // Output: 40
	fmt.Println(angka[4]) // Output: 50

	// Mendeklarasikan dan menginisialisasi array secara langsung
	buah := [3]string{"apel", "mangga", "pisang"}

	// Mengakses dan mencetak nilai elemen-elemen array
	fmt.Println(buah[0]) // Output: apel
	fmt.Println(buah[1]) // Output: mangga
	fmt.Println(buah[2]) // Output: pisang
}
