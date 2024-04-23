package main

import (
	"fmt"
	"math/rand"
	// "net/http"
)

func main() {
	// var usia int
	// fmt.Println("Masukan Usia : ")
	// fmt.Scan(&usia)

	// if usia < 18 {
	// 	fmt.Println("Belum cukup umur")
	// } else if usia >= 18 && usia <= 60 {
	// 	fmt.Println("Silahkan masuk")
	// } else {
	// 	fmt.Println("Tidak pada kriteria")
	// }

	// response, err := http.Get("https://milestone0-static-afemz.netlify.app")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// switch response.StatusCode {
	// case 200:
	// 	fmt.Println("OK")
	// case 404:
	// 	fmt.Println("Not Found")
	// case 500:
	// 	fmt.Println("Internal Server Error")
	// default:
	// 	fmt.Printf("Received HTTP Status Code : %d\n", response.StatusCode)
	// }

	NGCQ1()
	// NGCQ2()
}

func NGCQ1() {
	var nama string
	fmt.Printf("Masukan nama : ")
	fmt.Scan(&nama)

	randomNumber := rand.Intn(100) + 1

	if randomNumber > 80 {
		fmt.Printf("Nomor kamu : %v ", randomNumber)
		println()
		fmt.Printf("Selamat, %s sangat beruntung", nama)
	} else if randomNumber > 60 && randomNumber <= 80 {
		fmt.Printf("Nomor kamu : %v", randomNumber)
		println()
		fmt.Printf("Selamat, %s anda beruntung", nama)
	} else if randomNumber > 40 && randomNumber <= 60 {
		fmt.Printf("Nomor kamu : %v", randomNumber)
		println()
		fmt.Printf("Mohon Maaf, %s anda kurang beruntung", nama)
	} else if randomNumber <= 40 {
		fmt.Printf("Nomor kamu : %v", randomNumber)
		println()
		fmt.Printf("Mohon Maaf, %s anda sial", nama)
	}
}

func NGCQ2() {
	var nama string
	fmt.Printf("Masukan nama : ")
	fmt.Scan(&nama)

	var umur int
	fmt.Printf("Masukan umur : ")
	fmt.Scan(&umur)

	if umur <= 0 {
		fmt.Println("Umur tidak bisa 0")
	} else if umur >= 100 {
		fmt.Println("Tidak pada kriteria")
	} else if umur <= 18 {
		fmt.Println("Dilarang masuk, minimal umur 19")
	} else if umur >= 19 {
		fmt.Println("Silahkan masuk")
	} else {
		fmt.Println("Tidak pada kriteria")
	}

}
