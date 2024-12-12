package main

import (
    "fmt"
)

// Fungsi untuk melakukan insertion sort
func insertionSort(arr []int) {
    n := len(arr)
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

// Fungsi untuk memeriksa jarak tetap antar elemen
func checkEqualSpacing(arr []int) (bool, int) {
    if len(arr) < 2 {
        return true, 0
    }
    spacing := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        if arr[i]-arr[i-1] != spacing {
            return false, 0
        }
    }
    return true, spacing
}

func main() {
    var data []int
    var input int

    // Baca data sampai menemukan bilangan negatif
    for {
        fmt.Scan(&input)
        if input < 0 {
            break
        } else if input >= 0 {
            data = append(data, input)
        }
    }

    // Urutkan data menggunakan insertion sort
    insertionSort(data)

    // Cetak array yang sudah diurutkan
    fmt.Println("Array yang diurutkan:", data)

    // Periksa apakah data berjarak tetap
    isEqualSpacing, spacing := checkEqualSpacing(data)
    if isEqualSpacing {
        fmt.Printf("Data berjarak %d\n", spacing)
    } else {
        fmt.Println("Data berjarak tidak tetap")
    }
}
