package blogrenderer

import (
	"html/template"
	"io"
	"my-super-project/blogposts"
)

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(w io.Writer, p blogposts.BlogPost) error {
	temp, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}
	if err := temp.Execute(w, p); err != nil {
		return err
	}
	return nil
}
