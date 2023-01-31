package blogposts_test

import (
	"github.com/vovapasko/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts := blogposts.NewPostsFromFs(fs)
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, expected %d posts", len(posts), len(fs))
	}

}
