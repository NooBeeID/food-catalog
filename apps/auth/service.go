package auth

import (
	"database/sql"
	"fmt"
	"log"
	"project-catalog/internal/utils"
)

type repositoryContract interface {
	create(auth Auth) (err error)
	getByEmail(email string) (auth Auth, err error)
}

type service struct {
	// membutuhkan dependency ke repository
	// yang mana harus sesuai dengan kontrak yang sudah di sepakati (interface)
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

// method untuk register data auth
func (s service) create(auth Auth) (err error) {
	// hash password user sebelum di insert ke db
	auth.Password, err = utils.Hash(auth.Password)
	if err != nil {
		log.Println("error when try to hash password with error", err.Error())
		return
	}

	// insert ke datasource
	err = s.repo.create(auth)
	if err != nil {
		log.Println("error when try to create auth with error", err.Error())
		return
	}
	return
}

// method untuk login
func (s service) login(req Auth) (token string, err error) {
	// check apakah user dengan email tersebut ada atau tidak
	auth, err := s.repo.getByEmail(req.Email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err.Error())
		if err == sql.ErrNoRows {
			err = fmt.Errorf("username or password not found")
			return
		}
		return
	}

	// lakukan verifikasi
	// password yang dari database, dalam bentuk hash
	// jadi perlu kita verify
	err = utils.Verify(auth.Password, req.Password)
	if err != nil {
		log.Println("error when try to verify password with error", err.Error())
		err = fmt.Errorf("username or password not found")
		return
	}

	tokenJWT := utils.NewJWT(auth.Id)
	token, err = tokenJWT.GenerateToken()
	if err != nil {
		log.Println("error when try to GenerateToken with error", err.Error())
		return
	}
	return
}
