package blogrenderer

import (
	"fmt"
	"io"
	"my-super-project/blogposts"
)

func Render(w io.Writer, p blogposts.BlogPost) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
