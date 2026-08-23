package main

type WebsiteChecker func(url string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()

		// results[url] = wc(url)
	}

	for range len(urls) {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
