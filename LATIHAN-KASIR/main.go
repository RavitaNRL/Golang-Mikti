package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputJumlah, inputBarang string

	fmt.Println(`
	Daftar Barang :
	1. Mouse @ Rp 100.000
	2. Keyboard @ Rp 200.000
	3. Sound @ Rp 300.000
	`)

	fmt.Println("Silakan masukkan kode barang (1-3)")
	fmt.Scanln(&inputBarang)

	kodeBarang, err := strconv.Atoi(inputBarang)
	if err != nil || kodeBarang < 1 || kodeBarang > 3 {
		fmt.Println("Kode barang tidak valid.")
		return
	}

	hargaSatuan := 0

	switch kodeBarang {
	case 1:
		hargaSatuan = 100000
	case 2:
		hargaSatuan = 200000
	case 3:
		hargaSatuan = 300000
	}

	fmt.Println("Silakan masukkan jumlah barang yang dibeli: ")
	fmt.Scanln(&inputJumlah)

	jumlahBarang, err := strconv.Atoi(inputJumlah)
	if err != nil || jumlahBarang <= 0 {
		fmt.Println("Jumlah barang tidak valid.")
		return
	}

	total := hargaSatuan * jumlahBarang
	fmt.Println("Total belanjaan Anda adalah: Rp", total)

	cetakInvoice(inputBarang, hargaSatuan, jumlahBarang, total)
}

func cetakInvoice(kodeBarang string, hargaSatuan, jumlah, total int) {
	namaBarang := ""
	switch kodeBarang {
	case "1":
		namaBarang = "Mouse"
	case "2":
		namaBarang = "Keyboard"
	case "3":
		namaBarang = "Sound"
	}

	invoice := fmt.Sprintf(
		"INVOICE TOKO MAKMUR ABADI\nNama Barang: %s\nHarga Satuan: Rp%d\nJumlah Barang: %d\nTotal Belanja: Rp%d\n",
		namaBarang, hargaSatuan, jumlah, total,
	)

	file, err := os.Create("invoice.txt")
	if err != nil {
		fmt.Println("Gagal mencetak invoice: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(invoice)
	if err != nil {
		fmt.Println("Gagal menulis ke file: ", err)
		return
	}

	fmt.Println("Invoice berhasil tercetak ke file invoice.txt")
}
