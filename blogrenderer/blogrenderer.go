package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"my-super-project/blogposts"
)

var (
	//go:embed templates/*
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (renderer *PostRenderer) Render(w io.Writer, p blogposts.BlogPost) error {
	if err := renderer.templ.Execute(w, p); err != nil {
		return err
	}
	return nil
}
