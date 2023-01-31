package blogposts

import (
	"testing/fstest"
)

type BlogPost struct {
}

func NewPostsFromFs(fs fstest.MapFS) []BlogPost {
	return nil
}
