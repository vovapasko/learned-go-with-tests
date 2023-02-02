package blogposts

import (
	"io"
	"io/fs"
	"strings"
)

type BlogPost struct {
	Title, Description, Body string
	Tags                     []string
}

const (
	tagSeparator   = ", "
	bodySeparator  = "---\n"
	titleKey       = "Title: "
	descriptionKey = "Description: "
	tagsKey        = "Tags: "
)

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

	postMetadataMap := extractPostMetadata(string(postData))

	body, err := extractBody(string(postData))
	if err != nil {
		return BlogPost{
			Title: postMetadataMap[titleKey], Description: postMetadataMap[descriptionKey],
			Tags: parseTags(postMetadataMap[tagsKey], tagSeparator),
		}, err
	}
	return BlogPost{
		Title: postMetadataMap[titleKey], Description: postMetadataMap[descriptionKey],
		Tags: parseTags(postMetadataMap[tagsKey], tagSeparator), Body: body,
	}, nil
}

func extractPostMetadata(data string) (metadataMap map[string]string) {
	metadataMap = make(map[string]string)
	title, _ := extractTitle(data)
	description, _ := extractDescription(data)
	tags, _ := extractTags(data)
	metadataMap[titleKey] = title
	metadataMap[descriptionKey] = description
	metadataMap[tagsKey] = tags
	return
}

func extractBody(data string) (string, error) {
	_, afterSeparator, found := strings.Cut(data, bodySeparator)
	if !found {
		return "", WrongBlogPostFileFormatError
	}
	return afterSeparator, nil
}

func extractTags(data string) (string, error) {
	rawTags, err := extractByTheKey(data, tagsKey)
	if err != nil {
		return "", err
	}
	return rawTags, nil
}

func parseTags(data, tagSeparator string) (result []string) {
	splitted := strings.Split(data, tagSeparator)
	if splitted[0] == "" {
		return make([]string, 0)
	}
	return splitted
}

func extractTitle(data string) (string, error) {
	return extractByTheKey(data, titleKey)
}

func extractDescription(data string) (string, error) {
	return extractByTheKey(data, descriptionKey)
}

func extractByTheKey(data string, key string) (string, error) {
	_, afterTitle, found := strings.Cut(data, key)
	if !found {
		return "", WrongBlogPostFileFormatError
	}
	splitString := strings.Split(afterTitle, "\n")
	return splitString[0], nil
}
