package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// connect to postgreSQL with several params
//
// @host: localhost or ipaddress or domain of database host
//
// @port: merupakan port database postgres. by default adalah 5432
//
// @user: merupakan username dari database. by default adalah postgres
//
// @pass: merupakan password dari database. by default adalah empty string
//
// @dbname: merupakan nama database.
func ConnectPostgsres(host, port, user, pass, dbname string) (*sql.DB, error) {

	// proses pembuatan koneksi
	dbs, err := getPostgres(host, port, user, pass, dbname)
	if err != nil {
		panic(err)
	}

	// cek apakah database telah berhasil connect atau belum
	err = dbs.Ping()
	if err != nil {
		return nil, err
	}

	migrate(dbs)

	return dbs, nil

}

// get postgres connection based on config
func getPostgres(host, port, user, password, dbname string) (*sql.DB, error) {
	// membuat data source
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbs, err := createConnection(dsn)
	if err != nil {

		return nil, err
	}

	return dbs, nil
}

// create postgres connection by data source name (DSN)
func createConnection(dsn string) (*sql.DB, error) {
	// proses membuka koneksi
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// method ini berfungsi untuk melakukan set minimum koneksi yang dibuat
	// jadi saat program dijaklankan, dia akan membuat 10 koneksi yang pada posisi idle
	db.SetMaxIdleConns(10)

	// method ini berfungsi untuk melakukan set maximum jumlah koneksi yang dibuat
	db.SetMaxOpenConns(25)

	// method ini berfungsi untuk menentukan masa aktif sebuah koneksi saat posisi idle
	// jika melebih batas waktu yang ditentukan, maka koneksi akan dihapus hingga batas SetMaxIdleConns
	db.SetConnMaxIdleTime(5 * time.Minute)

	// method ini berfungsi untuk menentukan lamanya sebuah koneksi itu ada
	// jika lewat dari waktunya, maka koneksi akan dihapus dan di generate ulang
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

// migrations
func migrate(db *sql.DB) {
	log.Println("running migrations ...")
	var dir = "./docs/schema"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			// read file migrations
			f, err := os.ReadFile(dir + "/" + file.Name())
			if err != nil {
				panic(err)
			}

			query := string(f)
			_, err = db.ExecContext(context.Background(), query)
			if err != nil {
				panic(err)
			}
		}
	}
	log.Println("migrations completed !")
}
