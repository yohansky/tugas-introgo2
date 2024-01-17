package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func calculateArea(Width, Height float64) float64 {
	return Width * Height // untuk perhitungan rumus luas
}

func cetakLuas(luasSegitiga float64) {
	fmt.Println(luasSegitiga)// untuk cetak
}

func main() {

	segitiga := Rectangle{
		Width:  10,
		Height: 12}
	fmt.Println(segitiga)
	luasSegitiga := calculateArea(segitiga.Width, segitiga.Height)
	cetakLuas(luasSegitiga)
}