package main

import "fmt"

const nmax = 2024


/* Tipe bentukan struktur item dengan atribut:
	No Produk (integer), Price (integer), Quantity (integer), 
	Name (string), dan Transaction (integer) */
type item struct {
	noProduct, Price, Quantity int
	Name                       string
	Transaction                int
}
/* 	
	Tipe bentukan struktur transaction dengan atribut:
	Sold (integer) dan Price (integer) 
*/
type transaction struct {
	Sold  int
	Price int
}

/* 	
	Tipe alias arrItem untuk array dari item dengan ukuran nmax 
	dan tipe alias arrTrans untuk array dari Transaction dengan ukuran nmax
*/
type (
	arrItem  [nmax]item
	arrTrans [nmax]transaction
)

var (
	A arrItem
	B arrTrans
	n int
)

func main() {
	var pilih int

	intro()

	for {
		menuUtama(&pilih)
		if pilih == 1 {
			menuPenjualan(&A, &n)
		} else if pilih == 2 {
			menuTransaksi(&A, &B, &n)
		} else if pilih == 3 {
			bye()
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func intro() {
	clearScreen()
	fmt.Println("Selamat Datang di Aplikasi Kasir Minimart")
}

func bye() {
	clearScreen()
	fmt.Println("Terima kasih dan Sampai Jumpa!")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func menuUtama(p *int) {
	fmt.Println("-----------------------")
	fmt.Println("      M E N U          ")
	fmt.Println("-----------------------")
	fmt.Println("1. Penjualan")
	fmt.Println("2. Transaksi")
	fmt.Println("3. Keluar")
	fmt.Println("-----------------------")
	fmt.Print("Pilih: ")
	fmt.Scan(p)
}

func menuPenjualan(items *arrItem, n *int) {
	var pilih int

	for {
		clearScreen()
		fmt.Println("-----------------------")
		fmt.Println("     MENU PENJUALAN     ")
		fmt.Println("-----------------------")
		fmt.Println("1. Daftar Barang")
		fmt.Println("2. Daftar Barang Berdasarkan Stok Terbanyak")
		fmt.Println("3. Daftar Barang Berdasarkan Harga Terendah")
		fmt.Println("4. Tambah Barang")
		fmt.Println("5. Ubah Barang")
		fmt.Println("6. Hapus Barang")
		fmt.Println("7. Kembali")
		fmt.Println("-----------------------")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			listItem(*items, *n)
		} else if pilih == 2 {
			listItemByStock(*items, *n)
		} else if pilih == 3 {
			listItemByPrice(*items, *n)
		} else if pilih == 4 {
			addItem(items, n)
		} else if pilih == 5 {
			updateItem(items, *n)
		} else if pilih == 6 {
			deleteItem(items, n)
		} else if pilih == 7 {
			clearScreen()
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func findSeq(items arrItem, n int, noProduct int) int {
	for i := 0; i < n; i++ {
		if items[i].noProduct == noProduct {
			return i
		}
	}
	return -1
}

func selectionSortByStock(A *arrItem, n int) {
	for pass := 0; pass < n-1; pass++ {
		maxIdx := pass
		for i := pass + 1; i < n; i++ {
			if A[i].Quantity > A[maxIdx].Quantity || (A[i].Quantity == A[maxIdx].Quantity && A[i].noProduct < A[maxIdx].noProduct) {
				maxIdx = i
			}
		}
		A[pass], A[maxIdx] = A[maxIdx], A[pass]
	}
}

func insertionSortByPrice(A *arrItem, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && (A[j].Price > key.Price || (A[j].Price == key.Price && A[j].noProduct > key.noProduct)) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}

func listItem(A arrItem, n int) {
	clearScreen()

	fmt.Printf("%2s %15s %12s %10s\n", "No", "Nama Barang", "Harga", "Stok")
	for i := 0; i < n; i++ {
		fmt.Printf("%2d %15s %12d %10d\n", A[i].noProduct, A[i].Name, A[i].Price, A[i].Quantity)
	}
	fmt.Println()
}

func listItemByStock(A arrItem, n int) {
	clearScreen()
	selectionSortByStock(&A, n)

	fmt.Printf("%2s %15s %12s %10s\n", "No", "Nama Barang", "Harga", "Stok")
	for i := 0; i < n; i++ {
		fmt.Printf("%2d %15s %12d %10d\n", A[i].noProduct, A[i].Name, A[i].Price, A[i].Quantity)
	}
	fmt.Println()
}

func listItemByPrice(A arrItem, n int) {
	clearScreen()
	insertionSortByPrice(&A, n)

	fmt.Printf("%2s %15s %12s %10s\n", "No", "Nama Barang", "Harga", "Stok")
	for i := 0; i < n; i++ {
		fmt.Printf("%2d %15s %12d %10d\n", A[i].noProduct, A[i].Name, A[i].Price, A[i].Quantity)
	}
	fmt.Println()
}


func addItem(A *arrItem, n *int) {
	if *n >= nmax {
		fmt.Println("Kapasitas penyimpanan penuh!")
		return
	}

	fmt.Print("Masukkan No Produk: ")
	fmt.Scan(&A[*n].noProduct)
	fmt.Print("Masukkan Nama Barang: ")
	fmt.Scan(&A[*n].Name)
	fmt.Print("Masukkan Harga: ")
	fmt.Scan(&A[*n].Price)
	fmt.Print("Masukkan Jumlah: ")
	fmt.Scan(&A[*n].Quantity)
	*n++
	fmt.Println("Barang berhasil ditambahkan!")
}

func updateItem(A *arrItem, n int) {
	var noProduct int
	fmt.Print("Masukkan No Produk yang ingin diubah: ")
	fmt.Scan(&noProduct)
	index := findSeq(*A, n, noProduct)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Nama Barang: ")
	fmt.Scan(&A[index].Name)
	fmt.Print("Masukkan Harga: ")
	fmt.Scan(&A[index].Price)
	fmt.Print("Masukkan Jumlah: ")
	fmt.Scan(&A[index].Quantity)
	fmt.Println("Barang berhasil diubah!")
}

func deleteItem(A *arrItem, n *int) {
	var noProduct int
	fmt.Print("Masukkan No Produk yang ingin dihapus: ")
	fmt.Scan(&noProduct)
	index := findSeq(*A, *n, noProduct)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan!")
		return
	}
	for i := index; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	(*A)[*n-1] = item{}
	(*n)--
	fmt.Println("Barang berhasil dihapus!")
}

func menuTransaksi(A *arrItem, B *arrTrans, n *int) {
	var pilih int

	for {
		clearScreen()
		fmt.Println("-----------------------")
		fmt.Println("     MENU TRANSAKSI    ")
		fmt.Println("-----------------------")
		fmt.Println("1. Catat Transaksi")
		fmt.Println("2. Laporan Harian")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			CatatTransaksi(A, B, *n)
		} else if pilih == 2 {
			omzetHarian(*A, *B, *n)
		} else if pilih == 3 {
			clearScreen()
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func CatatTransaksi(A *arrItem, B *arrTrans, n int) {
	var noProduct, Sold int
	clearScreen()
	fmt.Print("Masukkan No Produk: ")
	fmt.Scan(&noProduct)
	index := findSeq(*A, n, noProduct)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Jumlah Terjual: ")
	fmt.Scan(&Sold)

	if Sold > (*A)[index].Quantity {
		fmt.Println("Jumlah terjual melebihi stok!")
	} else {
		B[index].Sold += Sold
		B[index].Price = (*A)[index].Price
		(*A)[index].Quantity -= Sold
		fmt.Println("Transaksi berhasil dicatat!")
	}
}

func omzetHarian(A arrItem, B arrTrans, n int) {
	var totalOmzet float64
	clearScreen()

	fmt.Printf("%2s %15s %12s %10s\n", "No", "Nama Barang", "Jumlah", "Total")
	for i := 0; i < n; i++ {
		if B[i].Sold > 0 {
			totalHarga := float64(B[i].Sold) * float64(B[i].Price)
			fmt.Printf("%2d %15s %12d %10.2f\n", A[i].noProduct, A[i].Name, B[i].Sold, totalHarga)
			totalOmzet += totalHarga
		}
	}
	if totalOmzet == 0 {
		fmt.Println("Belum ada transaksi yang dicatat.")
	} else {
		fmt.Printf("Total Omzet hari ini: Rp %.2f\n", totalOmzet)
	}
}
