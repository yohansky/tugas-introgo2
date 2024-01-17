package main

import (
	"fmt"
	"strings"
)

// catatan pointer (ambil dulu alamatnya menggunakan & lalu ditunjuk oleh pointer memakai *)


func searchName(names *[]string, prefix string, limit int) []string {
    // Definisikan variabel untuk menyimpan hasil pencarian
    foundNames := make([]string, 0, limit) // inisiasi/buat var foundNames dengan aprameter(array string, berindex 0, sampai "limit")

    // Cari nama yang mengandung huruf "an"
    for _, name := range *names { // lakukan perulangan pada pointer array
        if strings.Contains(name, prefix) { // jika nama mengandung(prefix) "an" 
            foundNames = append(foundNames, name) // append(tambahkan) kedalam 
        }

        if len(foundNames) == limit {// jika panjnag var foundNames sampai dengan limit yang ditentukan
            break // maka berhenti
        }
    }

    return foundNames
}

func main() {

	names := []string{
		"Abigail", "Alexandra", "Alison",
		"Amanda", "Angela", "Bella",
		"Carol", "Caroline", "Carolyn",
		"Deirdre", "Diana", "Elizabeth",
		"Ella", "Faith", "Olivia", "Penelope",
	}
	// ambil alamat memori array names untuk dijadikan pointer
	namesPtr := &names
	
	foundNames := searchName(namesPtr, "la", 3)

    // Cetak hasil pencarian
    fmt.Println("Nama yang mengandung huruf 'an':")
    //karena var foundNames adalah array string, maka lakukan perulangan
    for _, name := range foundNames {// iterasi/perulangan pada array foundNames
        fmt.Println(name)
    }
}