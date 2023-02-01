package blogposts

const (
	WrongBlogPostFileFormatError = BlogPostError("wrong format of the file to convert to blog post")
)

type BlogPostError string

func (err BlogPostError) Error() string {
	return string(err)
}
