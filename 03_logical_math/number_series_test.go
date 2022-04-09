package logical_math

import (
	"fmt"
	"testing"
)

func TestNumberSeries(t *testing.T) {
	// Buatkan flowchart dan pseudocodenya untuk menampilkan deret bilangan dengan inputan deret
	// pertama, deret kedua dan nilai bilangan ke x (15 point)
	// Contoh kasus pertama:
	// Inputan deret pertama = 2
	// Inputan deret kedua = 4
	// Value x = 5
	// Maka akan menghasilkan deret angka 2,4,6,8,10
	// Contoh kasus kedua:
	// Inputan deret pertama = 5
	// Inputan deret kedua = 8
	// Value x = 7
	// Maka akan menghasilkan deret angka 5,8,11,14,17,20,23

	// deklarasi test case
	numFirstCaseOne := 2
	numSecondCaseOne := 4
	maxCaseOne := 5
	// cetak hasil dari test case pertama
	fmt.Println("Number Series Case One: ", NumberSeries(numFirstCaseOne, numSecondCaseOne, maxCaseOne))

	// deklarasi test case kedua
	numFirstCaseTwo := 5
	numSecondCaseTwo := 8
	maxCaseTwo := 7
	// cetak hasil dari test case kedua
	fmt.Println("Number Series Case Two: ", NumberSeries(numFirstCaseTwo, numSecondCaseTwo, maxCaseTwo))
}

func NumberSeries(numFirst, numSecond, max int) []int {
	// deklarasi variabel tujuan dengan tipe data array integer
	dest := []int{}

	// masukkan nilai bilangan pertama dan kedua pada variabel tujuan
	dest = append(dest, numFirst, numSecond)

	// cari pembeda antara bilangan pertama dan kedua
	diff := numSecond - numFirst

	// lakukan perulangan
	for {
		// cari posisi terakhir di array tujuan
		lastPosition := len(dest) - 1

		// masukkan nilai baru (bilangan terakhir pada array tujuan dan di tambahkan dengan pembeda) ke dalam array tujuan
		dest = append(dest, dest[lastPosition]+diff)

		// jika panjang dari variabel tujuan sudah sama dengan maksimal
		if len(dest) == max {
			// maka berhentikan seluruh perulangan
			break
		}
	}

	// kembalikan nilai dari variabel tujuan
	return dest
}
