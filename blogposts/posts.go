package blogposts

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

func getPost(filesystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := filesystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readline := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readline(titlePrefix)
	description := readline(descriptionPrefix)
	tags := strings.Split(readline(tagsPrefix), ", ")
	body := readBody(scanner)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	// skip the seperator
	scanner.Scan()

	var b strings.Builder
	for scanner.Scan() {
		fmt.Fprintln(&b, scanner.Text())
	}

	return strings.TrimSuffix(b.String(), "\n")
}
