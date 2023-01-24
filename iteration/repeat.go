package iteration

func Repeat(text string, iterationAmount int) string {
	base := text
	for i := 0; i < iterationAmount; i++ {
		base += text
	}
	return base
}
