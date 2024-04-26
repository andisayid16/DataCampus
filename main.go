package main

import (
	"net/http"

	authcontroller "github.com/andisayid16/DataCampus/controllers"
	"github.com/andisayid16/DataCampus/controllers/dosencontroller"
	"github.com/andisayid16/DataCampus/controllers/fakultascontroller"
	"github.com/andisayid16/DataCampus/controllers/mahasiswacontroller"
	//"github.com/gobuffalo/packr/v2/jam/store"
)

func main() {

	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/", authcontroller.Login)

	http.HandleFunc("/mahasiswa", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/index", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/add", mahasiswacontroller.Add)
	http.HandleFunc("/mahasiswa/edit", mahasiswacontroller.Edit)
	http.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)

	http.HandleFunc("/dosen", dosencontroller.Index)
	http.HandleFunc("/dosen/index", dosencontroller.Index)
	http.HandleFunc("/dosen/add", dosencontroller.Add)
	http.HandleFunc("/dosen/edit", dosencontroller.Edit)
	http.HandleFunc("/dosen/delete", dosencontroller.Delete)

	http.HandleFunc("/fakultas", fakultascontroller.Index)
	http.HandleFunc("/fakultas/index", fakultascontroller.Index)
	http.HandleFunc("/fakultas/add", fakultascontroller.Add)
	http.HandleFunc("/fakultas/edit", fakultascontroller.Edit)
	http.HandleFunc("/fakultas/delete", fakultascontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
