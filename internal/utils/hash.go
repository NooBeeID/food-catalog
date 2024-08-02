package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// berfungsi untuk melakukan hash terhadap sebuah plain text
// function ini nantinya akan digunakan untuk hash password
func Hash(plain string) (hash string, err error) {
	// proses hashing menggunakan library bcrypt
	// salt yang digunakan adalah 10, yaitu DefaultCost
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	// hasil dari GenerateFromPassword adalah sebuah slice of byte
	// jadi perlu kita ubah ke sebuah string
	hash = string(hashByte)
	return
}

// function ini berfungsi untuk melakukan verifikasi terhadap hash yang diberi
// dengan plain text.
func Verify(hash string, plain string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return
}
