package blogposts

import (
	"io/fs"
)

type BlogPost struct {
}

func NewPostsFromFs(filesystem fs.FS) ([]BlogPost, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []BlogPost
	for range dir {
		posts = append(posts, BlogPost{})
	}
	return posts, nil
}
