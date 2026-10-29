package main

import (
	"errors" // Untuk membuat pesan error
	"fmt"    // Untuk menampilkan output ke terminal
)

// Struct untuk menyimpan data seorang mahasiswa
type Mahasiswa struct {
	Nim      string
	Nama     string
	Angkatan int
	Ipk      float64
}

// Fungsi untuk mencari mahasiswa berdasarkan NIM
func cariMahasiswa(daftar []Mahasiswa, nim string) (Mahasiswa, error) {
	// Periksa mahasiswa satu per satu
	for i := 0; i < len(daftar); i++ {
		if daftar[i].Nim == nim {
			return daftar[i], nil // Ditemukan, tidak ada error
		}
	}

	// Jika NIM tidak ditemukan
	return Mahasiswa{}, errors.New("mahasiswa tidak ditemukan")
}

// Fungsi untuk menghitung rata-rata IPK
func rerataIpk(daftar []Mahasiswa) float64 {
	if len(daftar) == 0 {
		return 0 // Mencegah pembagian dengan nol
	}

	var total float64

	// Menjumlahkan seluruh IPK mahasiswa
	for i := 0; i < len(daftar); i++ {
		total = total + daftar[i].Ipk
	}

	// Total IPK dibagi jumlah mahasiswa
	return total / float64(len(daftar))
}

func main() {
	// Slice berisi tiga data mahasiswa
	daftar := []Mahasiswa{
		{Nim: "H1A123001", Nama: "Andi Prasetyo", Angkatan: 2023, Ipk: 3.45},
		{Nim: "H1A123002", Nama: "Bunga Lestari", Angkatan: 2023, Ipk: 3.80},
		{Nim: "H1A124003", Nama: "Citra Ramadhani", Angkatan: 2024, Ipk: 3.20},
	}

	// Mencari mahasiswa dengan NIM tertentu
	ditemukan, err := cariMahasiswa(daftar, "H1A123002")
	if err != nil {
		fmt.Println("Galat:", err)
		return
	}

	// Menampilkan nama mahasiswa dan rata-rata IPK
	fmt.Println("Ditemukan:", ditemukan.Nama)
	fmt.Printf("Rerata IPK: %.2f\n", rerataIpk(daftar))

	// Map untuk menghitung jumlah mahasiswa per angkatan
	jumlahPerAngkatan := make(map[int]int)

	for i := 0; i < len(daftar); i++ {
		jumlahPerAngkatan[daftar[i].Angkatan] =
			jumlahPerAngkatan[daftar[i].Angkatan] + 1
	}

	// Menampilkan jumlah mahasiswa tiap angkatan
	fmt.Println("Jumlah per angkatan:", jumlahPerAngkatan)
}
