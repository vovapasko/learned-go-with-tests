package concurrency

type WebsiteChecker func(string) bool

type result struct {
	name  string
	value bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChanel := make(chan result)
	for _, url := range urls {
		go func(url string) {
			resultsChanel <- result{url, wc(url)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsChanel
		results[r.name] = r.value
	}

	return results
}
