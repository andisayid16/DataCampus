package dosencontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/andisayid16/DataCampus/entities"
	"github.com/andisayid16/DataCampus/libraries"
	"github.com/andisayid16/DataCampus/models"
)

var validation = libraries.NewValidation()
var dosenModel = models.NewDosenModel()

func Index(response http.ResponseWriter, request *http.Request) {

	dosen, _ := dosenModel.FindAll()

	data := map[string]interface{}{
		"dosen": dosen,
	}

	temp, err := template.ParseFiles("views/dosen/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/dosen/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var dosen entities.Dosen
		dosen.NamaLengkap = request.Form.Get("nama_lengkap")
		dosen.NIP = request.Form.Get("nip")
		dosen.JenisKelamin = request.Form.Get("jenis_kelamin")
		dosen.TempatLahir = request.Form.Get("tempat_lahir")
		dosen.TanggalLahir = request.Form.Get("tanggal_lahir")
		dosen.Alamat = request.Form.Get("alamat")
		dosen.NoHp = request.Form.Get("no_hp")
		dosen.Fakultas = request.Form.Get("fakultas")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(dosen)

		if vErrors != nil {
			data["dosen"] = dosen
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data dosen berhasil disimpan"
			dosenModel.Create(dosen)
		}

		temp, _ := template.ParseFiles("views/dosen/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var dosen entities.Dosen
		dosenModel.Find(id, &dosen)

		data := map[string]interface{}{
			"dosen": dosen,
		}

		temp, err := template.ParseFiles("views/dosen/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var dosen entities.Dosen
		dosen.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		dosen.NamaLengkap = request.Form.Get("nama_lengkap")
		dosen.NIP = request.Form.Get("nip")
		dosen.JenisKelamin = request.Form.Get("jenis_kelamin")
		dosen.TempatLahir = request.Form.Get("tempat_lahir")
		dosen.TanggalLahir = request.Form.Get("tanggal_lahir")
		dosen.Alamat = request.Form.Get("alamat")
		dosen.NoHp = request.Form.Get("no_hp")
		dosen.Fakultas = request.Form.Get("fakultas")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(dosen)

		if vErrors != nil {
			data["dosen"] = dosen
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data dosen berhasil diperbarui"
			dosenModel.Update(dosen)
		}

		temp, _ := template.ParseFiles("views/dosen/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	dosenModel.Delete(id)

	http.Redirect(response, request, "/dosen", http.StatusSeeOther)
}
