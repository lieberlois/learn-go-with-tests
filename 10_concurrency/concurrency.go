package concurrency

// WebsiteChecker is a function that checks a website for availability
type WebsiteChecker func(string) bool

type result struct {
    string
    bool
}

// CheckWebsites checks a slice of urls with a given WebsiteChecker
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resChannel := make(chan result)

    for _, url := range urls {
        go func(u string) {
            resChannel <- result{u, wc(u)}
        }(url)
    }
	
	for i := 0; i < len(urls); i++ {
		res := <- resChannel
		results[res.string] = res.bool
	}
	close(resChannel)
    return results
}