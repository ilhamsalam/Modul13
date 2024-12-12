package main

import (
    "fmt"
)

const nMax int = 7919

type Buku struct {
    id       string
    judul    string
    penulis  string
    penerbit string
    eksemplar int
    tahun    int
    rating   int
}

type DaftarBuku [nMax]Buku

var Pustaka DaftarBuku
var nPustaka int

// Procedure untuk mendaftarkan buku
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
    fmt.Scan(n)
    for i := 0; i < *n; i++ {
        fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
    }
	fmt.Println()
}

// Procedure untuk mencetak buku terfavorit
func CetakTerfavorit(pustaka *DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku dalam pustaka.")
        return
    }

    terfavorit := pustaka[0]
    for i := 1; i < n; i++ {
        if pustaka[i].rating > terfavorit.rating {
            terfavorit = pustaka[i]
        }
    }
    fmt.Printf("Buku Terfavorit: %s, %s, %s, %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun)
	fmt.Println()
}

// Procedure untuk mengurutkan buku menggunakan insertion sort
func UrutBuku(pustaka *DaftarBuku, n int) {
    for i := 1; i < n; i++ {
        key := pustaka[i]
        j := i - 1
        for j >= 0 && pustaka[j].rating < key.rating {
            pustaka[j+1] = pustaka[j]
            j = j - 1
        }
        pustaka[j+1] = key
    }
}

// Procedure untuk mencetak 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
    fmt.Println("5 Buku dengan Rating Tertinggi:")
    limit := 5
    if n < 5 {
        limit = n
    }
    for i := 0; i < limit; i++ {
        fmt.Printf("%s\n", pustaka[i].judul)
    }
	fmt.Println()
}

// Procedure untuk mencari buku berdasarkan rating
func CariBuku(pustaka *DaftarBuku, n int, r int) {
    left := 0
    right := n - 1

    for left <= right {
        mid := (left + right) / 2
        if pustaka[mid].rating == r {
            fmt.Printf("Buku dengan Rating %d: %s, %s, %s, %d, %d, %d\n", r, pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
            return
        } else if pustaka[mid].rating < r {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
    DaftarkanBuku(&Pustaka, &nPustaka)
    CetakTerfavorit(&Pustaka, nPustaka)
    UrutBuku(&Pustaka, nPustaka)
    Cetak5Terbaru(&Pustaka, nPustaka)

    var ratingCari int
	fmt.Printf("Masukkan rating buku yang dicari : ")
    fmt.Scan(&ratingCari)
    CariBuku(&Pustaka, nPustaka, ratingCari)
}
