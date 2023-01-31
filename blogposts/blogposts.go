package blogposts

import (
	"errors"
	"io"
	"io/fs"
	"strings"
)

const IndexOutOfRangeError = "Index out of range"

type BlogPost struct {
	Title string
}

func NewPostsFromFs(filesystem fs.FS) ([]BlogPost, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []BlogPost
	for _, filename := range dir {
		post, err := getPost(filesystem, filename)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, filename fs.DirEntry) (BlogPost, error) {
	postFile, err := filesystem.Open(filename.Name())
	if err != nil {
		return BlogPost{}, err
	}
	defer postFile.Close()

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return BlogPost{}, err
	}
	title, err := extractTitle(string(postData))
	if err != nil {
		return BlogPost{}, err
	}
	post := BlogPost{Title: title}
	return post, nil
}

func extractTitle(data string) (string, error) {
	splitString := strings.Split(data, "Title: ")
	if len(splitString) < 2 {
		return "", errors.New(IndexOutOfRangeError)
	}
	return splitString[1], nil
}
