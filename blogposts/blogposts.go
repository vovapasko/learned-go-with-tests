package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type BlogPost struct {
}

func NewPostsFromFs(filesystem fstest.MapFS) []BlogPost {
	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []BlogPost
	for range dir {
		posts = append(posts, BlogPost{})
	}
	return posts
}
