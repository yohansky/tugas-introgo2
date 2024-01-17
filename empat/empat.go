package main

import (
	"fmt"
	"sort"
)

// pakai metode variadict function

func SeleksiNilai(nilaiAwal int, nilaiAkhir int, dataArray *[]int) []int {
	if nilaiAkhir <= nilaiAwal {
		panic("Nilai akhir harus lebih besar dari nilai awal")
	}
	if len(*dataArray) <= 5 {
		panic("Jumlah angka dalam dataArray harus lebih dari 5 elemen")
	}
	
	var hasil []int
	for _, v := range *dataArray {
		if v >= nilaiAwal && v <= nilaiAkhir {
			hasil = append(hasil, v)
		}		
	}
	//no 4 
	sort.Ints(hasil)
	return hasil
}

func main() {
	dataArray := []int{2, 25, 4, 14, 17, 30, 8} // no 1 & 2
	// dataArray := []int{2, 25, 4} // no 3
	// dataArray := []int{2, 25, 4, 1, 30, 18} // no 4

	hasil := SeleksiNilai(5, 20, &dataArray) //no 1
	// hasil := SeleksiNilai(15, 3, &dataArray) //no 2
	// hasil := SeleksiNilai(5, 17, &dataArray) //no 3
	// hasil := SeleksiNilai(5, 17, &dataArray) //no 4
	if len(hasil) == 0 {
		fmt.Println("Nilai TIdak ditemukan")
	}else {
	fmt.Println(hasil)
	}
}