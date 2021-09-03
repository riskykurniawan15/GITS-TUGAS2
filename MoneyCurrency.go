package main

import (
	"fmt"
	"strings"
	"math"
)

func main(){
	// Risky Kurniawan - Universitas Adhirajasa Resawara Sanjaya
	// Buat program input output berupa konversi dari integer(numerik) menjadi format currency (Rupiah & dollar).
	var nilai int = 25000
	fmt.Println("========================================")
	fmt.Println("     PROGRAM KONVERSI ANGKA KE UANG")
	fmt.Println("========================================")
	fmt.Print("Masukan bilangan bulat : ")
	fmt.Scan(&nilai)
	fmt.Println("========================================")
	fmt.Printf("Format Rupiah    = %s\n", currency(nilai, "IDR"))
	fmt.Printf("Format US Dollar = %s\n", currency(nilai, "USD"))
	fmt.Println("========================================")
	fmt.Println("Kurs            = $ 1.00 USD -> Rp. 14.251,55")
	fmt.Println("========================================")
}

func currency(number int, format string) string{
	var money, sim string
	var nilai float64
	if format == "IDR"{
		money = "Rp. "
		sim = "."
		nilai = (math.Round((float64(number)) * 100)/100)
	}else if format == "USD"{
		money = "$ "
		sim = ","
		nilai = (math.Round((float64(number)/(14251.55)) * 100)/100)
	}else{
		return "Error Number Format"
	}

	valnum := strings.Split(fmt.Sprint(nilai), ".")[0]
	var reverse string = ""

	var index int = 1
	
	for i:=(strings.Count(valnum, "")-2); i>=0; i--{
		if (index%3) == 1 && index != 1{
			reverse += sim
		}
		reverse += string(valnum[i])
		
		index++
	}

	valnum = reverse

	for i:=(strings.Count(valnum, "")-2); i>=0; i--{
		money += string(valnum[i])
	}

	if format == "USD"{
		money += "." + strings.Split(fmt.Sprint(nilai), ".")[1] + " USD"
	}

	if format == "IDR"{
		money += ",00"
	}
	
	return money
}