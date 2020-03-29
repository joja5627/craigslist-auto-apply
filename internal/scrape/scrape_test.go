package scrape

import (
	"craigslist-auto-apply/internal"
	"fmt"
	"testing"
	"time"
)

func TestShouldScrapeCL(t *testing.T) {
	cScrape := CollyScrape{
		ActiveRequestMap:    make(map[string]time.Time),
		CompletedRequestMap: make(map[string]time.Duration),
	}
	cScrape.BuildCollector()

	startTime := time.Now()
	for _, state := range internal.StateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
		cScrape.C.Visit(fmt.Sprintf("%s/d/software-qa-dba-etc/search/sof", stateOrg))
		cScrape.C.Wait()
		cScrape.C.Visit(fmt.Sprintf("%s/search/sof?employment_type=3", stateOrg))
		cScrape.C.Wait()

	}
	totalTime := time.Now().Sub(startTime)
	fmt.Println(totalTime.Seconds())
	fmt.Println(len(cScrape.ListingURLS))
}