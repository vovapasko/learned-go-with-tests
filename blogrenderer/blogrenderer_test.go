package blogrenderer

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
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

		if err := Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})
}
