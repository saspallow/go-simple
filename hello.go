package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/admin", allowRoles("admin")(http.HandlerFunc(adminHandler)))
	http.Handle("/staff", allowRoles("staff")(http.HandlerFunc(staffHandler)))
	http.Handle("/admin-staff", allowRoles("admin", "staff")(http.HandlerFunc(adminStaffHandler)))
	err := http.ListenAndServe(":8900", nil)
	log.Println(err)
}

type middleware func(http.Handler) http.Handler

func allowRole(role string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if reqRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func allowRoles(roles ...string) middleware {
	allow := make(map[string]struct{})
	for _, role := range roles {
		allow[role] = struct{}{}
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\nHello"))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\nHello, Admin."))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\nHello, Staff."))
}

func adminStaffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Admin & Staff."))
}
