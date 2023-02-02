package blogrenderer

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"io"
	"my-super-project/blogposts"
	"testing"
)

func TestRenderer(t *testing.T) {
	var (
		aPost = blogposts.BlogPost{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		renderer, _ := NewRenderer()
		if err := renderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogposts.BlogPost{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	renderer, _ := NewRenderer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = renderer.Render(io.Discard, aPost)
	}
}
