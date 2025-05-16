package main

import (
	"fmt"
	"strings"
)

type Ide struct {
	Judul    string
	Kategori string
	Voting   int
	Tanggal  string
}

// const data = 100
// var ideas [data]Ide // ganti jadi array tetap
// var totalIde int
var ideas []Ide

func Tambahide(judul, kategori, tanggal string) {
	Idebaru := Ide{Judul: judul, Kategori: kategori, Voting: 0, Tanggal: tanggal}
	ideas = append(ideas, Idebaru)
	fmt.Println("Ide berhasil ditambahkan:", judul)
}

//func Tambahide(judul, kategori, tanggal string) {
//	if totalIde >= data {
//		fmt.Println("Gagal menambahkan ide: kapasitas penuh.")
//		return
//	}
//	ideas[totalIde] = Ide{
//		Judul:     judul,
//		Kategori: kategori,
//		Voting:    0,
//		Tanggal:   tanggal,
//	}
//	totalIde++
//	fmt.Println("Ide berhasil ditambahkan:", judul)
//}

func Editide(judulLama, judulBaru, kategoriBaru, tanggalBaru string) {
	for i := range ideas {
		if strings.EqualFold(ideas[i].Judul, judulLama) {
			ideas[i].Judul = judulBaru
			ideas[i].Kategori = kategoriBaru
			ideas[i].Tanggal = tanggalBaru

			fmt.Println("Ide berhasil diubah:", judulBaru)
			return
		}
	}
	fmt.Println("Ide tidak ditemukan.")
}

//func Hapuside

//func Votingide

//func Searchide

//func SortideVoting

//func SortideTanggal

func main() {
	var fitur int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambahkan Ide")
		fmt.Println("2. Edit Ide")
		fmt.Println("3. Hapus Ide")
		fmt.Println("4. Voting Ide")
		fmt.Println("5. Cari Ide")
		fmt.Println("6. Urutkan Ide Berdasarkan Voting")
		fmt.Println("7. Urutan Ide Berdasarkan Tanggal")
		fmt.Println("8. Tampilkan Semua Ide")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&fitur)

		switch fitur {
		case 1:
			var judul, kategori, tanggal string

			fmt.Print("Masukkan judul ide: ")
			fmt.Scanln(&judul)
			fmt.Print("Masukkan deskripsi ide: ")
			fmt.Scanln(&kategori)
			fmt.Print("Masukkan tanggal dibuat (YYYY-MM-DD): ")
			fmt.Scanln(&tanggal)
			Tambahide(judul, kategori, tanggal)
		case 2:
			var judulLama, judulBaru, kategoriBaru, tanggalBaru string

			fmt.Print("Masukkan judul ide yang ingin diedit: ")
			fmt.Scan(&judulLama)
			fmt.Print("Masukkan judul baru: ")
			fmt.Scan(&judulBaru)
			fmt.Print("Masukkan deskripsi baru: ")
			fmt.Scan(&kategoriBaru)
			fmt.Print("Masukkan tanggal baru: ")
			fmt.Scan(&tanggalBaru)

			Editide(judulLama, judulBaru, kategoriBaru, tanggalBaru)

		case 3:
			var judul string
			fmt.Print("Masukkan judul ide yang ingin dihapus: ")
			fmt.Scan(&judul)
			//		if !Hapuside(judul) {
			//			fmt.Println("Ide tidak ditemukan.")
			//		}

		case 4:
			var judul string
			var rating int
			fmt.Print("Masukkan judul ide: ")
			fmt.Scan(&judul)
			fmt.Print("Masukkan rating baru: ")
			fmt.Scan(&rating)
			//		if !Ratingide(judul, rating) {
			//			fmt.Println("Ide tidak ditemukan.")
			//		}

		case 5:

		case 6:

		case 7:
			var urutan int

			for {
				fmt.Println("\nmenu: ")
				fmt.Println("1. ascending")
				fmt.Println("2. descanding")
				fmt.Println("Pilih urutan: ")
				fmt.Scanln(&urutan)

				switch urutan {
				case 1:

				case 2:

				}
			}
		case 8:
			if len(ideas) == 0 {
				fmt.Println("Belum ada ide yang ditambahkan.")
			} else {
				fmt.Println("Daftar semua ide:")

				for i := 0; i < len(ideas); i++ {
					ide := ideas[i]
					fmt.Printf("Judul: %s\nDeskripsi: %s\nTanggal: %s\nRating: %d\n\n",
						ide.Judul, ide.Kategori, ide.Tanggal, ide.Voting)
				}
			}
			//			kalau  pakai const tetap
			//			fmt.Println("Daftar semua ide:")
			//			for i := 0; i < totalIde; i++ {
			//			fmt.Printf("Judul: %s, Deskripsi: %s, Tanggal: %s, Rating: %d\n",
			//			ideas[i].Judul, ideas[i].Deskripsi, ideas[i].Tanggal, ideas[i].Rating)
			//	}

		case 9:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
