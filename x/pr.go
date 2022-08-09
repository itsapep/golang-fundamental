package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type institute struct {
	instituteName          string
	instituteAddress       string
	instituteAccreditation string
	instituteSince         int
}

type course struct {
	name string
}

type student struct {
	npm, name string
	ins       institute
	courses   []course
}

func main() {
	/*
		Soal 1

		Pada tugas ini kamu diminta untuk membuat sebuah aplikasi basis data mahasiswa sederhana berbasis console dengan kriteria yang sudah ditentukan

		// Model
		type institute struct {
		    instituteName          string
		    instituteAddress       string
		    instituteAccreditation string
		    instituteSince         int
		}

		type course struct {
		    name string
		}

		type student struct {
		    npm, name string
		    ins       institute
		    courses   []course
		}

		1. Bisa menampung sampai dengan 5 mahasiswa (jika lebih ada informasi)
		2. Validasi sesuai dengan screen, jika validasi tidak terpenuhi maka input tersebut akan terus diminta sampai terpenuhi
		3. Design screen harus mempunyai menu tambah mahasiswa, hapus mahasiswa, lihat mahasiswa dan exit
		4. Untuk menu tambah mahasiswa berikan validasi yaitu nama harus minimal 3 dan maksimal 20 karakter
		5. Untuk menu hapus mahasiswa akan otomatis index terakhir yang terhapus atau bisa by index
		6. untuk menu lihat mahasiswa akan ada sub menu lagi yaitu lihat menggunakan index dan lihat semua
		7. untuk menu lihat menggunakan index di input sesuai index data mahasiswa
		8. untuk menu lihat semua akan menampilkan seluruh data mahasiswa yang terdaftar

		repo: challenge-struct-student
	*/
	var mahasiswa []student
	var state, inputan int

	mahasiswa1 := student{
		name: "Gow Lank",
		npm:  "44444",
		ins: institute{
			instituteName:          "Institut Koding Jakarta",
			instituteAddress:       "Jakarta",
			instituteAccreditation: "A",
			instituteSince:         1970,
		},
		courses: []course{
			{
				name: "Basic Programming",
			},
			{
				name: "Cloud Computing",
			},
			{
				name: "Internet of Thing",
			},
		},
	}

	mahasiswa2 := student{
		name: "Phi Tony",
		npm:  "88888",
		ins: institute{
			instituteName:          "Institut Informasi Teknologi",
			instituteAddress:       "Bandung",
			instituteAccreditation: "A",
			instituteSince:         1930,
		},
		courses: []course{
			{
				name: "Jaringan Teknologi",
			},
			{
				name: "Jaringan Teknologi 2",
			},
			{
				name: "Jaringan Teknologi 3",
			},
		},
	}

	tambahMahasiswa(&mahasiswa, mahasiswa1, mahasiswa2)

	mahasiswa3 := student{
		name: "Pradipta Megantara",
		npm:  "11111",
		ins: institute{
			instituteName:          "Universitas Diponegoro",
			instituteAddress:       "Semarang",
			instituteAccreditation: "A",
			instituteSince:         1957,
		},
		courses: []course{
			{
				name: "Probabilitas",
			},
			{
				name: "Kalkulus",
			},
			{
				name: "Sistem Digital",
			},
		},
	}

	mahasiswa4 := student{
		name: "John Tor",
		npm:  "22222",
		ins: institute{
			instituteName:          "Universitas Dipta",
			instituteAddress:       "Bekasi",
			instituteAccreditation: "A",
			instituteSince:         2022,
		},
		courses: []course{
			{
				name: "Sistem Jaringan",
			},
			{
				name: "Dasar Pemgoraman",
			},
			{
				name: "Sistem Waktu Kontinyu",
			},
		},
	}

	tambahMahasiswa(&mahasiswa, mahasiswa3, mahasiswa4)

	state = 0
	for {
		if state == 0 {
			fmt.Println(strings.Repeat("-", 50))
			fmt.Println("- Selamat datang di Program Basis Data Mahasiswa -")
			fmt.Println(strings.Repeat("-", 50))

			fmt.Println("Silakan pilih tindakan yang ingin dilakukan:")
			fmt.Println("1. Lihat")
			fmt.Println("2. Tambah")
			fmt.Println("3. Hapus")
			fmt.Println("4. Keluar")
			fmt.Print("-> ")
			fmt.Scanln(&inputan)
			if inputan == 1 {
				state = 1
			}
			if inputan == 2 {
				state = 2
			}
			if inputan == 3 {
				state = 3
			}
			if inputan == 4 {
				break
			}
			inputan = 0

		} else if state == 1 {
			for ke, tampil := range mahasiswa {
				fmt.Println("Mahasiswa ke -", ke+1)
				fmt.Println("NPM\t\t\t:", tampil.npm)
				fmt.Println("Nama\t\t\t:", tampil.name)
				fmt.Printf("Perguruan Tinggi\t: %s di %s dengan akreditasi %s yang berdiri sejak %d\n", tampil.ins.instituteName, tampil.ins.instituteAddress, tampil.ins.instituteAccreditation, tampil.ins.instituteSince)
				fmt.Print("Mata Kuliah\t\t: ")
				for index, matkul := range tampil.courses {
					fmt.Print(matkul.name)
					if index < len(tampil.courses)-1 {
						fmt.Print(", ")
					} else {
						fmt.Print("\n\n")
					}
				}
			}
			validInput := bufio.NewScanner(os.Stdin)
			fmt.Print("ketik y untuk kembali ke menu utama -> ")
			validInput.Scan()

			if validInput.Text() == "y" {
				state = 0
			}
		} else if state == 2 {
			if len(mahasiswa) == 5 {
				fmt.Println("Program ini hanya dapat menampung maksimal 5 mahasiswa.")
				fmt.Println("Silakan hapus terlebih dahulu data mahasiswa agar dapat menambahkan data baru")
				state = 0
			}
			for {
				fmt.Println("Silakan masukkan data-data berikut:")
				mhsNama := bufio.NewScanner(os.Stdin)
				fmt.Print("Nama : ")
				mhsNama.Scan()
				mhsNpm := bufio.NewScanner(os.Stdin)
				fmt.Print("NPM : ")
				mhsNpm.Scan()
				mhsPt := bufio.NewScanner(os.Stdin)
				fmt.Print("Perguruan Tinggi : ")
				mhsPt.Scan()
				mhsPtAddress := bufio.NewScanner(os.Stdin)
				fmt.Print("Alamat Perguruan Tinggi : ")
				mhsPtAddress.Scan()
				mhsPtAccred := bufio.NewScanner(os.Stdin)
				fmt.Print("Akreditasi Perguruan Tinggi : ")
				mhsPtAccred.Scan()
				mhsPtSince := bufio.NewScanner(os.Stdin)
				fmt.Print("Tahun Berdiri Perguruan Tinggi : ")
				mhsPtSince.Scan()
				mhsPtSinceTemp, _ := strconv.Atoi(mhsPtSince.Text())
				mhsCourse := bufio.NewScanner(os.Stdin)
				fmt.Print("Mata Kuliah (Jika lebih dari 1 pisahkan dengan tanda koma(,)) : ")
				mhsCourse.Scan()

				// prosedur validasi
				fmt.Println("Berikut data yang anda inputkan:")
				fmt.Println("Nama :", mhsNama.Text())
				fmt.Println("NPM :", mhsNpm.Text())
				fmt.Println("Perguruan Tinggi :", mhsPt.Text())
				fmt.Println("Alamat Perguruan Tinggi :", mhsPtAddress.Text())
				fmt.Println("Akreditasi Perguruan Tinggi :", mhsPtAccred.Text())
				fmt.Println("Tahun Berdiri Perguruan Tinggi :", mhsPtSince.Text())
				fmt.Println("Mata Kuliah :", mhsCourse.Text())

				validInput := bufio.NewScanner(os.Stdin)
				fmt.Print("Apakah data yang anda inputkan sudah benar?\n(masukkan y jika benar atau n jika salah)\n-> ")
				validInput.Scan()

				if validInput.Text() == "y" {
					courseTemp := []course{}
					for _, courseNya := range strings.Split(mhsCourse.Text(), ",") {
						bantuan := course{courseNya}
						courseTemp = append(courseTemp, bantuan)
					}

					// prosedur penambahan mahasiswa
					mhsTemp := student{
						name: mhsNama.Text(),
						npm:  mhsNpm.Text(),
						ins: institute{
							instituteName:          mhsPt.Text(),
							instituteAddress:       mhsPtAddress.Text(),
							instituteAccreditation: mhsPtAccred.Text(),
							instituteSince:         mhsPtSinceTemp,
						},
						courses: courseTemp,
					}
					tambahMahasiswa(&mahasiswa, mhsTemp)
					fmt.Println("Data Berhasil ditambahkan\nkembali ke menu utama dalam\n1")
					time.Sleep(1 * time.Second)
					fmt.Println(2)
					time.Sleep(1 * time.Second)
					fmt.Println(3)
					state = 0
					break
				}
				fmt.Println("Baik silakan ulangi.")
			}
		} else if state == 3 {
			hapusIndexTemp := 0
			for {
				for ke, tampil := range mahasiswa {
					fmt.Printf("%d. %s dengan NPM  %s\n", ke+1, tampil.name, tampil.npm)
				}
				fmt.Println("\n\nMahasiswa nomor berapa yang ingin anda hapus")
				fmt.Println("(jika kosong mahasiswa dengan nomor terakhir yang akan terhapus)")
				fmt.Print("-> ")
				hapusIndex := bufio.NewScanner(os.Stdin)
				fmt.Print("ketik y untuk kembali ke menu utama -> ")
				hapusIndex.Scan()
				if hapusIndex.Text() == "" {
					hapusIndexTemp = len(mahasiswa)
				} else {
					hapusIndexTemp, _ = strconv.Atoi(hapusIndex.Text())
				}

				if hapusIndexTemp > 0 && hapusIndexTemp <= len(mahasiswa) {
					mhsTemp := hapusMahasiswa(mahasiswa, hapusIndexTemp-1)
					mahasiswa = mhsTemp
					break
				}
				fmt.Printf("Tidak ada data mahasiswa dengan nomor %d\n\n", hapusIndexTemp)
			}
			fmt.Printf("Mahasiswa ke - %d berhasil dihapus\nkembali ke menu utama dalam\n1\n", hapusIndexTemp)
			time.Sleep(1 * time.Second)
			fmt.Println(2)
			time.Sleep(1 * time.Second)
			fmt.Println(3)
			state = 0
		}
	}
}

func tambahMahasiswa(sliceNya *[]student, mahaNya ...student) {
	*sliceNya = append(*sliceNya, mahaNya...)
}

func hapusMahasiswa(sliceNya []student, s int) []student {
	return append(sliceNya[:s], sliceNya[s+1:]...)
}
