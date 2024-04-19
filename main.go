package main

import "fmt"

// Produk adalah antarmuka untuk semua produk yang dijual di supermarket.
type Produk interface {
	TampilkanInfo() string
}

// Keripik adalah representasi keripik yang dijual di supermarket.
type Keripik struct {
	Nama string
}

// TampilkanInfo mengimplementasikan metode TampilkanInfo untuk menampilkan informasi keripik.
func (k *Keripik) TampilkanInfo() string {
	return fmt.Sprintf("Menghadirkan %s di bagian makanan ringan", k.Nama)
}

// Kerupuk adalah representasi kerupuk yang dijual di supermarket.
type Kerupuk struct {
	Nama string
}

// TampilkanInfo mengimplementasikan metode TampilkanInfo untuk menampilkan informasi kerupuk.
func (k *Kerupuk) TampilkanInfo() string {
	return fmt.Sprintf("Menyediakan %s di bagian makanan ringan", k.Nama)
}

// Pabrik adalah antarmuka untuk semua pabrik produk di supermarket.
type Pabrik interface {
	BuatProduk() Produk
}

// PabrikKeripik adalah pabrik konkret yang membuat keripik.
type PabrikKeripik struct{}

// BuatProduk membuat instance baru dari keripik.
func (f *PabrikKeripik) BuatProduk() Produk {
	return &Keripik{Nama: "Keripik Kentang"}
}

// PabrikKerupuk adalah pabrik konkret yang membuat kerupuk.
type PabrikKerupuk struct{}

// BuatProduk membuat instance baru dari kerupuk.
func (f *PabrikKerupuk) BuatProduk() Produk {
	return &Kerupuk{Nama: "Kerupuk Udang"}
}

func main() {
	// Membuat sebuah instance dari PabrikKeripik
	pabrikKeripik := &PabrikKeripik{}
	// Menggunakan pabrik untuk membuat sebuah produk keripik
	produkKeripik := pabrikKeripik.BuatProduk()
	// Menampilkan informasi produk keripik
	fmt.Println(produkKeripik.TampilkanInfo())

	// Membuat sebuah instance dari PabrikKerupuk
	pabrikKerupuk := &PabrikKerupuk{}
	// Menggunakan pabrik untuk membuat sebuah produk kerupuk
	produkKerupuk := pabrikKerupuk.BuatProduk()
	// Menampilkan informasi produk kerupuk
	fmt.Println(produkKerupuk.TampilkanInfo())
}
