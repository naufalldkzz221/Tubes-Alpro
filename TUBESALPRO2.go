package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Soal struct {
	Pertanyaan string
	Pilihan    [4]string
	Jawaban    int
	BenarCnt   int
	SalahCnt   int
}

type Peserta struct {
	Nama string
	Skor int
}

const MAX_SOAL = 100
const MAX_PESERTA = 100

var bankSoal [MAX_SOAL]Soal
var daftarPeserta [MAX_PESERTA]Peserta
var jumlahSoal, jumlahPeserta int

func main() {
	rand.Seed(time.Now().UnixNano())
	inisialisasiSoal()

	for {
		clearScreen()
		printMenu("Aplikasi Quiz - Who Wants to Be a Millionaire", []string{
			"Admin",
			"Peserta",
			"METU SIT",
		})

		var pilihan int
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			panelAdmin()
		case 2:
			panelPeserta()
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi. Babay Ketemu dilain waktu!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func inisialisasiSoal() {
	bankSoal[0] = Soal{"Apa ibu kota Prancis?", [4]string{"Paris", "Berlin", "Madrid", "Roma"}, 1, 0, 0}
	bankSoal[1] = Soal{"Berapa hasil 2 + 2?", [4]string{"3", "4", "5", "6"}, 2, 0, 0}
	bankSoal[2] = Soal{"Siapa penulis 'Hamlet'?", [4]string{"Shakespeare", "Hemingway", "Tolkien", "Austen"}, 1, 0, 0}
	jumlahSoal = 3
}

func panelAdmin() {
	for {
		clearScreen()
		printMenu("Panel Admin", []string{
			"Lihat Soal",
			"Tambah Soal",
			"Edit Soal",
			"Hapus Soal",
			"Lihat Statistik Soal",
			"Kembali ke Menu Utama",
		})

		var pilihan int
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatSoal()
		case 2:
			tambahSoal()
		case 3:
			editSoal()
		case 4:
			hapusSoal()
		case 5:
			lihatStatistikSoal()
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func lihatSoal() {
	clearScreen()
	fmt.Println("Daftar Soal:")
	for i := 0; i < jumlahSoal; i++ {
		fmt.Printf("%d. %s\n", i+1, bankSoal[i].Pertanyaan)
		for j, opsi := range bankSoal[i].Pilihan {
			fmt.Printf("   %d. %s\n", j+1, opsi)
		}
	}
	tekanEnter()
}

func tambahSoal() {
	if jumlahSoal >= MAX_SOAL {
		fmt.Println("Bank soal kebek. Tidak dapat menambah soal baru.")
		tekanEnter()
		return
	}

	var soal Soal
	fmt.Print("Masukkan pertanyaan: ")
	fmt.Scanln(&soal.Pertanyaan)
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan pilihan %d: ", i+1)
		fmt.Scanln(&soal.Pilihan[i])
	}
	fmt.Print("Masukkan nomor pilihan yang benar (1-4): ")
	fmt.Scanln(&soal.Jawaban)

	bankSoal[jumlahSoal] = soal
	jumlahSoal++
	fmt.Println("Soal berhasil ditambahkan!")
	tekanEnter()
}

func editSoal() {
	lihatSoal()
	fmt.Print("Masukkan nomor soal yang ingin diedit: ")
	var nomor int
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > jumlahSoal {
		fmt.Println("Nomor soal tidak valid.")
		tekanEnter()
		return
	}

	tambahSoal()
	bankSoal[nomor-1] = bankSoal[jumlahSoal-1]
	jumlahSoal--
	fmt.Println("Soal berhasil diedit!")
	tekanEnter()
}

func hapusSoal() {
	lihatSoal()
	fmt.Print("Masukkan nomor soal yang ingin dihapus: ")
	var nomor int
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > jumlahSoal {
		fmt.Println("Nomor soal tidak valid.")
		tekanEnter()
		return
	}

	for i := nomor - 1; i < jumlahSoal-1; i++ {
		bankSoal[i] = bankSoal[i+1]
	}
	jumlahSoal--
	fmt.Println("Soal berhasil dihapus!")
	tekanEnter()
}

func lihatStatistikSoal() {
	clearScreen()
	fmt.Println("Statistik Soal:")
	for i := 0; i < jumlahSoal; i++ {
		fmt.Printf("%d. %s\n   Benar: %d, Salah: %d\n", i+1, bankSoal[i].Pertanyaan, bankSoal[i].BenarCnt, bankSoal[i].SalahCnt)
	}
	tekanEnter()
}

func panelPeserta() {
	if jumlahPeserta >= MAX_PESERTA {
		fmt.Println("Pendaftaran peserta penuh. Tidak dapat mendaftar lebih banyak peserta.")
		tekanEnter()
		return
	}

	var peserta Peserta
	fmt.Print("Masukkan nama peserta: ")
	fmt.Scanln(&peserta.Nama)

	daftarPeserta[jumlahPeserta] = peserta
	jumlahPeserta++
	fmt.Printf("\nSelamat datang, %s! Mari mulai kuis.\n", peserta.Nama)
	kuis(&daftarPeserta[jumlahPeserta-1])
}

func kuis(peserta *Peserta) {
	clearScreen()
	skor := 0
	soalDigunakan := make(map[int]bool)

	for len(soalDigunakan) < jumlahSoal {
		index := rand.Intn(jumlahSoal)
		if soalDigunakan[index] {
			continue
		}
		soalDigunakan[index] = true

		soal := bankSoal[index]
		fmt.Printf("\n%s\n", soal.Pertanyaan)
		for i, opsi := range soal.Pilihan {
			fmt.Printf("%d. %s\n", i+1, opsi)
		}

		fmt.Print("Jawaban Anda (1-4): ")
		var jawaban int
		fmt.Scanln(&jawaban)

		if jawaban == soal.Jawaban {
			fmt.Println("Benar! Kamu Pinter dehh...")
			soal.BenarCnt++
			skor += 10
		} else {
			fmt.Printf("Salah! Jawaban yang benar adalah: %s\n", soal.Pilihan[soal.Jawaban-1])
			soal.SalahCnt++
		}
	}

	peserta.Skor = skor
	fmt.Printf("\nKuis selesai! Skor Anda: %d\n", skor)
	tekanEnter()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape sequence untuk membersihkan layar
}

func printMenu(title string, options []string) {
	border := strings.Repeat("=", len(title)+4)
	fmt.Println(border)
	fmt.Printf("= %s =\n", title)
	fmt.Println(border)
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Println(border)
}

func tekanEnter() {
	fmt.Print("\nPejet Enter Ya Broo buat ngelanjutin...")
	fmt.Scanln()
}
