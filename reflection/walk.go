package reflection

func Walk(x interface{}, fn func(string)) {
	fn("hello")
}
