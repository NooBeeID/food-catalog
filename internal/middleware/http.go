package middleware

import (
	"context"
	"log"
	"net/http"
	"project-catalog/internal/helper"
	"project-catalog/internal/utils"
	"strings"
	"time"
)

func Tracer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get current time
		now := time.Now()

		// setup log for incoming request
		log.Printf("method=%v path=%v type=[INFO] message='incoming request'", r.Method, r.URL.Path)

		// pass to next handler
		h.ServeHTTP(w, r)

		// compare process and get time by ms
		end := time.Since(now).Milliseconds()

		// log for after request with response time
		log.Printf("method=%v path=%v type=[INFO] message='finish request' response_time=%vms", r.Method, r.URL.Path, end)
	})
}

func CheckToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get header Authorization
		// karena peletakan token ada di header ini
		// bentuknya nanti akan seperti ini
		// Authorization : Bearer <token>
		bearer := r.Header.Get("Authorization")

		// jika gaada bearer token, maka return unauthorized
		if bearer == "" {
			resp := helper.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Error:   "no token provided",
			}
			resp.WriteJsonResponse(w)
			return
		}

		// lakukan split string
		// artinya stringnya akan kita potong berdasarkan "Bearer "
		tokSlice := strings.Split(bearer, "Bearer ")

		// expectnya akan ada 2 data
		if len(tokSlice) < 2 {
			resp := helper.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Error:   "invalid token",
			}
			resp.WriteJsonResponse(w)
			return
		}

		// data pada index 0 nya akan kosong
		// dan pada index 1 adalah isi dari tokennya (token string)
		tokString := tokSlice[1]

		// setelah itu, lakukan verifikasi token tersebut
		token, err := utils.VerifyToken(tokString)

		// jika error, return unauthorized
		if err != nil {
			resp := helper.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Error:   err.Error(),
			}
			resp.WriteJsonResponse(w)
			return
		}

		// nah bagiain ini, kita akan mengirim value dari id nya
		// via context.
		ctx := context.WithValue(r.Context(), "AUTH_ID", token.Id)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
