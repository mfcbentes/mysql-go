# Projeto de Gerenciamento de Álbuns

Este projeto é um sistema simples de gerenciamento de álbuns. Ele permite que você adicione, atualize e recupere álbuns de um banco de dados.

## Como usar

Primeiro, você precisa configurar o banco de dados. As configurações do banco de dados estão localizadas no arquivo db.go.

Depois de configurar o banco de dados, você pode começar a usar o sistema de gerenciamento de álbuns. Aqui estão alguns exemplos de como você pode fazer isso:

### Adicionar um álbum

```go
alb := model.Album{
Title: "Some Title",
Artist: "Some Artist",
Price: 19.99,
}

id, err := repository.AddAlbum(alb)
if err != nil {
log.Fatal(err)
}

fmt.Printf("Added album with ID %d\n", id)
```

### Atualizar um álbum

```go
alb := model.Album{
ID: 1,
Title: "New Title",
Artist: "New Artist",
Price: 19.99,
}

err := repository.UpdateAlbum(alb)
if err != nil {
log.Fatal(err)
}

fmt.Println("Updated album")
```

### Recuperar um álbum

```go
alb, err := repository.GetAlbumByID(1)
if err != nil {
    log.Fatal(err)
}

fmt.Println(alb)
```

### Buscar todos os álbuns

```go
albs, err := repository.GetAllAlbums()
if err != nil {
    log.Fatal(err)
}

for _, alb := range albs {
    fmt.Println(alb)
}

```

### Excluir um álbum

```go
err := repository.DeleteAlbum(1) // Replace 1 with the ID of the album you want to delete
if err != nil {
  log.Fatal(err)
}

fmt.Println("Deleted album")
```

### Contribuindo

Contribuições são bem-vindas! Por favor, sinta-se à vontade para abrir uma issue ou enviar um pull request.
