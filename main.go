package main

import "fmt"

type Produk interface {
	TampilkanInfo() string
}

type Keripik struct {
	Nama string
}

func (k *Keripik) TampilkanInfo() string {
	return fmt.Sprintf("Menghadirkan %s di bagian makanan ringan", k.Nama)
}

type Kerupuk struct {
	Nama string
}

func (k *Kerupuk) TampilkanInfo() string {
	return fmt.Sprintf("Menyediakan %s di bagian makanan ringan", k.Nama)
}

type Pabrik interface {
	BuatProduk() Produk
}

type PabrikKeripik struct{}

func (f *PabrikKeripik) BuatProduk() Produk {
	return &Keripik{Nama: "Keripik Kentang"}
}

type PabrikKerupuk struct{}

func (f *PabrikKerupuk) BuatProduk() Produk {
	return &Kerupuk{Nama: "Kerupuk Udang"}
}

func main() {
	pabrikKeripik := &PabrikKeripik{}
	produkKeripik := pabrikKeripik.BuatProduk()
	fmt.Println(produkKeripik.TampilkanInfo())

	pabrikKerupuk := &PabrikKerupuk{}
	produkKerupuk := pabrikKerupuk.BuatProduk()
	fmt.Println(produkKerupuk.TampilkanInfo())
}
