package blogposts_test

import (
	"my-super-project/blogposts"
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
Tags: myTag1
---
Hello world in one line text 1!`
			secondPost = `Title: Post 2
Description: My test description 2
Tags: myTag2, Personal
---

Hello world in many lines text 2!
Here goes something 

and here!

Lorem ipsum`
			thirdPost = `Title: Post 3
Description: My test description 3
Tags:
---
Hello world in many lines text 3!
Here goes something 

and here!

Lorem ipsum
`
		)

		want := []blogposts.BlogPost{
			{
				Title:       "Post 1",
				Description: "My test description 1",
				Tags:        []string{"myTag1"},
				Body:        "Hello world in one line text 1!",
			},
			{
				Title:       "Post 2",
				Description: "My test description 2",
				Tags:        []string{"myTag2", "Personal"},
				Body:        "\nHello world in many lines text 2!\nHere goes something \n\nand here!\n\nLorem ipsum",
			},
			{
				Title:       "Post 3",
				Description: "My test description 3",
				Tags:        []string{},
				Body:        "Hello world in many lines text 3!\nHere goes something \n\nand here!\n\nLorem ipsum\n",
			},
		}
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstPost)},
			"hello-world2.md": {Data: []byte(secondPost)},
			"hello-world3.md": {Data: []byte(thirdPost)},
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
