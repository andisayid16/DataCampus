package models

import (
	"database/sql"
	"fmt"

	"github.com/andisayid16/DataCampus/config"
	"github.com/andisayid16/DataCampus/entities"
)

type FakultasModel struct {
	conn *sql.DB
}

func NewFakultasModel() *FakultasModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &FakultasModel{
		conn: conn,
	}
}

// func (p *FakultasModel) FindAll() ([]entities.Fakultas, error) {
// 	rows, err := p.conn.Query("select * from fakultas")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var fakultas []entities.Fakultas
// 	for rows.Next() {
// 		var fakultas entities.Fakultas
// 		err := rows.Scan(
// 			&fakultas.Id,
// 			&fakultas.Fakultas,
// 		)
// 	}
// 	return fakultas, nil
// }

func (p *FakultasModel) FindAll() ([]entities.Fakultas, error) {
	rows, err := p.conn.Query("SELECT * FROM fakultas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fakultas []entities.Fakultas
	for rows.Next() {
		var f entities.Fakultas
		err := rows.Scan(&f.Id, &f.Fakultas)
		if err != nil {
			return nil, err
		}
		fakultas = append(fakultas, f)
	}

	return fakultas, nil
}

func (p *FakultasModel) Create(fakultas entities.Fakultas) bool {

	result, err := p.conn.Exec("insert into fakultas (fakultas) values(?)", fakultas.Fakultas)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *FakultasModel) Find(id int64, fakultas *entities.Fakultas) error {

	return p.conn.QueryRow("select * from fakultas where id = ?", id).Scan(
		&fakultas.Id,
		&fakultas.Fakultas,
	)
}

func (p *FakultasModel) Update(fakultas entities.Fakultas) error {

	_, err := p.conn.Exec(
		"update fakultas set fakultas = ? where id = ?", fakultas.Fakultas, fakultas.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *FakultasModel) Delete(id int64) {
	p.conn.Exec("delete from fakultas where id = ?", id)
}
