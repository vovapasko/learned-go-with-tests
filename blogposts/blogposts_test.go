package blogposts_test

import (
	"github.com/vovapasko/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("test extracting tests from file system", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}
		_, err := blogposts.NewPostsFromFs(fs)

		assertError(t, err)
	})
	t.Run("test specific post after extracting", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}
		posts, err := blogposts.NewPostsFromFs(fs)
		got := posts[0]
		want := blogposts.BlogPost{Title: "Post 1"}

		assertNoError(t, err)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, expected %d posts", len(posts), len(fs))
		}

	})

}

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Error("Should give an error")
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
