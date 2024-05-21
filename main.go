package main

import (
	"fmt"
	"log"

	"github.com/mfcbentes/mysql-go/model"
	"github.com/mfcbentes/mysql-go/model/repository"
)

func main() {

	albums, err := repository.GetAlbums()
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albums {
		fmt.Println(album)
	}

	alb, err := repository.GetAlbumByID(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %d, Title: %s, Artist: %s, Price: %0.2f\n", alb.ID, alb.Title, alb.Artist, alb.Price)

	albID, err := repository.AddAlbum(model.Album{
		Title:  "Toxicity",
		Artist: "System Of A Down",
		Price:  65.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

}
