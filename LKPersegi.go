package main

import (
	"fmt"
)

func main(){
	// Risky Kurniawan - Universitas Adhirajasa Resawara Sanjaya
	// Buat program untuk menghitung Luas dan keliling dari persegi panjang.Untuk Panjang dan Lebar nya sendiri itu di dapat dari inputan user

	var panjang, lebar int
	fmt.Println("============================================")
	fmt.Println("   HITUNG LUAS & KELILING PERSEGI PANJANG   ")
	fmt.Println("============================================")
	fmt.Print("Masukan Panjang : ")
	fmt.Scan(&panjang)
	fmt.Print("Masukan Lebar   : ")
	fmt.Scan(&lebar)

	luas := panjang * lebar
	keliling := (panjang * 2) + (lebar * 2)
	fmt.Println("============================================")
	fmt.Printf("Luas persegi panjang     = %d\n", luas)
	fmt.Printf("Keliling persegi panjang = %d\n", keliling)
	fmt.Println("============================================")
}