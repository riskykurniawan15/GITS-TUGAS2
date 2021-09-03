package main

import (
	"fmt"
)

func main(){
	// Risky Kurniawan - Universitas Adhirajasa Resawara Sanjaya
	// Buat sebuah program Konversi suhu ruangan dari celcius ke fahrenheit dan kelvin.
	var celcius int
	fmt.Println("===========================================")
	fmt.Print("Masukan besar suhu dalam celcius : ")
	fmt.Scan(&celcius)
	
	fmt.Println("===========================================")
	fmt.Printf("Suhu Celcius = %d 'C \n", celcius)
	fmt.Println("===========================================")
	fahrenheit := ((float64(celcius) * 9)/5) + 32
	fmt.Printf("Dalam Fahrenheit = %.3f 'F\n", fahrenheit)
	kelvin := float64(celcius) + 273.15
	fmt.Printf("Dalam Kelvin     = %.2f 'K\n", kelvin)
	fmt.Println("===========================================")
}