package entities

type Fakultas struct {
	Id       int64
	Fakultas string `validate:"required" label:"Nama Fakultas"`
}
