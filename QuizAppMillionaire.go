package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
	Nama  string
	Email string
	Skor  int
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
		printMenu("Aplikasi Quiz - Who One To Be A Milionare ", []string{
			"Admin",
			"Peserta",
			"Keluar",
		})

		var pilihan int
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			loginAdmin()
		case 2:
			panelPeserta()
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func inisialisasiSoal() {
	bankSoal[0] = Soal{"Apa ibu kota Prancis?", [4]string{"Paris", "Berlin", "Madrid", "Roma"}, 1, 0, 0}
	bankSoal[1] = Soal{"Berapa hasil 2 + 2?", [4]string{"3", "4", "5", "6"}, 2, 0, 0}
	bankSoal[2] = Soal{"Siapa penulis 'Hamlet'?", [4]string{"Shakespeare", "Hemingway", "Tolkien", "Austen"}, 1, 0, 0}
	jumlahSoal = 3
}

func loginAdmin() {
	clearScreen()
	fmt.Println("=== Login Admin ===")
	username := "milionare"
	password := "1234"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan username: ")
	inputUsername, _ := reader.ReadString('\n')
	inputUsername = strings.TrimSpace(inputUsername)

	fmt.Print("Masukkan password: ")
	inputPassword, _ := reader.ReadString('\n')
	inputPassword = strings.TrimSpace(inputPassword)

	if inputUsername == username && inputPassword == password {
		fmt.Println("Login berhasil!")
		tekanEnter()
		panelAdmin()
	} else {
		fmt.Println("Username atau password salah.")
		tekanEnter()
	}
}

// Panel Admin
func panelAdmin() {
	for {
		clearScreen()
		printMenu("Panel Admin", []string{
			"Lihat Soal",
			"Tambah Soal",
			"Edit Soal",
			"Hapus Soal",
			"Lihat Statistik Soal",
			"Urutkan Soal",
			"Cari Soal",
			"Lihat Peserta", // Menu untuk melihat peserta
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
			urutkanSoal()
		case 7:
			cariSoal()
		case 8:
			lihatPeserta()
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
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

// Lihat Peserta
func lihatPeserta() {
	clearScreen()
	fmt.Println("Daftar Peserta yang Telah Mengerjakan Kuis:")
	if jumlahPeserta == 0 {
		fmt.Println("Belum ada peserta yang mengerjakan kuis.")
	} else {
		for i := 0; i < jumlahPeserta; i++ {
			fmt.Printf("%d. Nama: %s, Email: %s, Skor: %d\n", i+1, daftarPeserta[i].Nama, daftarPeserta[i].Email, daftarPeserta[i].Skor)
		}
	}
	tekanEnter()
}

func tambahSoal() {
	if jumlahSoal >= MAX_SOAL {
		fmt.Println("Bank soal penuh.")
		tekanEnter()
		return
	}

	var soal Soal
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan pertanyaan: ")
	pertanyaan, _ := reader.ReadString('\n')
	soal.Pertanyaan = strings.TrimSpace(pertanyaan)

	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan pilihan %d: ", i+1)
		pilihan, _ := reader.ReadString('\n')
		soal.Pilihan[i] = strings.TrimSpace(pilihan)
	}

	for {
		fmt.Print("Masukkan nomor pilihan yang benar (1-4): ")
		_, err := fmt.Scanln(&soal.Jawaban)
		if err == nil && soal.Jawaban >= 1 && soal.Jawaban <= 4 {
			break
		}
		fmt.Println("Input tidak valid. Masukkan angka antara 1 hingga 4.")
	}

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



// Sorting
func urutkanSoal() {
	clearScreen()
	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	var pilihan int
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		selectionSort(true)
	} else if pilihan == 2 {
		selectionSort(false)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
	tekanEnter()
}

func selectionSort(ascending bool) {
	for i := 0; i < jumlahSoal-1; i++ {
		minMax := i
		for j := i + 1; j < jumlahSoal; j++ {
			if ascending && strings.Compare(bankSoal[j].Pertanyaan, bankSoal[minMax].Pertanyaan) < 0 {
				minMax = j
			} else if !ascending && strings.Compare(bankSoal[j].Pertanyaan, bankSoal[minMax].Pertanyaan) > 0 {
				minMax = j
			}
		}
		bankSoal[i], bankSoal[minMax] = bankSoal[minMax], bankSoal[i]
	}
	fmt.Println("Soal berhasil diurutkan.")
}

func lihatStatistikSoal() {
	clearScreen()
	fmt.Println("Statistik Soal:")
	for i := 0; i < jumlahSoal; i++ {
		fmt.Printf("%d. %s\n   Benar: %d, Salah: %d\n", i+1, bankSoal[i].Pertanyaan, bankSoal[i].BenarCnt, bankSoal[i].SalahCnt)
	}
	tekanEnter()
}

// Searching
func cariSoal() {
	clearScreen()
	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (Pastikan data sudah diurutkan)")
	var pilihan int
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilihan)

	fmt.Print("Masukkan kata kunci pencarian: ")
	var keyword string
	fmt.Scanln(&keyword)

	switch pilihan {
	case 1:
		sequentialSearch(keyword)
	case 2:
		binarySearch(keyword)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
	tekanEnter()
}

func sequentialSearch(keyword string) {
	fmt.Println("Hasil pencarian (Sequential Search):")
	found := false
	for i := 0; i < jumlahSoal; i++ {
		if strings.Contains(strings.ToLower(bankSoal[i].Pertanyaan), strings.ToLower(keyword)) {
			fmt.Printf("%d. %s\n", i+1, bankSoal[i].Pertanyaan)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada soal yang ditemukan dengan kata kunci tersebut.")
	}
}

func binarySearch(keyword string) {
	fmt.Println("Hasil pencarian (Binary Search):")

	low := 0
	high := jumlahSoal - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		midValue := strings.ToLower(bankSoal[mid].Pertanyaan)
		searchValue := strings.ToLower(keyword)

		if strings.Contains(midValue, searchValue) {
			fmt.Printf("%d. %s\n", mid+1, bankSoal[mid].Pertanyaan)
			found = true

			// Cari ke kiri dan kanan untuk memastikan semua kecocokan ditemukan
			left := mid - 1
			right := mid + 1
			for left >= 0 && strings.Contains(strings.ToLower(bankSoal[left].Pertanyaan), searchValue) {
				fmt.Printf("%d. %s\n", left+1, bankSoal[left].Pertanyaan)
				left--
			}
			for right < jumlahSoal && strings.Contains(strings.ToLower(bankSoal[right].Pertanyaan), searchValue) {
				fmt.Printf("%d. %s\n", right+1, bankSoal[right].Pertanyaan)
				right++
			}
			break
		} else if midValue < searchValue {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ada soal yang ditemukan dengan kata kunci tersebut.")
	}
}


func panelPeserta() {
	clearScreen()
	reader := bufio.NewReader(os.Stdin)

	// Membaca nama peserta
	fmt.Print("Masukkan nama peserta: ")
	namaPeserta, _ := reader.ReadString('\n')
	namaPeserta = strings.TrimSpace(namaPeserta)

	// Membaca email peserta
	fmt.Print("Masukkan email peserta: ")
	emailPeserta, _ := reader.ReadString('\n')
	emailPeserta = strings.TrimSpace(emailPeserta)

	// Membuat objek peserta baru dan menyimpannya dalam daftarPeserta
	peserta := Peserta{Nama: namaPeserta, Email: emailPeserta, Skor: 0}
	daftarPeserta[jumlahPeserta] = peserta
	jumlahPeserta++

	// Memastikan pesan ditampilkan sebelum lanjut
	fmt.Printf("\nSelamat datang, %s! Email Anda: %s. Mari mulai kuis.\n", namaPeserta, emailPeserta)
	fmt.Println("Persiapkan diri Anda!")
	time.Sleep(2 * time.Second) // Memberi jeda untuk menampilkan pesan

	// Memulai kuis
	kuis(&daftarPeserta[jumlahPeserta-1])
}

func kuis(peserta *Peserta) {
	clearScreen()
	skor := 0
	soalTerjawab := make(map[int]bool)

	// Mengacak urutan soal
	shuffledSoalIndices := rand.Perm(jumlahSoal) // Mengacak indeks soal

	// Mencatat waktu mulai
	startTime := time.Now()

	for i := 0; i < jumlahSoal; i++ {
		index := shuffledSoalIndices[i] // Mendapatkan soal berdasarkan urutan acak
		soal := bankSoal[index]

		// Memastikan soal belum dijawab
		if soalTerjawab[index] {
			continue
		}
		soalTerjawab[index] = true

		// Menampilkan soal dan pilihan jawaban
		fmt.Printf("\n%s\n", soal.Pertanyaan)
		for j, opsi := range soal.Pilihan {
			fmt.Printf("%d. %s\n", j+1, opsi)
		}

		// Membaca input jawaban
		fmt.Print("Jawaban Anda (1-4): ")
		var jawaban int
		fmt.Scanln(&jawaban)

		// Mengevaluasi jawaban
		if jawaban == soal.Jawaban {
			fmt.Println("Benar!")
			bankSoal[index].BenarCnt++
			skor += 10
		} else {
			fmt.Printf("Salah! Jawaban yang benar adalah: %s\n", soal.Pilihan[soal.Jawaban-1])
			bankSoal[index].SalahCnt++
		}
	}

	// Mencatat waktu selesai
	endTime := time.Now()
	duration := endTime.Sub(startTime) // Menghitung durasi pengerjaan kuis

	// Menyimpan skor akhir peserta
	peserta.Skor = skor
	fmt.Printf("\nKuis selesai! Skor Anda: %d\n", skor)
	fmt.Printf("Waktu yang dihabiskan: %v\n", duration)

	// Menambahkan pesan hasil kuis akan dikirimkan ke email
	fmt.Println("Hasil quiz akan kami segera kirim ke email kalian ya.")
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
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	fmt.Scanln()
}