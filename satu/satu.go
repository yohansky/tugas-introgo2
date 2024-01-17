package main

import "fmt"

func main() {
	//1.len() untuk mengetahui panjang suatu tipe data
	var nama string = "Yohanes"
	fmt.Println("panjang nama :", len(nama))

	//2.cap() untuk menghitung kapasitas variabel "slice"
	var slice1 []int = []int{1, 2, 3, 4}
	fmt.Println("kapasitas slice :", cap(slice1))

	//3.append() untuk menambahkan elemen ke variabel "slice"
	var slice2 []int = []int{1, 2, 3, 4}
	fmt.Println("elemen awal slice sebelum append: ", slice2)
	slice2 = append(slice2, 5)
	fmt.Println("elemen awal slice setelah di append: ", slice2)

	//4.copy() untuk mengcopy nilai 1 variabel ke variabel lainnya
	var angkas1 []int = []int{1, 2, 3}
    var angkas2 []int = []int{}
    copy(angkas2, angkas1)
    fmt.Println(angkas2) // [1 2 3]

	//5.make() untuk membuat var slice,map baru
	slice := make([]int, 10)
    fmt.Println(slice) // [0 0 0 0 0 0 0 0 0 0]

	//6.print(), printf(), println() untuk mencetak nilai
	fmt.Print("Halo\n")
	fmt.Printf("Nama saya\n")
	fmt.Println("Yohanes Hubert")

	//7.new() mengalokasikan memori (sama seperti & pada pointer)
	var bilangan int
    fmt.Println(bilangan) // 0
    // Alokasikan memori untuk variabel integer baru
    bilangan2 := new(int)
    fmt.Println(*bilangan2) // 0
    // Atur nilai variabel bilangan2
    *bilangan2 = 10
    fmt.Println(*bilangan2) // 10

	//8.delete() untuk menghapus kunci dan elemen pada map
	m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    delete(m, "a")
    fmt.Println(m) // map[b:2 c:3]

	//9.clear() untuk membersihkan tipe data map,slice 
	tes := []string{"one", "two", "three"}
	clear(tes)
	fmt.Printf("%#v\n", tes)// []string{"", "", ""}

	//10.panic() untuk menghentikan program secara tiba-tiba
	panic("Terjadi kesalahan")


}