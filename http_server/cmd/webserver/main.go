package main

import (
	"log"
	"net/http"

	poker "github.com/minhnghia2k3/_learn_go_with_test/http_server"
)

const gameDatabase = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(gameDatabase)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	log.Println("listing on port :8000")
	log.Fatal(http.ListenAndServe(":8000", server))
}
