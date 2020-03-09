package main

import (
	"craigslist-auto-apply/internal"
	"craigslist-auto-apply/internal/scrape"
	"fmt"
	"time"
)

func main(){
	c := scrape.BuildCollector()

	for _, state := range internal.StateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
		time.Sleep(10 * time.Millisecond)
		c.Visit(fmt.Sprintf("%s/d/software-qa-dba-etc/search/sof", stateOrg))
		time.Sleep(10 * time.Millisecond)
		c.Visit(fmt.Sprintf("%s/search/sof?employment_type=3", stateOrg))
	}
	c.Wait()
}
