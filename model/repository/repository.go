package repository

import (
	"github.com/mfcbentes/mysql-go/db"
	"github.com/mfcbentes/mysql-go/model"
)

func GetAlbums() ([]model.Album, error) {
	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var albums []model.Album

	stmt, err := conn.Prepare("SELECT * FROM album")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var alb model.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return albums, nil
}
