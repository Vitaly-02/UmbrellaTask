package mw

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("User-Role")

		if strings.EqualFold(role, roleAdmin) {
			log.Println("red button user detected")
		}
		next.ServeHTTP(w, r)
	})
}

func AdminHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("User-Role")

		if strings.EqualFold(role, roleAdmin) {
			log.Println("red button user detected")
			next.ServeHTTP(w, r)
		} else {
			fmt.Println(r.Header)
			w.WriteHeader(http.StatusUnauthorized)
		}

	})
}
