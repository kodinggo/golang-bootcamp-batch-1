package main

import (
	"errors"
	"fmt"
)

func main() {
	err := cekUsia(2)
	if err != nil {
		fmt.Println(err.Error()) // Output: Anda dibawah umur!
	}
}

func cekUsia(usia int64) error {
	if usia < 17 {
		return errors.New("anda dibawah umur!") // Menginisialisasi pesan error
	}
	return nil
}
