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

		assertError(t, err, blogposts.WrongBlogPostFileFormatError)
	})
	t.Run("test specific post after extracting", func(t *testing.T) {
		const (
			firstPost = `Title: Post 1
Description: My test description 1
Tags: myTag1`
			secondPost = `Title: Post 2
Description: My test description 2
Tags: myTag2, Personal`
		)

		want := []blogposts.BlogPost{
			{
				Title:       "Post 1",
				Description: "My test description 1",
				Tags:        []string{"myTag1"},
			},
			{
				Title:       "Post 2",
				Description: "My test description 2",
				Tags:        []string{"myTag2", "Personal"},
			},
		}
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstPost)},
			"hello-world2.md": {Data: []byte(secondPost)},
		}
		posts, err := blogposts.NewPostsFromFs(fs)
		got := posts

		assertNoError(t, err)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, expected %d posts", len(posts), len(fs))
		}

	})

}

func assertError(t *testing.T, gotError error, wantError error) {
	if gotError == nil {
		t.Error("Should give an error")
	}
	if gotError != wantError {
		t.Errorf("Expected error %v, but got %v", wantError, gotError)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
