package fakultascontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/andisayid16/DataCampus/entities"
	"github.com/andisayid16/DataCampus/libraries"
	"github.com/andisayid16/DataCampus/models"
)

var validation = libraries.NewValidation()
var fakultasModel = models.NewFakultasModel()

func Index(response http.ResponseWriter, request *http.Request) {

	fakultas, _ := fakultasModel.FindAll()

	data := map[string]interface{}{
		"fakultas": fakultas,
	}

	temp, err := template.ParseFiles("views/fakultas/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/fakultas/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var fakultas entities.Fakultas
		fakultas.Fakultas = request.Form.Get("data_fakultas")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(fakultas)

		if vErrors != nil {
			data["fakultas"] = fakultas
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data fakultas berhasil disimpan"
			fakultasModel.Create(fakultas)
		}

		temp, _ := template.ParseFiles("views/fakultas/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var fakultas entities.Fakultas
		fakultasModel.Find(id, &fakultas)

		data := map[string]interface{}{
			"fakultas": fakultas,
		}

		temp, err := template.ParseFiles("views/fakultas/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var fakultas entities.Fakultas
		fakultas.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		fakultas.Fakultas = request.Form.Get("data_fakultas")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(fakultas)

		if vErrors != nil {
			data["fakultas"] = fakultas
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data fakultas berhasil diperbarui"
			fakultasModel.Update(fakultas)
		}

		temp, _ := template.ParseFiles("views/fakultas/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	fakultasModel.Delete(id)

	http.Redirect(response, request, "/fakultas", http.StatusSeeOther)
}
