package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Ide struct {
	ID       int
	Judul    string
	Kategori string
	Voting   int
	Tanggal  string
}

var ideas []Ide
var noID = 1

func Tambahide() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan judul ide: ")
	judul, _ := reader.ReadString('\n') // misal user input: " Ide Baru\n"
	judul = strings.TrimSpace(judul)    // hasilnya jadi "Ide Baru" tanpa spasi atau newline

	//Kenapa penting? Karena ReadString('\n') itu akan menyertakan karakter newline \n di akhir input yang kita ketik. Kalau kita tidak bersihkan dengan TrimSpace, stringnya akan ada karakter tak terlihat yang bisa bikin logic kita error atau string jadi tidak sesuai harapan.

	fmt.Print("Masukkan kategori ide: ")
	kategori, _ := reader.ReadString('\n')
	kategori = strings.TrimSpace(kategori)

	fmt.Print("Masukkan tanggal dibuat (YYYY-MM-DD): ")
	tanggal, _ := reader.ReadString('\n')
	tanggal = strings.TrimSpace(tanggal)

	Idebaru := Ide{
		ID:       noID,
		Judul:    judul,
		Kategori: kategori,
		Voting:   0,
		Tanggal:  tanggal}
	ideas = append(ideas, Idebaru)
	noID++
	fmt.Println("✅Ide berhasil ditambahkan:", judul, "(ID: ", Idebaru.ID, ") 🎉")
}

func EditideByID() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID ide yang ingin diedit: ")
	var id int
	fmt.Scanln(&id)

	// Cari ide berdasarkan ID
	for i := range ideas {
		if ideas[i].ID == id {
			// Jika ditemukan, input data baru
			fmt.Print("Masukkan judul baru: ")
			judul, _ := reader.ReadString('\n')
			judul = strings.TrimSpace(judul)

			fmt.Print("Masukkan kategori baru: ")
			kategori, _ := reader.ReadString('\n')
			kategori = strings.TrimSpace(kategori)

			fmt.Print("Masukkan tanggal baru (YYYY-MM-DD): ")
			tanggal, _ := reader.ReadString('\n')
			tanggal = strings.TrimSpace(tanggal)

			// Update data ide
			ideas[i].Judul = judul
			ideas[i].Kategori = kategori
			ideas[i].Tanggal = tanggal

			fmt.Println("✅ Ide dengan ID", id, "berhasil diubah. ✏️")
			return
		}
	}

	fmt.Println("⚠️ Ide dengan ID", id, "tidak ditemukan. ❌")
}

func HapusideByID() bool {
	var id int
	fmt.Print("Masukkan ID ide yang ingin dihapus: ")
	fmt.Scanln(&id)

	for i := 0; i < len(ideas); i++ {
		if ideas[i].ID == id {
			ideas = append(ideas[:i], ideas[i+1:]...)
			fmt.Printf("✅ Ide dengan ID %d berhasil dihapus. 🗑️\n", id)
			return true
		}
	}
	fmt.Printf("❌ Ide dengan ID %d tidak ditemukan.🚫\n", id)
	return false
}

func VotingByID() bool {
	var id, rating int
	fmt.Print("Masukkan ID ide yang ingin diberi voting: ")
	fmt.Scanln(&id)

	fmt.Print("Masukkan nilai voting baru (misal 1-10): ")
	fmt.Scanln(&rating)

	for i := 0; i < len(ideas); i++ {
		if ideas[i].ID == id {
			ideas[i].Voting = rating
			fmt.Printf("✅ Voting ide dengan ID %d berhasil diubah menjadi %d. 👍\n", id, rating)
			return true
		}
	}

	fmt.Printf("❌ Ide dengan ID %d tidak ditemukan. 🚫\n", id)
	return false
}

func SearchSequentialByID(id int) *Ide {
	for i := 0; i < len(ideas); i++ {
		if ideas[i].ID == id {
			return &ideas[i]
		}
	}
	return nil
}

func SearchSequentialByJudul(judul string) *Ide {
	judul = strings.ToLower(judul)
	for i := 0; i < len(ideas); i++ {
		if strings.ToLower(ideas[i].Judul) == judul {
			return &ideas[i]
		}
	}
	return nil
}

func SearchBinaryByID(id int) *Ide {
	kr := 0
	kn := len(ideas) - 1

	for kr <= kn {
		mid := (kr + kn) / 2
		if ideas[mid].ID == id {
			return &ideas[mid]
		} else if ideas[mid].ID < id {
			kr = mid + 1
		} else {
			kn = mid - 1
		}
	}
	return nil
}

func SearchBinaryByJudul(judul string) *Ide {
	judul = strings.ToLower(judul)
	kr := 0
	kn := len(ideas) - 1

	for kr <= kn {
		mid := (kr + kn) / 2
		midJudul := strings.ToLower(ideas[mid].Judul)
		if midJudul == judul {
			return &ideas[mid]
		} else if midJudul < judul {
			kr = mid + 1
		} else {
			kn = mid - 1
		}
	}
	return nil
}

func UrutkanVotingAscSelection() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		idx_min := i - 1
		j := i
		for j < n {
			if ideas[idx_min].Voting > ideas[j].Voting {
				idx_min = j
			}
			j = j + 1
		}
		// Tukar posisi
		temp := ideas[idx_min]
		ideas[idx_min] = ideas[i-1]
		ideas[i-1] = temp

		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan Voting secara ascending. 📈")
}

func UrutkanVotingDescSelection() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		idx_max := i - 1
		j := i
		for j < n {
			if ideas[idx_max].Voting < ideas[j].Voting {
				idx_max = j
			}
			j = j + 1
		}
		// Tukar posisi
		temp := ideas[idx_max]
		ideas[idx_max] = ideas[i-1]
		ideas[i-1] = temp

		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan Voting secara descending. 📉")
}

func UrutkanVotingAscInsertion() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		j := i
		temp := ideas[j]
		for j > 0 && temp.Voting < ideas[j-1].Voting {
			ideas[j] = ideas[j-1]
			j = j - 1
		}
		ideas[j] = temp
		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan voting (ascending) menggunakan Insertion Sort. 📈")
}

func UrutkanVotingDescInsertion() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		j := i
		temp := ideas[j]
		for j > 0 && temp.Voting > ideas[j-1].Voting {
			ideas[j] = ideas[j-1]
			j = j - 1
		}
		ideas[j] = temp
		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan voting (descending) menggunakan Insertion Sort. 📉")
}

func UrutkanTanggalAscSelection() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		idx_min := i - 1
		j := i
		for j < n {
			if ideas[j].Tanggal < ideas[idx_min].Tanggal {
				idx_min = j
			}
			j = j + 1
		}
		ideas[idx_min], ideas[i-1] = ideas[i-1], ideas[idx_min]
		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan tanggal (ascending) menggunakan Selection Sort. 📅⬆️")
}

func UrutkanTanggalDescSelection() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		idx_max := i - 1
		j := i
		for j < n {
			if ideas[j].Tanggal > ideas[idx_max].Tanggal {
				idx_max = j
			}
			j = j + 1
		}
		ideas[idx_max], ideas[i-1] = ideas[i-1], ideas[idx_max]
		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan tanggal (descending) menggunakan Selection Sort. 📅⬇️")
}

func UrutkanTanggalAscInsertion() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		j := i
		temp := ideas[j]
		for j > 0 && temp.Tanggal < ideas[j-1].Tanggal {
			ideas[j] = ideas[j-1]
			j = j - 1
		}
		ideas[j] = temp
		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan tanggal (ascending) menggunakan Insertion Sort. 📅⬆️")
}

func UrutkanTanggalDescInsertion() {
	n := len(ideas)
	i := 1
	for i <= n-1 {
		j := i
		temp := ideas[j]
		for j > 0 && temp.Tanggal > ideas[j-1].Tanggal {
			ideas[j] = ideas[j-1]
			j = j - 1
		}
		ideas[j] = temp
		i = i + 1
	}
	fmt.Println("✅ Ide berhasil diurutkan berdasarkan tanggal (descending) menggunakan Insertion Sort. 📅⬇️")
}

func UrutkanIDAsc() {
	for i := 1; i < len(ideas); i++ {
		temp := ideas[i]
		j := i
		for j > 0 && temp.ID < ideas[j-1].ID {
			ideas[j] = ideas[j-1]
			j--
		}
		ideas[j] = temp
	}
	fmt.Println("✅ Data berhasil diurutkan kembali berdasarkan ID (urutan awal penambahan). 🔄")
}

func TampilkanSemuaIde() {
	if len(ideas) == 0 {
		fmt.Println("📭 Belum ada ide yang ditambahkan.")
		return
	}

	fmt.Println("\n📋 Daftar Semua Ide:")
	for _, ide := range ideas {
		fmt.Println("------------------------------")
		fmt.Printf("🆔 ID       : %d\n", ide.ID)
		fmt.Printf("📝 Judul    : %s\n", ide.Judul)
		fmt.Printf("🏷️ Kategori : %s\n", ide.Kategori)
		fmt.Printf("📅 Tanggal  : %s\n", ide.Tanggal)
		fmt.Printf("👍 Voting   : %d\n", ide.Voting)
	}
	fmt.Println("------------------------------")
}

func main() {
	var fitur int
	for {
		fmt.Println("\n📋 Menu:")
		fmt.Println("1️⃣ Tambahkan Ide")
		fmt.Println("2️⃣ Edit Ide")
		fmt.Println("3️⃣ Hapus Ide")
		fmt.Println("4️⃣ Voting Ide")
		fmt.Println("5️⃣ Cari Ide")
		fmt.Println("6️⃣ Urutkan Ide Berdasarkan Voting")
		fmt.Println("7️⃣ Urutan Ide Berdasarkan Tanggal")
		fmt.Println("8️⃣ Reset Urutan Berdasarkan ID")
		fmt.Println("9️⃣ Tampilkan Semua Ide")
		fmt.Println("0️⃣ Keluar")
		fmt.Print("👉 Pilih menu: ")
		fmt.Scanln(&fitur)

		switch fitur {
		case 1:
			Tambahide()

		case 2:
			EditideByID()

		case 3:
			HapusideByID()

		case 4:
			VotingByID()

		case 5:
			reader := bufio.NewReader(os.Stdin)
			var metode, tipe int
			fmt.Println("🔎 Cari berdasarkan:")
			fmt.Println("1️⃣ ID")
			fmt.Println("2️⃣ Judul")
			fmt.Print("👉 Pilih tipe pencarian: ")
			fmt.Scanln(&tipe)

			fmt.Println("⚙️ Pilih metode pencarian:")
			fmt.Println("1️⃣ Sequential Search")
			fmt.Println("2️⃣ Binary Search")
			fmt.Print("👉 Pilih metode: ")
			fmt.Scanln(&metode)

			switch tipe {
			case 1:
				fmt.Print("🆔 Masukkan ID: ")
				var id int
				fmt.Scanln(&id)
				var result *Ide
				if metode == 1 {
					result = SearchSequentialByID(id)
				} else {
					sort.Slice(ideas, func(i, j int) bool {
						return ideas[i].ID < ideas[j].ID
					})
					result = SearchBinaryByID(id)
				}

				if result != nil {
					fmt.Println("✅ Ide ditemukan:")
					fmt.Printf("🆔 ID      : %d\n📝 Judul    : %s\n🏷️ Kategori : %s\n📅 Tanggal  : %s\n👍 Voting   : %d\n\n",
						result.ID, result.Judul, result.Kategori, result.Tanggal, result.Voting)
				} else {
					fmt.Println("❌ Ide tidak ditemukan. 😞")
				}

			case 2:
				fmt.Print("📝 Masukkan judul ide: ")
				judul, _ := reader.ReadString('\n')
				judul = strings.TrimSpace(judul)

				var result *Ide
				if metode == 1 {
					result = SearchSequentialByJudul(judul)
				} else {
					sort.Slice(ideas, func(i, j int) bool {
						return strings.ToLower(ideas[i].Judul) < strings.ToLower(ideas[j].Judul)
					})
					result = SearchBinaryByJudul(judul)
				}

				if result != nil {
					fmt.Println("✅ Ide ditemukan:")
					fmt.Printf("🆔 ID      : %d\n📝 Judul    : %s\n🏷️ Kategori : %s\n📅 Tanggal  : %s\n👍 Voting   : %d\n\n",
						result.ID, result.Judul, result.Kategori, result.Tanggal, result.Voting)
				} else {
					fmt.Println("❌ Ide tidak ditemukan. 😞")
				}

			default:
				fmt.Println("❗ Pilihan tidak valid. ⚠️")
			}

		case 6:
			var urutan, metode int
			fmt.Println("📋 Pilih urutan voting:")
			fmt.Println("1️⃣ Ascending")
			fmt.Println("2️⃣ Descending")
			fmt.Print("👉 Masukkan pilihan: ")
			fmt.Scanln(&urutan)

			fmt.Println("🔼 Pilih metode sorting:")
			fmt.Println("1️⃣ Selection Sort")
			fmt.Println("2️⃣ Insertion Sort")
			fmt.Print("👉 Masukkan pilihan: ")
			fmt.Scanln(&metode)

			if urutan == 1 {
				if metode == 1 {
					UrutkanVotingAscSelection()
				} else if metode == 2 {
					UrutkanVotingAscInsertion()
				} else {
					fmt.Println("❌ Metode sorting tidak valid. ⚠️")
				}
			} else if urutan == 2 {
				if metode == 1 {
					UrutkanVotingDescSelection()
				} else if metode == 2 {
					UrutkanVotingDescInsertion()
				} else {
					fmt.Println("❌ Metode sorting tidak valid. ⚠️")
				}
			} else {
				fmt.Println("❌ Pilihan urutan tidak valid. ⚠️")
			}

		case 7:
			var urutan, metode int
			fmt.Println("📅 Urutkan ide berdasarkan tanggal:")
			fmt.Println("1️⃣ Gunakan Selection Sort")
			fmt.Println("2️⃣ Gunakan Insertion Sort")
			fmt.Print("👉 Masukkan pilihan metode: ")
			fmt.Scanln(&metode)

			fmt.Println("🔼 Pilih urutan pengurutan:")
			fmt.Println("1️⃣ Ascending (dari paling lama)")
			fmt.Println("2️⃣ Descending (dari paling baru)")
			fmt.Print("👉 Masukkan pilihan urutan: ")
			fmt.Scanln(&urutan)

			if metode == 1 {
				if urutan == 1 {
					UrutkanTanggalAscSelection()
					fmt.Println("✅ Berhasil mengurutkan ide berdasarkan tanggal secara ascending 📅⬆️")
				} else if urutan == 2 {
					UrutkanTanggalDescSelection()
					fmt.Println("✅ Berhasil mengurutkan ide berdasarkan tanggal secara descending 📅⬇️")
				} else {
					fmt.Println("❌ Pilihan urutan tidak valid. ⚠️")
				}
			} else if metode == 2 {
				if urutan == 1 {
					UrutkanTanggalAscInsertion()
					fmt.Println("✅ Berhasil mengurutkan ide berdasarkan tanggal secara ascending 📅⬆️")
				} else if urutan == 2 {
					UrutkanTanggalDescInsertion()
					fmt.Println("✅ Berhasil mengurutkan ide berdasarkan tanggal secara descending 📅⬇️")
				} else {
					fmt.Println("❌ Pilihan urutan tidak valid. ⚠️")
				}
			} else {
				fmt.Println("❌ Metode pengurutan tidak valid. ⚠️")
			}

		case 8:
			UrutkanIDAsc()

		case 9:
			TampilkanSemuaIde()

		case 0:
			fmt.Println("👋 Keluar dari aplikasi. Sampai jumpa! ✨")
			return

		default:
			fmt.Println("❗ Pilihan tidak valid. Silakan coba lagi. 🔄")
		}
	}
}
