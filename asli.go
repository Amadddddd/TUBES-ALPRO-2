package main

import "fmt"

const nmax int = 100

type bahanMakanan struct {
	namaBahan  string
	jumlahStok int
	tgl        int
}
type arrBahanMakanan [nmax]bahanMakanan

// var tempBahan arrBahanMakanan

func main() {
	var pilihanMenu string
	var index int = 1
	var dataBahan arrBahanMakanan

	menuUtama(&pilihanMenu)

	// for pilihanMenu != "1" && pilihanMenu != "2" && pilihanMenu != "3" && pilihanMenu != "4" && pilihanMenu != "5" {
	// 	fmt.Print("Pilihan: ")
	// 	fmt.Scan(&pilihanMenu)
	// }

	// perulangan di sini berfungsi sebagai penentu menu yang akan dimasuki beredasarkan nilai pilihanMenu
	for {
		switch pilihanMenu {
		case "1":
			menuTambahan(&dataBahan, &index)
		case "2":
			menuUbah(&dataBahan, index)
		case "3":
			menuDelete(&dataBahan, &index)
		case "4":
			daftarBahan(&dataBahan, index)
		case "5":
			fmt.Println("Terima kasih sudah menggunakan aplikasi kami :)")
			return
		}
		menuUtama(&pilihanMenu)

	}

	// fmt.Scan(&pilihanMenu)
	// for {
	// 	if pilihanMenu == "1" {
	// 		menuTambahan(&dataBahan, &index)

	// 	} else if pilihanMenu == "2" {
	// 		menuUbah()

	// 	} else if pilihanMenu == "3" {
	// 		menuDelete()

	// 	} else if pilihanMenu == "4" {
	// 		daftarBahan(&dataBahan, index)

	// 	} else if pilihanMenu == "5" {
	// 		return
	// 	} else {
	// 		// fmt.Print("Pilihan: ")

	// 	}
	// 	menuUtama(&pilihanMenu)

	// switch pilihanMenu {
	// case 1:
	// 	menuTambahan(&dataBahan, &index)
	// case 2:
	// 	menuUbah()
	// case 3:
	// 	menuDelete()
	// case 4:
	// 	daftarBahan(&dataBahan, index)
	// case 5:
	// 	fmt.Println("Terima kasih sudah menggunakan aplikasi kami :)")
	// 	return
	// default:
	// 	fmt.Print("Pilihan: ")
	// 	fmt.Scan(&pilihanMenu)

	//}

	// }
}

func menuUtama(pilihanMenu *string) {
	/*
		IS. diberikan sebuah string yang berfungsi sebagai penentu masuknya kedalam menu bagian mana berdasarkan angka uang nantinya di beri masukan oleh user, selain itu prosedur ini juga mengeluarkan tampilan menunya
		FS. string yang dimasukan oleh user tadi di kembalikan ke main fungsi
	*/

	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Menu Utama--------")
	fmt.Println("1. Tambahan ")
	fmt.Println("2. Ubah ")
	fmt.Println("3. Delete ")
	fmt.Println("4. Daftar Bahan ")
	fmt.Println("5. Keluar ")
	fmt.Println("----------------------------")
	fmt.Print("Pilihan: ")
	fmt.Scan(&*pilihanMenu)

	// akan melakukan perulangan selama nilai dari pilihan menu bukan 1,2,3,4,5 . angka tersebut merupakan string, lalu user akan terus melakukan inputan sampai nilai yang di maksudkan ada
	for *pilihanMenu != "1" && *pilihanMenu != "2" && *pilihanMenu != "3" && *pilihanMenu != "4" && *pilihanMenu != "5" {
		fmt.Print("Pilihan: ")
		fmt.Scan(&*pilihanMenu)
	}

	// for *pilihanMenu != "1" && *pilihanMenu != "2" && *pilihanMenu != "3" && *pilihanMenu != "4" && *pilihanMenu != "5" {
	// 	fmt.Print("Pilihan: ")
	// 	fmt.Scan(&*pilihanMenu)

	// }
	// switch *pilihanMenu {
	// case "1":
	// 	menuTambahan(dataBahan, i)
	// case "2":
	// 	menuUbah(*dataBahan, *i)
	// case "3":
	// 	menuDelete()
	// case "4":
	// 	daftarBahan(*dataBahan, *i)
	// case "5":
	// 	fmt.Println("Terima kasih sudah menggunakan aplikasi kami :)")
	// 	return
	// }
	// menuUtama(pilihanMenu, i, dataBahan)

}
func menuTambahan(dataBahan *arrBahanMakanan, n *int) {
	/*
		IS. diberikan sebuah data array yang berfungsi sebagai menyimpan data yang akan di tambahkan pleh user, lalu ada n yang fungsi nya banyaknya koleksi data dari array
		FS. dataBahan akan di kembalikan setelah user menginputkan data yang akan di tambahkan, lalu n dikembalikan untuk diketahui banyaknya data yang telah di tambahkan
	*/
	// variabel piilihan berguna untuk menentukan data yang ditambahkan akan disimpan atau tidak berdasarkan inputana angka 1 atau 2
	var pilihan int
	// variabel simpanSementara akan digunakan sebagai array sementara yang nantinya akan di simpan ke dalam koleksi data utama atau tidak
	var simpanSementara arrBahanMakanan

	fmt.Println("------------------------------------------------")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Tambahan--------")
	fmt.Print("Nama Bahan: ")
	fmt.Scan(&simpanSementara[*n-1].namaBahan)

	fmt.Print("Jumlah Stok(masukan angka tidak boleh selain angka): ")
	fmt.Scan(&simpanSementara[*n-1].jumlahStok)

	fmt.Print("Masukan Berapa lama sampai Kadaluarsa(masukan angka tidak boleh selain angka): ")
	fmt.Scan(&simpanSementara[*n-1].tgl)

	fmt.Println("----------------------------")
	fmt.Println("1. Simpan, 2. Batal")
	fmt.Scan(&pilihan)

	// switch case ini berfungsi sebagai penentu data yang telah di inputkan oleh user akan di simpan atau tidak
	switch pilihan {
	case 1:
		dataBahan[*n-1].namaBahan = simpanSementara[*n-1].namaBahan
		dataBahan[*n-1].jumlahStok = simpanSementara[*n-1].jumlahStok
		dataBahan[*n-1].tgl = simpanSementara[*n-1].tgl
		*n = *n + 1

	}
	//fmt.Printf("Nama Bahan : %s \nJumlah Stok : %d\nKadaluarsa : %d\n", simpanSementara[0].namaBahan, simpanSementara[0].jumlahStok, simpanSementara[0].tgl)

}

func menuUbah(dataBahan *arrBahanMakanan, n int) {
	/*
		IS. diberikan sebuah array yang memiliki beberapa data yang berfungsi sebagai pilihan yang akan di ubah
		FS. User memilih data yang akan di ubah berdasarkan nama dari suatu data yang ada di dalam array, lalu memilih mau menyimpan perubahan atau tidak,
	*/

	var cari string
	var hasilPencarian int

	fmt.Println("------------------------------------------------")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Ubah--------")

	// menampilkan data dalam array fungsinya agar user lebih mudah melihat data yang ada
	daftarBahan(dataBahan, n)

	fmt.Print("Cari Nama Data yang akan di ubah(sequencial search): ")
	fmt.Scan(&cari)
	hasilPencarian = sequentialSearch(*dataBahan, n, cari)

	// perulangan disini tujuannya jika inputan user tidak sesuai dengan data yang ada maka inputan akan terus berulang
	for hasilPencarian < 0 {

		fmt.Println("Data tidak ditemukan !!")
		fmt.Print("Cari Nama Data yang akan di ubah: ")
		fmt.Scan(&cari)
		hasilPencarian = sequentialSearch(*dataBahan, n, cari)
	}

	fmt.Println("Data ditemukan !!\nData saat ini : ")

	// menampilkan data array yang dicari
	fmt.Printf("Nama Bahan : %s \nJumlah Stok : %d\nKadaluarsa : %d\n", dataBahan[hasilPencarian-1].namaBahan, dataBahan[hasilPencarian-1].jumlahStok, dataBahan[hasilPencarian-1].tgl)
	fmt.Println("----------------------------")

	// memanggil fungsu mengubah untuk mengubah data
	mengubah(&*dataBahan, hasilPencarian)
}

func menuDelete(databahan *arrBahanMakanan, n *int) {
	var cari string
	var pilihan int
	var hasilPencarian int
	fmt.Println("------------------------------------------------")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Delete--------")
	fmt.Print("Masukan Nama Bahan(binary search): ")
	fmt.Scan(&cari)
	hasilPencarian = binSearch(*databahan, *n, cari)

	for hasilPencarian < 0 {

		fmt.Println("Data tidak ditemukan !!")
		fmt.Print("Cari Nama Data yang akan di ubah: ")
		fmt.Scan(&cari)
		hasilPencarian = binSearch(*databahan, *n, cari)
	}
	fmt.Println("Data ditemukan !!\nData saat ini : ")

	// menampilkan data array yang dicari
	fmt.Printf("Nama Bahan : %s \nJumlah Stok : %d\nKadaluarsa : %d\n", databahan[hasilPencarian].namaBahan, databahan[hasilPencarian].jumlahStok, databahan[hasilPencarian].tgl)
	fmt.Println("----------------------------")
	fmt.Println("Apakah Datanya akan di hapus?\n1. Iya, 2. Batal")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		hapusData(&*databahan, &*n, hasilPencarian)
	}

}
func daftarBahan(dataBahan *arrBahanMakanan, n int) {

	var i int
	if n-1 < 1 {
		fmt.Println("Data Belum Ada")
	} else {
		fmt.Println("------------------------------------------------")
		fmt.Println("")
		fmt.Println("             APLIKASI DEMO TUBES")
		fmt.Println("")
		fmt.Println("------------------------------------------------")
		fmt.Println("")
		fmt.Println("----------Daftar--------")
		fmt.Println("Daftar Bahan: ")
		for i = 1; i <= n-1; i++ {
			fmt.Printf("Nama Bahan : %s \nJumlah Stok : %d\nKadaluarsa : %d\n", dataBahan[i-1].namaBahan, dataBahan[i-1].jumlahStok, dataBahan[i-1].tgl)

		}

		fmt.Printf("")
		fmt.Println("----------------------------")

	}

}

func sequentialSearch(T arrBahanMakanan, n int, x string) int {
	var found bool = false
	var i int = 0
	for i < n && !found {
		found = T[i].namaBahan == x
		i++
	}
	if found {
		return i
	} else {
		return 0
	}

}

func mengubah(databahan *arrBahanMakanan, n int) {
	var tempBahan arrBahanMakanan
	var pilihan int

	fmt.Print("Masukan Nama Bahan yang akan di Ubah: ")
	fmt.Scan(&tempBahan[n-1].namaBahan)
	fmt.Print("Masukan Jumlah stok yang akan di Ubah(Hanya boleh masukan angka): ")
	fmt.Scan(&tempBahan[n-1].jumlahStok)
	fmt.Print("Masukan jadi berapa hari lagi tanggal kadaluarsa bahan tersebut(hanya boleh masukan angka): ")
	fmt.Scan(&tempBahan[n-1].tgl)
	fmt.Println("1. Simpan, 2. Batal")
	fmt.Scan(&pilihan)

	// switch case ini berfungsi sebagai penentu data yang telah di inputkan oleh user akan di simpan atau tidak
	switch pilihan {
	case 1:
		databahan[n-1].namaBahan = tempBahan[n-1].namaBahan
		databahan[n-1].jumlahStok = tempBahan[n-1].jumlahStok
		databahan[n-1].tgl = tempBahan[n-1].tgl

	}

}

func binSearch(tab arrBahanMakanan, n int, x string) int {
	var mid, left, right int
	var found int = -1
	left = 0
	right = n - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x > tab[mid].namaBahan {
			right = mid - 1
		} else if x < tab[mid].namaBahan {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func hapusData(databahan *arrBahanMakanan, n *int, x int) {
	var i int
	for i = x; i < *n; i++ {
		*&databahan[x] = *&databahan[x+1]
	}
	*n = *n - 1

}

// func InsertionSort(A *arrBahanMakanan, N int) {
// 	var i, pass int
// 	var temp int
// 	pass = 1
// 	for pass <= N-1 {
// 		i = pass
// 		temp = A[pass]
// 		for i > 0 && temp < A[i-1] {
// 			A[i] = A[i-1]
// 			i--
// 		}
// 		A[i] = temp
// 		pass = pass + 1
// 	}
// }

// func SelectionSort(A *arrBahanMakanan, N int) {
// 	var i, idx, pass int
// 	var temp int

// 	pass = 1
// 	for i < N {
// 		idx = pass - 1
// 		i = pass
// 		for i < N {
// 			if A[i] > A[idx] {
// 				idx = i
// 			}
// 			i++
// 		}
// 		temp = A[pass-1]
// 		A[pass-1] = A[idx]
// 		A[idx] = temp
// 		pass = pass + 1
// 	}
// }
