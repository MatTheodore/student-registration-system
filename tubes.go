package main

import (
	"fmt"
)

const NMAXMahasiswa int = 100
const NMAXJurusan int = 10

type daftarMahasiswa [NMAXMahasiswa]mahasiswa
type daftarJurusan [NMAXJurusan]jurusan

type mahasiswa struct {
	nama, jurusan, username, password string
	nilai                             int
	diterima, aktif                   bool
	jumlah                            int
}

type infotype struct {
	data mahasiswa
}

type jurusan struct {
	namaJurusan string
}

type infoMahasiswa struct {
	data            [NMAXMahasiswa]mahasiswa
	jumlahMahasiswa int
	counterRegister int
}

var dataMahasiswa infoMahasiswa

type infoJurusan struct {
	data          [NMAXJurusan]jurusan
	jumlahJurusan int
}

var dataJurusan infoJurusan

//FLOW APLIKASI

func main() {
	mainFlow()
}

func mainFlow() {
	pilihRoleLaluEksekusi()
}

func pilihRoleLaluEksekusi() {
	var pilihan int
	header()
	pilihLaluMasukMenu(&pilihan)
}

func header() {
	fmt.Println("	Selamat datang di	")
	fmt.Println("Aplikasi Pendaftaran Mahasiswa Baru")

	fmt.Println("=====================================")
}

func pilihLaluMasukMenu(pilihan *int) {
	fmt.Println("	Pilih tipe pengguna: ")
	fmt.Println("	1. Admin")
	fmt.Println("	2. Mahasiswa")
	fmt.Print("Opsi: ")
	fmt.Scan(pilihan)
	pilihOpsi(pilihan)
}

func loginAsAdmin() {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	if username == "admin" && password == "admin" {
		fmt.Println("Login berhasil")
		menuAdmin()
	} else {
		fmt.Println("Username atau password salah")
		loginAsAdmin()
	}
}

// Mahasiswa 1st Phase
func registerMahasiswa() {
	var nama, username, password string
	var jurusan, pilihan string

	tampilkanDaftarJurusan()

	fmt.Print("Pilih jurusan: ")
	fmt.Scan(&jurusan)
	if jurusan == "kembali" {
		opsiMahasiswa()
	} else {
		dataJurusan.data[dataJurusan.jumlahJurusan].namaJurusan = jurusan
	}
	if findJurusanifExist(jurusan) {
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&nama)
		dataMahasiswa.data[dataMahasiswa.jumlahMahasiswa].nama = nama
		fmt.Print("Buat Username: ")
		fmt.Scan(&username)
		for searchUsernameifExist(username) != username || dataMahasiswa.counterRegister < 1 {
			dataMahasiswa.data[dataMahasiswa.jumlahMahasiswa].username = username
			fmt.Print("Buat Sandi: ")
			fmt.Scan(&password)
			dataMahasiswa.data[dataMahasiswa.jumlahMahasiswa].password = password
			dataMahasiswa.jumlahMahasiswa++
			dataMahasiswa.counterRegister++
			fmt.Println("Registrasi berhasil, selamat datang", nama)
			opsiMahasiswa()
		}
		fmt.Println("Username telah terdaftar")
		fmt.Println("Ingin Login? (y/n)")
		fmt.Scan(&pilihan)
		if pilihan == "y" {
			loginAsMahasiswa()
		} else {
			registerMahasiswa()
		}

	} else {
		fmt.Println("Jurusan tidak ditemukan")
		registerMahasiswa()
	}

}

func searchUsernameifExist(username string) string {
	var i int
	for i = 0; i < dataMahasiswa.jumlahMahasiswa; i++ {
		if dataMahasiswa.data[i].username == username {
			return username
		}
	}
	return username
}

func loginAsMahasiswa() {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	for i := 0; i < dataMahasiswa.jumlahMahasiswa; i++ {
		if username == dataMahasiswa.data[i].username && password == dataMahasiswa.data[i].password {
			fmt.Println("Login berhasil")
			dataMahasiswa.data[i].aktif = true
			fmt.Println("")
			fiturMahasiswa()
		} else {
			fmt.Println("Username atau password salah")
			loginAsMahasiswa()
		}
	}
}

func fiturMahasiswa() {
	var opsi int
	fmt.Println("1. Input Nilai")
	fmt.Println("2. Edit Profil")
	fmt.Println("3. Pengumuman Kelulusan")
	fmt.Println("4. Logout")
	fmt.Print("Opsi: ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		inputNilai()
	} else if opsi == 2 {
		editProfil()
	} else if opsi == 3 {
		pengumumanKelulusan()
	} else if opsi == 4 {
		opsiMahasiswa()
	} else {
		fmt.Println("Pilihan tidak tersedia")
		fiturMahasiswa()
	}
}

func inputNilai() {
	var nilai int
	var namaCari string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan(&namaCari)
	var indeksUbah int = searchMahasiswa(namaCari)
	if indeksUbah != -1 {
		fmt.Print("Masukkan rata-rata nilai: ")
		fmt.Scan(&nilai)
		dataMahasiswa.data[indeksUbah].nilai = nilai
		fmt.Println("Input nilai berhasil")
		fiturMahasiswa()
	}
}

func editProfil() {
	var namaCari string
	var namaBaru string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan(&namaCari)
	var indeksUbah int = searchMahasiswa(namaCari)
	if indeksUbah != -1 {
		fmt.Print("Nama semula: ", dataMahasiswa.data[indeksUbah].nama, "\n")
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&namaBaru)
		dataMahasiswa.data[indeksUbah].nama = namaBaru
	}
}

func pengumumanKelulusan() {
	var namaCari string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan(&namaCari)
	var indeksCari int = searchMahasiswa(namaCari)
	if indeksCari != -1 {
		if dataMahasiswa.data[indeksCari].nilai >= 80 {
			fmt.Println("Selamat Anda lulus")
		} else {
			fmt.Println("Anda tidak lulus")
		}
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func pilihOpsi(pilihan *int) {
	if *pilihan == 1 {
		loginAsAdmin()
	} else if *pilihan == 2 {
		opsiMahasiswa()
	} else {
		fmt.Println("Pilihan tidak tersedia")
		fmt.Println("")
		pilihLaluMasukMenu(pilihan)
	}
}

func menuAdmin() {
	headerAdmin()
	opsiAdmin()
}

func headerAdmin() {
	fmt.Println("=====================================")
	fmt.Println("		Menu Admin")
	fmt.Println("=====================================")
}

func opsiAdmin() {
	var opsi int
	fmt.Println("	1. Tambah data mahasiswa")
	fmt.Println("	2. Ubah data mahasiswa")
	fmt.Println("	3. Hapus data mahasiswa")
	fmt.Println("	4. Tambah data jurusan")
	fmt.Println("	5. Ubah data jurusan")
	fmt.Println("	6. Hapus data jurusan")
	fmt.Println("	7. Tampilkan Data Mahasiswa")
	fmt.Println("	8. Tampilkan Data Jurusan")
	fmt.Println("	9. Kembali")
	fmt.Println("=====================================")

	fmt.Print("Opsi: ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		inputDataMahasiswa()
		fmt.Println("Data berhasil ditambahkan")
		opsiAdmin()
	} else if opsi == 2 {
		ubahDataMahasiswa()
		opsiAdmin()
	} else if opsi == 3 {
		hapusDataMahasiswa()
		opsiAdmin()
	} else if opsi == 4 {
		inputDataJurusan()
		fmt.Println("Data berhasil ditambahkan")
		opsiAdmin()
	} else if opsi == 5 {
		ubahDataJurusan()
		opsiAdmin()
	} else if opsi == 6 {
		hapusDataJurusan()
		opsiAdmin()
	} else if opsi == 7 {
		tampilkanDataMahasiswa()
		opsiAdmin()
	} else if opsi == 8 {
		tampilkanDataJurusan()
		opsiAdmin()
	} else if opsi == 9 {
		pilihRoleLaluEksekusi()
		//opsiAdmin()
	} else {
		fmt.Println("Pilihan tidak tersedia")
		//opsiAdmin()
	}
}

func opsiMahasiswa() {
	var opsi int
	fmt.Println("=====================================")
	fmt.Println("	Menu Mahasiswa		")
	fmt.Println("=====================================")
	fmt.Println("	1. Daftar")
	fmt.Println("	2. Masuk")
	fmt.Println("	3. Kembali")
	fmt.Println("=====================================")

	fmt.Print("Opsi: ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		registrasiMahasiswa()
	} else if opsi == 2 {
		loginAsMahasiswa()
	} else if opsi == 3 {
		pilihRoleLaluEksekusi()
	} else {
		fmt.Println("Pilihan tidak tersedia")
		opsiMahasiswa()
	}
}

func menuMahasiswa() {
	opsiMahasiswa()

}

// OPSI ADMIN
// MHS
func inputDataMahasiswa() {
	inputDM()
}

func ubahDataMahasiswa() {
	ubahDM()
}

func hapusDataMahasiswa() {
	hapusDM()
}

func tampilkanDataMahasiswa() {
	var opsi int
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Nilai")
	fmt.Println("3. Berdasarkan Jurusan")
	fmt.Scan(&opsi)

	if opsi == 1 {
		mahasiswaTerurutBerdasarkanNamaAscending()
		tampilkanDaftarMahasiswa()
	} else if opsi == 2 {
		mahasiswaTerurutBerdasarkanNilaiDescending()
		tampilkanDaftarMahasiswa()
	} else if opsi == 3 {
		mahasiswaTerurutBerdasarkanJurusanAscending()
		tampilkanDaftarMahasiswa()
	} else {
		fmt.Println("Pilihan tidak tersedia")
		tampilkanDataMahasiswa()
	}

}

// JRSN
func inputDataJurusan() {
	inputDJ()
}

func ubahDataJurusan() {
	ubahDJ()
}

func hapusDataJurusan() {
	hapusDJ()
}

func tampilkanDataJurusan() {
	tampilkanDaftarJurusan()
}

//INTERFACE OPSI MAHASISWA

func registrasiMahasiswa() {
	registerMahasiswa()
	opsiMahasiswa()
}

// func loginMahasiswa(){
// 	masukMahasiswa()
// }

// MAHASISWA

// func masukMahasiswa( ){
// 	var user, password string

// }

//ADMIN

func isMahasiswaDiterimaDitolak() {
	var nilaiUntukDiterima int
	fmt.Print("Masukkan nilai untuk diterima: ")
	fmt.Scanln(&nilaiUntukDiterima)

	for i := 0; i < dataMahasiswa.jumlahMahasiswa; i++ {
		if dataMahasiswa.data[i].nilai >= nilaiUntukDiterima {
			dataMahasiswa.data[i].diterima = true
		} else {
			dataMahasiswa.data[i].diterima = false
		}
	}
}

//MAHASISWA

func inputDM() {
	var idx int
	var jumlah int
	var nama string
	var jurusan string
	var nilai int
	fmt.Println("Jurusan yang tersedia: ")
	for i := 0; i < dataJurusan.jumlahJurusan; i++ {
		fmt.Println(i+1, ".", dataJurusan.data[i].namaJurusan, " ")
	}

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	idx = dataMahasiswa.jumlahMahasiswa
	for i := 0; i < jumlah; i++ {
		fmt.Println("Mahasiswa ke-", i+1)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&nama)
		dataMahasiswa.data[idx].nama = nama
		fmt.Print("Masukkan Jurusan: ")
		fmt.Scan(&jurusan)
		for searchJurusan(jurusan) == -1 {
			fmt.Print("Jurusan tidak tersedia\n")
			fmt.Print("Masukkan Jurusan: ")
			fmt.Scan(&jurusan)
			if jurusan == "kembali" {
				opsiAdmin()
			}
		}
		dataMahasiswa.data[idx].jurusan = jurusan
		fmt.Print("Masukkan Nilai: ")
		fmt.Scan(&nilai)
		dataMahasiswa.data[idx].nilai = nilai
		idx++
	}
	dataMahasiswa.jumlahMahasiswa = idx
	fmt.Println("idx:", idx)
}

func searchMahasiswa(nama string) int { //sequential/linear search
	var i int
	for i < dataMahasiswa.jumlahMahasiswa && (dataMahasiswa.data[i].nama != "" && dataMahasiswa.data[i].jurusan != "" && dataMahasiswa.data[i].nilai != 0) {
		if dataMahasiswa.data[i].nama == nama {
			return i
		}
		i++
	}
	return -1
}

func ubahDM() {
	var namaCari string
	var namaBaru string
	var jurusan string
	fmt.Print("Masukkan nama mahasiswa yang ingin diubah: ")
	fmt.Scan(&namaCari)
	var indeksUbah int = searchMahasiswa(namaCari)
	if indeksUbah != -1 {
		fmt.Print("Nama semula: ", dataMahasiswa.data[indeksUbah].nama, "\n")
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&namaBaru)
		dataMahasiswa.data[indeksUbah].nama = namaBaru
		fmt.Print("Jurusan semula: ", dataMahasiswa.data[indeksUbah].jurusan, "\n")
		fmt.Print("Masukkan jurusan baru: ")
		fmt.Scan(&jurusan)
		for searchJurusan(jurusan) == -1 {
			fmt.Print("Jurusan tidak tersedia\n")
			fmt.Print("Masukkan Jurusan: ")
			fmt.Scan(&jurusan)
			if jurusan == "kembali" {
				opsiAdmin()
			}
		}
		dataMahasiswa.data[indeksUbah].jurusan = jurusan
		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusDM() {
	var nama string
	fmt.Print("Masukkan nama mahasiswa yang ingin dihapus: ")
	fmt.Scan(&nama)
	var indeksHapus int = searchMahasiswa(nama)
	if dataMahasiswa.data[indeksHapus].nama == nama {
		for i := 0; i < dataMahasiswa.jumlahMahasiswa && ((dataMahasiswa.data[i].nama != "") && (dataMahasiswa.data[i].jurusan != "") && (dataMahasiswa.data[i].nilai != 0)); i++ {
			if dataMahasiswa.data[i].nama == nama {
				for j := i; j < dataMahasiswa.jumlahMahasiswa && ((dataMahasiswa.data[i].nama != "") && (dataMahasiswa.data[i].jurusan != "") && (dataMahasiswa.data[i].nilai != 0)); j++ {
					dataMahasiswa.data[j] = dataMahasiswa.data[j+1]
				}
			}
		}
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

// LAYOUT

//JURUSAN

func inputDJ() {
	var jumlah int
	var nama string
	var idx int
	idx = dataJurusan.jumlahJurusan
	fmt.Print("Masukkan jumlah jurusan: ")
	fmt.Scan(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Println("Jurusan ke-", i+1)
		fmt.Print("Masukkan Nama Jurusan: ")
		fmt.Scan(&nama)
		dataJurusan.data[idx].namaJurusan = nama
		idx++
	}
	dataJurusan.jumlahJurusan = idx
	fmt.Println("idx:", idx)
}

func searchJurusan(nama string) int {
	low := 0
	high := dataJurusan.jumlahJurusan - 1

	for low <= high {
		mid := (low + high) / 2

		if dataJurusan.data[mid].namaJurusan == nama {
			return mid
		} else if dataJurusan.data[mid].namaJurusan > nama {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func findJurusanifExist(nama string) bool {
	var i int
	for i < dataJurusan.jumlahJurusan && (dataJurusan.data[i].namaJurusan != "") {
		if dataJurusan.data[i].namaJurusan == nama {
			return true
		}
		i++
	}
	return false
}

func ubahDJ() {
	var namaCari string
	var namaBaru string
	fmt.Print("Masukkan nama jurusan yang ingin diubah: ")
	fmt.Scan(&namaCari)
	var indeksUbah int = searchJurusan(namaCari)
	if indeksUbah != -1 {
		fmt.Print("Nama semula: ", dataJurusan.data[indeksUbah].namaJurusan, "\n")
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&namaBaru)
		dataJurusan.data[indeksUbah].namaJurusan = namaBaru
		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusDJ() {
	var nama string
	fmt.Print("Masukkan nama jurusan yang ingin dihapus: ")
	fmt.Scan(&nama)
	var indeksHapus int = searchJurusan(nama)
	if dataJurusan.data[indeksHapus].namaJurusan == nama {
		for i := 0; i < dataJurusan.jumlahJurusan && (dataJurusan.data[i].namaJurusan != ""); i++ {
			if dataJurusan.data[i].namaJurusan == nama {
				for j := i; j < dataJurusan.jumlahJurusan-1 && (dataJurusan.data[i].namaJurusan != ""); j++ {
					dataJurusan.data[j] = dataJurusan.data[j+1]
				}
			}
		}
		dataJurusan.jumlahJurusan--
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

//TAMPIL DATA MAHASISWA

func tampilkanMahasiswaBerdasarkanJurusan(jumlahMahasiswa *int, jurusan string) {
	fmt.Println("Data Mahasiswa Jurusan ", jurusan)
	fmt.Println("============================")
	var i int
	for i < *jumlahMahasiswa {
		if dataMahasiswa.data[i].jurusan == jurusan {
			fmt.Println("Nama: ", dataMahasiswa.data[i].nama)
			fmt.Println("Jurusan: ", dataMahasiswa.data[i].jurusan)
			fmt.Println("Nilai: ", dataMahasiswa.data[i].nilai)
		}
		i++
	}
}

func tampilkanMahasiswaYangDitolakBerdasarkanJurusan(jurusan string) {
	fmt.Println("Data Mahasiswa Jurusan ", jurusan)
	fmt.Println("============================")
	var i int
	for i < *&dataMahasiswa.jumlahMahasiswa {
		if dataMahasiswa.data[i].jurusan == jurusan && dataMahasiswa.data[i].diterima == false {
			fmt.Println("Nama: ", dataMahasiswa.data[i].nama)
			fmt.Println("Jurusan: ", dataMahasiswa.data[i].jurusan)
			fmt.Println("Nilai: ", dataMahasiswa.data[i].nilai)
		}
		i++
	}
}

func tampilkanMahasiswaYangDiterimaBerdasarkanJurusan(jurusan string) {
	fmt.Println("Data Mahasiswa Jurusan ", jurusan)
	fmt.Println("============================")
	var i int
	for i < *&dataMahasiswa.jumlahMahasiswa {
		if dataMahasiswa.data[i].jurusan == jurusan && dataMahasiswa.data[i].diterima {
			fmt.Println("Nama: ", dataMahasiswa.data[i].nama)
			fmt.Println("Jurusan: ", dataMahasiswa.data[i].jurusan)
			fmt.Println("Nilai: ", dataMahasiswa.data[i].nilai)
		}
		i++
	}
}

func tampilkanDaftarMahasiswa() {
	fmt.Println("Daftar Mahasiswa:")
	fmt.Println("============================")
	for i := 0; i < dataMahasiswa.jumlahMahasiswa; i++ {
		if dataMahasiswa.data[i].nama != "" {
			fmt.Println("Nama:", dataMahasiswa.data[i].nama)
			fmt.Println("Jurusan:", dataMahasiswa.data[i].jurusan)
			fmt.Println("Nilai:", dataMahasiswa.data[i].nilai)
			fmt.Println("============================")
		}
	}
}

func mahasiswaTerurutBerdasarkanNilaiAscending() {
	for i := 1; i < dataMahasiswa.jumlahMahasiswa; i++ {
		key := dataMahasiswa.data[i]
		j := i - 1

		for j >= 0 && dataMahasiswa.data[j].nilai > key.nilai {
			dataMahasiswa.data[j+1] = dataMahasiswa.data[j]
			j--
		}

		dataMahasiswa.data[j+1] = key
	}
}

func mahasiswaTerurutBerdasarkanNilaiDescending() {
	for i := 1; i < dataMahasiswa.jumlahMahasiswa; i++ {
		key := dataMahasiswa.data[i]
		j := i - 1

		for j >= 0 && dataMahasiswa.data[j].nilai < key.nilai {
			dataMahasiswa.data[j+1] = dataMahasiswa.data[j]
			j--
		}
		dataMahasiswa.data[j+1] = key
	}
}

func mahasiswaTerurutBerdasarkanJurusanAscending() {
	for i := 0; i < dataMahasiswa.jumlahMahasiswa-1; i++ {
		minIndex := i

		for j := i + 1; j < dataMahasiswa.jumlahMahasiswa; j++ {
			if dataMahasiswa.data[j].jurusan < dataMahasiswa.data[minIndex].jurusan {
				minIndex = j
			}
		}

		dataMahasiswa.data[i], dataMahasiswa.data[minIndex] = dataMahasiswa.data[minIndex], dataMahasiswa.data[i]
	}
}

func mahasiswaTerurutBerdasarkanJurusanDescending() {
	for i := 0; i < dataMahasiswa.jumlahMahasiswa-1; i++ {
		maxIndex := i

		for j := i + 1; j < dataMahasiswa.jumlahMahasiswa; j++ {
			if dataMahasiswa.data[j].jurusan > dataMahasiswa.data[maxIndex].jurusan {
				maxIndex = j
			}
		}

		dataMahasiswa.data[i], dataMahasiswa.data[maxIndex] = dataMahasiswa.data[maxIndex], dataMahasiswa.data[i]
	}
}

func mahasiswaTerurutBerdasarkanNamaAscending() {
	for i := 0; i < dataMahasiswa.jumlahMahasiswa-1; i++ {
		for j := 0; j < dataMahasiswa.jumlahMahasiswa-i-1; j++ {
			if dataMahasiswa.data[j].nama > dataMahasiswa.data[j+1].nama {
				dataMahasiswa.data[j], dataMahasiswa.data[j+1] = dataMahasiswa.data[j+1], dataMahasiswa.data[j]
			}
		}
	}
}

func mahasiswaTerurutBerdasarkanNamaDescending() {
	for i := 0; i < dataMahasiswa.jumlahMahasiswa-1; i++ {
		for j := 0; j < dataMahasiswa.jumlahMahasiswa-i-1; j++ {
			if dataMahasiswa.data[j].nama < dataMahasiswa.data[j+1].nama {
				dataMahasiswa.data[j], dataMahasiswa.data[j+1] = dataMahasiswa.data[j+1], dataMahasiswa.data[j]
			}
		}
	}
}

// TAMPIL DATA JURUSAN

func tampilkanDaftarJurusan() {
	fmt.Println("\nDaftar Jurusan:")
	fmt.Println("============================")
	for i := 0; i < dataJurusan.jumlahJurusan; i++ {
		if dataJurusan.data[i].namaJurusan != "" {
			fmt.Println("Nama Jurusan: ", dataJurusan.data[i].namaJurusan)
			fmt.Println("============================")
		}
	}
}
