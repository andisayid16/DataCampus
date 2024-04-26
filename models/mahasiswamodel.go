package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andisayid16/DataCampus/config"
	"github.com/andisayid16/DataCampus/entities"
)

type MahasiswaModel struct {
	conn *sql.DB
}

func NewMahasiswaModel() *MahasiswaModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &MahasiswaModel{
		conn: conn,
	}
}

func (p *MahasiswaModel) FindAll() ([]entities.Mahasiswa, error) {
	rows, err := p.conn.Query("select * from mahasiswa")
	if err != nil {
		return []entities.Mahasiswa{}, err
	}
	defer rows.Close()

	var dataMahasiswa []entities.Mahasiswa
	for rows.Next() {
		var mahasiswa entities.Mahasiswa
		rows.Scan(&mahasiswa.Id,
			&mahasiswa.NamaLengkap,
			&mahasiswa.NIM,
			&mahasiswa.JenisKelamin,
			&mahasiswa.TempatLahir,
			&mahasiswa.TanggalLahir,
			&mahasiswa.Alamat,
			&mahasiswa.NoHp,
			&mahasiswa.Fakultas,
		)

		if mahasiswa.JenisKelamin == "1" {
			mahasiswa.JenisKelamin = "Laki-laki"
		} else {
			mahasiswa.JenisKelamin = "Perempuan"
		}

		// yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", mahasiswa.TanggalLahir)
		mahasiswa.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataMahasiswa = append(dataMahasiswa, mahasiswa)
	}

	return dataMahasiswa, nil
}

func (p *MahasiswaModel) Create(mahasiswa entities.Mahasiswa) bool {

	result, err := p.conn.Exec("insert into mahasiswa (nama_lengkap, nim, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp, nama_fakultas) values(?,?,?,?,?,?,?,?)",
		mahasiswa.NamaLengkap, mahasiswa.NIM, mahasiswa.JenisKelamin, mahasiswa.TempatLahir, mahasiswa.TanggalLahir, mahasiswa.Alamat, mahasiswa.NoHp, mahasiswa.Fakultas)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *MahasiswaModel) Find(id int64, mahasiswa *entities.Mahasiswa) error {

	return p.conn.QueryRow("select * from mahasiswa where id = ?", id).Scan(
		&mahasiswa.Id,
		&mahasiswa.NamaLengkap,
		&mahasiswa.NIM,
		&mahasiswa.JenisKelamin,
		&mahasiswa.TempatLahir,
		&mahasiswa.TanggalLahir,
		&mahasiswa.Alamat,
		&mahasiswa.NoHp,
		&mahasiswa.Fakultas,
	)
}

func (p *MahasiswaModel) Update(mahasiswa entities.Mahasiswa) error {

	_, err := p.conn.Exec(
		"update mahasiswa set nama_lengkap = ?, nim = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ?, nama_fakultas = ? where id = ?",
		mahasiswa.NamaLengkap, mahasiswa.NIM, mahasiswa.JenisKelamin, mahasiswa.TempatLahir, mahasiswa.TanggalLahir, mahasiswa.Alamat, mahasiswa.NoHp, mahasiswa.Fakultas, mahasiswa.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *MahasiswaModel) Delete(id int64) {
	p.conn.Exec("delete from mahasiswa where id = ?", id)
}
