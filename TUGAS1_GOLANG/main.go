package main

import (
	"fmt"
)

func main() {
	// Masukkan invoice /  tagihan
	fmt.Print("Masukkan Invoice: ")
	var tagihan int
	_, err := fmt.Scanf("%d", &tagihan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pendaklarasian variabel tip
	var tip float64

	// Menghitung tiap inputan berdasarkan aturan
	if tagihan >= 50 && tagihan <= 300 {
		tip = float64(tagihan) * 0.15
	} else {
		tip = float64(tagihan) * 0.20
	}

	// Hitung total nilai (tagihan + tip)
	total := float64(tagihan) + tip

	// Menampilkan hasil di konsol
	fmt.Printf("Invoice : %d\n tipnya : %.2f\n dan nilai akhir : %.2f\n", tagihan, tip, total)
}
