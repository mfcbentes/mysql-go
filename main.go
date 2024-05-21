package main

import (
	"fmt"
	"log"

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

}
