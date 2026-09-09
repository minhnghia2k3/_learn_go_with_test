package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/minhnghia2k3/_learn_go_with_test/http_server"
)

const dbFileName = "game.db.json"

func main() {
	store, cleanup, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
