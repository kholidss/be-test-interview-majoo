package code_block

import (
	"errors"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

//! abaikan kode di bawah ini
// dalam code block ini, menggunakan packae gorm & zerolog
// contoh pemanggilan init db
func InitDB() *gorm.DB {
	return &gorm.DB{}
}

// contoh dependency injection repository
type AreaRepository interface {
	InsertArea(float64, float64, string) error
}

type areaRepository struct {
	DB *gorm.DB
}

func NewAreaRepository(db *gorm.DB) AreaRepository {
	return &areaRepository{
		DB: db,
	}
}

// contoh dependency injection service
type AreaService struct {
	repository AreaRepository
}

func NewAreaService() AreaService {
	return AreaService{
		repository: NewAreaRepository(InitDB()),
	}
}

// contoh translation
type Translation struct {
	ERROR_DATABASE string
}

func NewTranslation(lang string) Translation {
	switch lang {
	case "en":
		return Translation{
			ERROR_DATABASE: "Error Database",
		}
	default:
		return Translation{
			ERROR_DATABASE: "Ada masalah di Database",
		}
	}
}

//! Start Code Here
// Struct Area
type (
	Area struct {
		ID int64 `gorm:"column:id;primaryKey;"`
		//! AreaValue int64  `gorm:"column:area_value"`
		// saya ganti tipe data dari AreaValue dari int64 menjadi float64
		// karena float64 itu bisa menampung data desimal sedangkan int64 itu tidak bisa
		AreaValue float64 `gorm:"column:area_value"`
		AreaType  string  `gorm:"column:type"`
	}
)

// Example Service in clean architecture
func (_u AreaService) AreaService() error {
	// contoh deklarasi error with translation
	en := NewTranslation("en")

	//! err = _u.repository.InsertArea(10, 10, ‘persegi’)
	// ‘persegi’ = salah, string harus menggunakan (") "petik dua"
	// "err = " (sama dengan saja) salah, jika variabel err adalah deklarasi pertama, maka cara pendeklarasiannya ada 2 macam:
	// a. buat variable err dengan tipe data error (var err error) sebelum (di atas) menulis kode tersebut (err = ...)
	// b. menambah titik dua (:) sebelum sama dengan (:=) untuk mendeklarasikan secara langsung
	err := _u.repository.InsertArea(10, 10, "persegi")
	if err != nil {
		// with package zerolog
		log.Error().Msg(err.Error())
		// ganti error message dengan translation
		err = errors.New(en.ERROR_DATABASE)
		return err
	}

	// kembalikan nilai error = nil ketika AreaService berhasil sampai akhir (selesai)
	return nil
}

//! func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area) (err error) {
// karena saya memberi contoh AreaRepository interface, maka saya harus areaRepository karena dia struct yang merujuk pada struct yang sudah di buat pada dependency injection
// karena param yang dimasukkan pada area itu adalah string, maka attribut ke-3 dan ke-4 itu itu tidak perlu ("arrayType []string, ar *Model.Area") dan ganti sesuai dengan ketentuannya yakni dengan tipe data string
// penamaan variabel type pada attribute ke-3 itu redundant, karena type yang di maksud bisa diarikan sebagai type (type dari golang) bukan nama sebuah variabel, jadi harus di ganti seuai dengen nama variabel yang akan digunakan, sebagai contoh areaType
// untuk tipe data attribute ke-1 dan ke-2, menurut saya float64 karena attribute tersebut bisa memasukkan bilangan desimal
func (_r *areaRepository) InsertArea(param1 float64, param2 float64, areaType string) (err error) {
	//!
	// Perlu pendeklarasian variabel ar itu adalah tipe data/struct yang mana
	// karena ar adalah data yang akan dimasukkan ke dalam gorm, maka ar adalah struct dari Area (sesuai dengan ketentuan)
	ar := Area{}

	//! inst := _r.DB.Model(ar)
	// saya hapus karena tidak perlu

	//! Var area int
	// untuk penulisan pendeklarasian variable (var) itu lowercase, huruf V nya itu kecil
	// saya mengganti tipe data menjadi float64 supaya bisa menampung nilai desimal
	var area float64

	//! switch type {
	// disesuaikan dengan nama yang ada di attribute ke-3, karena sudah di ganti menjadi areaType maka harus di ganti
	switch areaType {

	//! case ‘persegi panjang’:
	// ‘persegi panjang’ = salah, string harus menggunakan (") "petik dua"
	case "persegi panjang":
		//! var area := param1 * param2
		// karena variable area sudah di deklarasi, maka penulisan area tersebut jangan menggunakan titik dua (:), karena dengan adanya titik dua maka area akan di deklarasikan kembali dan  menimbulkan error
		// untuk var saya hapus karena tanpa var juga sudah menjelaskan bahwa area itu adalah variabel
		area = param1 * param2

		ar.AreaValue = area

		//! ar.AreaType = ‘persegi panjang’
		// ‘persegi panjang’ = salah, string harus menggunakan (") "petik dua"
		ar.AreaType = "persegi panjang"

		//! err = _r.DB.create(&ar).Error
		// pemanggilan gorm Create itu, penulisan function create huruf C nya itu harus kapital, karena kalau c nya kecil hanya bisa di akses dalam satu package dan dalam folder yang sama
		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}

	//! case ‘persegi’:
	// ‘persegi’ = salah, string harus menggunakan (") "petik dua"
	case "persegi":
		//! var area = param1 * param2
		// untuk var saya hapus karena tanpa var juga sudah menjelaskan bahwa area itu adalah variabel
		area = param1 * param2

		ar.AreaValue = area

		//! ar.AreaType = ‘persegi’
		// ‘persegi’ = salah, string harus menggunakan (") "petik dua"
		ar.AreaType = "persegi"

		//! err = _r.DB.create(&ar).Error
		// pemanggilan gorm Create itu, penulisan function create huruf C nya itu harus kapital, karena kalau c nya kecil hanya bisa di akses dalam satu package dan dalam folder yang sama
		// untuk tidak memanggilnya secara berulang ulang, lebih baik di panggil sekali setelah switch case ini
		// err = _r.DB.Create(&ar).Error
		// if err != nil {
		// 	return err
		// }

	//! case segitiga:
	// ‘segitiga’ = salah, string harus menggunakan (") "petik dua"
	case "segitiga":
		area = 0.5 * (param1 * param2)
		ar.AreaValue = area

		//! ar.AreaType = ‘segitiga’
		// ‘segitiga’ = salah, string harus menggunakan (") "petik dua"
		ar.AreaType = "segitiga"

		//! err = _r.DB.create(&ar).Error
		// pemanggilan gorm Create itu, penulisan function create huruf C nya itu harus kapital, karena kalau c nya kecil hanya bisa di akses dalam satu package dan dalam folder yang sama
		// untuk tidak memanggilnya secara berulang ulang, lebih baik di panggil sekali setelah switch case ini
		// err = _r.DB.Create(&ar).Error
		// if err != nil {
		// 	return err
		// }

	default:
		ar.AreaValue = 0

		//! ar.AreaType = ‘undefined data’
		// ‘undefined data’ = salah, string harus menggunakan (") "petik dua"
		ar.AreaType = "undefined data"

		//! err = _r.DB.create(&ar).Error
		// pemanggilan gorm Create itu, penulisan function create huruf C nya itu harus kapital, karena kalau c nya kecil hanya bisa di akses dalam satu package dan dalam folder yang sama
		// untuk tidak memanggilnya secara berulang ulang, lebih baik di panggil sekali setelah switch case ini
		// err = _r.DB.Create(&ar).Error
		// if err != nil {
		// 	return err
		// }
	}

	// dipanggil sekali setelah switch case
	err = _r.DB.Create(&ar).Error
	if err != nil {
		return err
	}

	// kembalikan nilai error = nil ketika InsertArea berhasil sampai akhir (selesai)
	return nil
}
