package main

import (
	"net/http"
)

func main() {
	// Menentukan direktori yang akan disajikan
	fs := http.FileServer(http.Dir("./view"))

	// Menggunakan FileServer untuk menyajikan file dari root ("/")
	http.Handle("/", fs)

	// Menjalankan server HTTP di port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
