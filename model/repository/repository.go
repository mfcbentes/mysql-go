package repository

import (
	"database/sql"
	"fmt"

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

func GetAlbumByID(id int64) (*model.Album, error) {
	alb := &model.Album{}

	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return alb, nil
}

func AddAlbum(alb model.Album) (int64, error) {
	conn, err := db.GetConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateAlbum(alb model.Album) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT id FROM album WHERE id = ?", alb.ID)
	var id int64
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no album found with id %d", alb.ID)
		}
		return err
	}

	stmt, err := conn.Prepare("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(alb.Title, alb.Artist, alb.Price, alb.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAlbum(id int64) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("DELETE FROM album WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
