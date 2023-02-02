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

func Render(w io.Writer, p blogposts.BlogPost) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}
	if err := templ.Execute(w, p); err != nil {
		return err
	}
	return nil
}
