package blogposts

const (
	IndexOutOfRangeError         = "index out of range"
	WrongBlogPostFileFormatError = "wrong format of the file to convert to blog post"
)

type BlogPostError string

func (err BlogPostError) Error() string {
	return string(err)
}
