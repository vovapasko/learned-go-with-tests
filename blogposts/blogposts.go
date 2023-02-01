package blogposts

import (
	"io"
	"io/fs"
	"strings"
)

type BlogPost struct {
	Title, Description string
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

	defer func(postFile fs.File) {
		err := postFile.Close()
		if err != nil {

		}
	}(postFile)

	post, err := createPost(postFile)

	return post, err
}

func createPost(r io.Reader) (BlogPost, error) {
	postData, err := io.ReadAll(r)
	if err != nil {
		return BlogPost{}, err
	}
	title, err := extractTitle(string(postData))
	if err != nil {
		return BlogPost{}, err
	}
	description, err := extractDescription(string(postData))
	if err != nil {
		return BlogPost{Title: title}, err
	}
	return BlogPost{Title: title, Description: description}, nil
}

func extractTitle(data string) (string, error) {
	_, afterTitle, found := strings.Cut(data, "Title: ")
	if !found {
		return "", WrongBlogPostFileFormatError
	}
	splitString := strings.Split(afterTitle, "\n")
	return splitString[0], nil
}

func extractDescription(data string) (string, error) {
	_, afterTitle, found := strings.Cut(data, "Description: ")
	if !found {
		return "", WrongBlogPostFileFormatError
	}
	splitString := strings.Split(afterTitle, "\n")
	return splitString[0], nil
}
