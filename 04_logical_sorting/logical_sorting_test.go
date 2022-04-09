package logical_sorting

import (
	"fmt"
	"testing"
)

func TestLogicalSortings(t *testing.T) {
	// Terdapat variabel array dengan nilai [4, -7, -5, 3, 3.3, 9, 0, 10, 0.2]. Buatkan flowchart dan
	// pseudocodeu ntuk melakukan logical sorting variabel tersebut baik secara ascending maupun
	// descending (10 point)
	// Hasil yang diharapkan ketika sorting secara ascending
	// [-7, -5, 0, 0.2, 3, 3.3, 4, 9, 10]
	// Hasil yang diharapkan ketika sorting secara descending
	// [10, 9, 4, 3.3, 3, 0.2, 0, -5, -7]

	// deklarasi test case
	dataInput := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}

	// cetak test case yang masih acak
	fmt.Println("Test Case: ", dataInput)

	// buat dependency injection baru untuk urutan ascending
	sortAsc := NewQuickSort(dataInput)
	// cetak hasil dari urutan ascending
	fmt.Println("Quick Sort Ascending: ", sortAsc.Ascending())

	// buat dependency injection baru untuk urutan descending
	sortDesc := NewQuickSort(dataInput)
	// cetak hasil dari urutan descending
	fmt.Println("Quick Sort Descending: ", sortDesc.Descending())
}

// inisialisasi interface, adalah function yang bisa di panggil di luar dari package logical_sorting setelah menginisialisasi fungsi NewQuickSort
type QuickSort interface {
	Ascending() []float64
	Descending() []float64
}

// inisialisasi struct
type quickSort struct {
	// isi array yang akan di sort, sesuaikan dengan tipe data dari data yang akan di urutkan
	nums []float64

	// default boolean = false
	order bool
}

// buat dependency injection
func NewQuickSort(nums []float64) QuickSort {
	return quickSort{
		nums: nums,
	}
}

func (q *quickSort) swap(i, j int) {
	// menukar posisi i menjadi j dan j menjadi 1
	q.nums[i], q.nums[j] = q.nums[j], q.nums[i]
}

func (q *quickSort) partition(low, high int) int {
	// mengambil posisi yang paling tinggi, dari posisi poros (posisi terakhir)
	pivot := q.nums[high]

	// cari posisi dari yang paling bawah berdasarkan poros (posisi terakhir)
	i := (low - 1)

	// melakukan perulangan
	for j := low; j <= high-1; j++ {
		// lambda function
		swap := func() {
			// increment posisi i ditambah dengan 1
			i++

			// tukar posisi i dengan j
			q.swap(i, j)
		}

		// jika ordernya adalah true, maka default ordernya adalah ascending
		// dan posisi array float64 yang berada di posisi j lebih besar dari posisi poros (posisi terakhir), maka panggil lambda function swap
		if q.order && q.nums[j] > pivot {
			swap()
		}

		// default nilai dari boolean adalah false, maka default ordernya adalah ascending
		// dan posisi array float64 yang berada di posisi j lebih kecil dari posisi poros (posisi terakhir), maka panggil lambda function swap
		if !q.order && q.nums[j] < pivot {
			swap()
		}
	}

	// tukar posisi i ditambah dengan 1 (posisi setelahnya) dengan posisi poros (posisi terakhir)
	q.swap(i+1, high)

	// kembalikan nilai dari posisi setelah i (i diambah dengan 1)
	return (i + 1)
}

// function recursive (function yang akan memanggil dirinya sendiri)
func (q *quickSort) quickSort(low, high int) {
	// jika posisi pertama lebih kecil dari poros (posisi terakhir)
	if low < high {

		// cari posisi partisi dari array float64
		pi := q.partition(low, high)

		// mengurutkan array float sebelum partisi
		q.quickSort(low, pi-1)
		// mengurutkan array float setelah partisi
		q.quickSort(pi+1, high)
	}
}

func (q quickSort) Ascending() []float64 {
	// panggil fungsi quickSort dengan posisi pertama dan posisi poros (posisi terakhir)
	q.quickSort(0, len(q.nums)-1)

	// kembalikan array float64
	return q.nums
}

func (q quickSort) Descending() []float64 {
	// ubah order menjadi true
	q.order = true

	// panggil fungsi quickSort dengan posisi pertama dan posisi poros (posisi terakhir)
	q.quickSort(0, len(q.nums)-1)

	// kembalikan array float64
	return q.nums
}
