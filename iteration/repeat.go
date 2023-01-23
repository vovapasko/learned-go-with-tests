package iteration

func Repeat(text string) string {
	base := text
	for i := 0; i < 5; i++ {
		base += text
	}
	return base
}
