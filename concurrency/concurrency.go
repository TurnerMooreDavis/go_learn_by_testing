package concurrency

type WebsiteChecker func(string) bool

func asyncGet(m *chan result, url string, wc WebsiteChecker) {
	(*m) <- result{url, wc(url)}
}

type result struct {
	url string
	res bool
}
type resultAnon struct {
	string
	bool
}

func CheckWebsitesSlow(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		// go func() {
		// 	results[url] = wc(url)
		// }()
		go asyncGet(&resultsChannel, url, wc)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.url] = r.res
	}
	return results
}

func CheckWebsitesWAnonymousFunc(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan resultAnon)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- resultAnon{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
