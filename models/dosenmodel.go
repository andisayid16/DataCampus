package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andisayid16/DataCampus/config"
	"github.com/andisayid16/DataCampus/entities"
)

type DosenModel struct {
	conn *sql.DB
}

func NewDosenModel() *DosenModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &DosenModel{
		conn: conn,
	}
}

func (p *DosenModel) FindAll() ([]entities.Dosen, error) {
	rows, err := p.conn.Query("select * from dosen")
	if err != nil {
		return []entities.Dosen{}, err
	}
	defer rows.Close()

	var dataDosen []entities.Dosen
	for rows.Next() {
		var dosen entities.Dosen
		rows.Scan(&dosen.Id,
			&dosen.NamaLengkap,
			&dosen.NIP,
			&dosen.JenisKelamin,
			&dosen.TempatLahir,
			&dosen.TanggalLahir,
			&dosen.Alamat,
			&dosen.NoHp,
			&dosen.Fakultas)

		if dosen.JenisKelamin == "1" {
			dosen.JenisKelamin = "Laki-laki"
		} else {
			dosen.JenisKelamin = "Perempuan"
		}

		// yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", dosen.TanggalLahir)
		dosen.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataDosen = append(dataDosen, dosen)
	}

	return dataDosen, nil
}

func (p *DosenModel) Create(dosen entities.Dosen) bool {

	result, err := p.conn.Exec("insert into dosen (nama_lengkap, nip, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp, fakultas) values(?,?,?,?,?,?,?,?)",
		dosen.NamaLengkap, dosen.NIP, dosen.JenisKelamin, dosen.TempatLahir, dosen.TanggalLahir, dosen.Alamat, dosen.NoHp, dosen.Fakultas)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *DosenModel) Find(id int64, dosen *entities.Dosen) error {

	return p.conn.QueryRow("select * from dosen where id = ?", id).Scan(
		&dosen.Id,
		&dosen.NamaLengkap,
		&dosen.NIP,
		&dosen.JenisKelamin,
		&dosen.TempatLahir,
		&dosen.TanggalLahir,
		&dosen.Alamat,
		&dosen.NoHp,
		&dosen.Fakultas,
	)
}

func (p *DosenModel) Update(dosen entities.Dosen) error {

	_, err := p.conn.Exec(
		"update dosen set nama_lengkap = ?, nip = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ?, fakultas = ? where id = ?",
		dosen.NamaLengkap, dosen.NIP, dosen.JenisKelamin, dosen.TempatLahir, dosen.TanggalLahir, dosen.Alamat, dosen.NoHp, dosen.Fakultas, dosen.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *DosenModel) Delete(id int64) {
	p.conn.Exec("delete from dosen where id = ?", id)
}
