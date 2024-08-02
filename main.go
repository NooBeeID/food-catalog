package main

import (
	"project-catalog/apps"
	"project-catalog/external/database"
	"project-catalog/internal/config"
)

func main() {
	// load config
	err := config.LoadEnvConfig(".env")
	if err != nil {
		panic(err)
	}

	// silahkan sesuaikan dengan yang kamu punya
	// normalnya untuk host = localhost
	dbHost := config.GetenvString(config.CFG_DB_HOST, "postgres")

	// port biasanya 5432, tergantung dengan yang kamu buat saat proses installasi
	dbPort := config.GetenvString(config.CFG_DB_PORT, "5432")

	// username dari database kamu
	dbUser := config.GetenvString(config.CFG_DB_USER, "postgres")

	// password dari database kamu
	dbPass := config.GetenvString(config.CFG_DB_PASS, "")

	// nama database
	dbName := config.GetenvString(config.CFG_DB_NAME, "")

	// apps port
	db, err := database.ConnectPostgsres(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		panic(err)
	}

	// setup port aplikasi yang akan kita buka
	appPort := ":4444"

	// run apps
	apps.Run(appPort, db)
}
