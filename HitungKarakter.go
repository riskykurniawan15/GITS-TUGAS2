package main

import (
	"fmt"
	"strings"
)

func main(){
	// Risky Kurniawan - Universitas Adhirajasa Resawara Sanjaya
	// Buat sebuah program untuk mengetahui jumlah karakter yang di input dan karakter yang sama.

	var karakter string
	fmt.Println("========================================")
	fmt.Print("Masukan Karakter : ")
	fmt.Scan(&karakter)
	fmt.Println("========================================")
	fmt.Printf("Jumlah Karakter '%s' adalah '%d' karakter\n", karakter, (strings.Count(karakter, "")-1))
	
	fmt.Println("========================================")
	fmt.Println("Jumlah Karakter Setiap Huruf (CamelCase)")
	fmt.Println("========================================")

	for {
		fmt.Printf("Jumlah Huruf '%s' berjumlah '%d'\n", string(karakter[0]), (strings.Count(karakter, string(karakter[0]))))
		karakter = strings.Replace(karakter, string(karakter[0]), "", -1)
		if (strings.Count(karakter, "")-1) <= 0{
			break;
		}
	}

	fmt.Println("========================================")
}