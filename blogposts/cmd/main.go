package main

import (
	"log"
	"os"

	"github.com/minhnghia2k3/_learn_go_with_test/blogposts"
)

func main() {
	posts, _ := blogposts.NewPostFromFS(os.DirFS("posts"))

	log.Printf("posts: %#v\n", posts)
}
