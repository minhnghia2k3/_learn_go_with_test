package render

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"

	"github.com/minhnghia2k3/_learn_go_with_test/blogposts"
)

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	tmpl *template.Template
}

// NewPostRenderer renders a template
// avoids re-rendering template every function calls.
func NewPostRenderer() (*PostRenderer, error) {
	tpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{tmpl: tpl}, nil
}

func (r *PostRenderer) Render(w io.Writer, post blogposts.Post) error {
	return r.tmpl.ExecuteTemplate(w, "blog.gohtml", post)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	return r.tmpl.ExecuteTemplate(w, "index.gohtml", posts)
}

func RenderBody(w io.Writer, post blogposts.Post) error {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)

	html := markdown.ToHTML([]byte(post.Body), p, nil)

	_, err := fmt.Fprint(w, string(html))
	return err
}
