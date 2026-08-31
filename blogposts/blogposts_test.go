package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/minhnghia2k3/_learn_go_with_test/blogposts"
)

type MockFileSystem struct{}

func (fs *MockFileSystem) Open(name string) (fs.File, error) {
	return nil, errors.New("i always return error")
}

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`

	secondBody = `Title: Post 2
Description: Description 2
Tags: rust
---
Good
Luck
With Rust
---`
)

func TestNewBlogPost(t *testing.T) {
	t.Run("should create new post from fs", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world.md":   {Data: []byte(firstBody)},
			"hello_world_2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`}

		assertPostEqual(t, got, want)
	})

	t.Run("should return err", func(t *testing.T) {
		mockFs := &MockFileSystem{}

		_, err := blogposts.NewPostFromFS(mockFs)
		if err == nil {
			t.Errorf("expected an error, but got none")
		}
	})
}

func assertPostEqual(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
