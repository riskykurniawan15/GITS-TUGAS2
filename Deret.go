package main

import (
	"fmt"
)

func main(){
	// Risky Kurniawan - Universitas Adhirajasa Resawara Sanjaya
	// Buat sebuah program deret ganjil dan genap dengan jumlah deret harus sesuai dengan inputan user.

	var awal, akhir, n_ganjil, n_genap int
	fmt.Println("========================================")
	fmt.Println("          DERET GANJIL GENAP            ")
	fmt.Println("========================================")
	fmt.Print("Masukan nilai awal  : ")
	fmt.Scan(&awal)
	fmt.Print("Masukan nilai akhir : ")
	fmt.Scan(&akhir)

	if (awal%2) == 1 {
		n_ganjil = awal
		n_genap = awal + 1
	}else{
		n_ganjil = awal + 1
		n_genap = awal
	}
	fmt.Println("========================================")
	fmt.Print("BILANGAN GANJIL = ")
	
	for i:= n_ganjil; i <= akhir; i += 2 {
		fmt.Printf("%d ", i)
	}
	
	fmt.Print("\nBILANGAN GENAP  = ")
	for i:= n_genap; i <= akhir; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n========================================")
}