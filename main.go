package main

import (
	"fmt"
	"log"

	"github.com/mfcbentes/mysql-go/model"
	"github.com/mfcbentes/mysql-go/model/repository"
)

func main() {

	// Get all albums
	albums, err := repository.GetAlbums()
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albums {
		fmt.Println(album)
	}

	// Get album by ID
	alb, err := repository.GetAlbumByID(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %d, Title: %s, Artist: %s, Price: %0.2f\n", alb.ID, alb.Title, alb.Artist, alb.Price)

	// Add album
	albID, err := repository.AddAlbum(model.Album{
		Title:  "Complicated",
		Artist: "Avril Lavigne",
		Price:  17.50,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	// Update album
	albUpdate := model.Album{
		ID:     1,
		Title:  "Fallen",
		Artist: "Evanescence",
		Price:  19.99,
	}

	err = repository.UpdateAlbum(albUpdate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Album updated successfully")

	// Delete album
	err = repository.DeleteAlbum(3)
	if err != nil {
		log.Fatal(err)
	}

}
