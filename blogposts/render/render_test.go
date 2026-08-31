package render

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/minhnghia2k3/_learn_go_with_test/blogposts"
)

func TestRender(t *testing.T) {
	post := blogposts.Post{
		Title: "hello world",
		Body: `## Body
- This
- Body
### End
**The end**`,
		Description: "this is a description",
		Tags:        []string{"tdd", "go"},
	}

	rs, err := NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should convert single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := rs.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("should render index of posts", func(t *testing.T) {
		posts := []blogposts.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		buf := bytes.Buffer{}

		err = rs.RenderIndex(&buf, posts)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("should render body of a post", func(t *testing.T) {
		buf := bytes.Buffer{}

		RenderBody(&buf, post)

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	post := blogposts.Post{
		Title:       "hello world",
		Body:        "this is a body",
		Description: "this is a description",
		Tags:        []string{"tdd", "go"},
	}

	rs, err := NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		rs.Render(io.Discard, post)
	}
}
