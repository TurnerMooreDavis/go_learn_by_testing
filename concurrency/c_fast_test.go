package concurrency

import (
	"testing"
)

func BenchmarkCheckWebsitesFast(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsitesWAnonymousFunc(slowStubWebsiteChecker, urls)
	}
}
