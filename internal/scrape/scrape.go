package scrape

import (
	"craigslist-auto-apply/internal/socks5"
	"fmt"
	"github.com/corpix/uarand"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/gorilla/websocket"
	"regexp"
	"strings"
	"time"
)

var (
	selectors        = []string{".result-row .result-image", "#sortable-results > ul > li:nth-child(1) > p > a"}
	emailTagSelector = "body > section > section > header > div.reply-button-row > button"
	c                = colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(false),
	)
)

type CollyScraperInterface interface {
	baseCollector() *colly.Collector
}
type CollyScrape struct {
	ActiveRequestMap    map[string]time.Time
	CompletedRequestMap map[string]time.Duration
	requestCount        int
	errors              []error
	contactInfos        []string
	links               []string
	ListingURLS         []string
	C                   *colly.Collector
	U                   *websocket.Conn
}
func baseCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.IgnoreRobotsTxt())

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*.craigslist.org.*",
		Parallelism: 1,
		RandomDelay: 30 * time.Second,
	})
	return collector
}

func getProxyFunc(proxyService *socks5.Service) (colly.ProxyFunc, error) {
	newServers := proxyService.RotateServers(30)
	fmt.Printf("rotating servers")
	proxyFunc, err := proxy.RoundRobinProxySwitcher(newServers...)
	if err != nil {
		return nil, err
	}
	return proxyFunc, nil
}
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}


//BuildCollector ...
func (s *CollyScrape) BuildCollector() {

	//socks5Service := &socks5.Service{RotatingServers: false}
	//
	//proxyFunc, err := getProxyFunc(socks5Service)
	//if err != nil {
	//	fmt.Println("can't set proxy")
	//}
	//collector.SetProxyFunc(proxyFunc)
	collector := baseCollector()

	collector.OnError(func(r *colly.Response, err error) {
		s.errors = append(s.errors, err)
		time.Sleep(5 * time.Second)
		c.Visit(r.Request.URL.String())
	})
	collector.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
		listingURL := e.Attr("href")
		if !contains(s.ListingURLS, listingURL) {
			s.ListingURLS = append(s.ListingURLS, listingURL)
			s.U.WriteJSON(SocketMessage{MessageType: "listings", Payload: string(listingURL)})
			c.Visit(listingURL)

		}
	})

	collector.OnHTML(".reply-button.js-only", func(e *colly.HTMLElement) {
		info := e.Attr("data-href")
		r, _ := regexp.Compile("https://([a-z]+).craigslist.org")
		contactInfoTop := fmt.Sprintf("%s/contactinfo/", r.FindString(e.Request.URL.String()))
		info = strings.Replace(info, "/__SERVICE_ID__/", contactInfoTop, 1)
		s.contactInfos = append(s.contactInfos, info)

	})
	//c.OnHTML("button.reply-button.js-only", func(e *colly.HTMLElement) {
	//	fmt.Println(e)
	//	//doc.Find("form").Each(func(i int, formDoc *goquery.Selection) {
	//	//	if loginFormSelection != nil {
	//	//		return
	//	//	}
	//	//	formDoc.Find("input").Each(func(_ int, inputDoc *goquery.Selection) {
	//	//		if loginFormSelection != nil {
	//	//			return
	//	//		}
	//	//		if name, ok := inputDoc.Attr("name"); ok {
	//	//			if strings.Contains(strings.ToLower(name), "pass") {
	//	//				loginFormSelection = formDoc
	//	//			}
	//	//		}
	//	//	})
	//	//})
	//})


	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Del("User-Agent")
		r.Headers.Add("User-Agent", uarand.GetRandom())
		s.ActiveRequestMap[r.URL.String()] = time.Now()

	})

	collector.OnResponse(func(r *colly.Response) {
		url := r.Request.URL.String()
		startTime := s.ActiveRequestMap[url]
		delete(s.ActiveRequestMap,url)
		s.CompletedRequestMap[url] = startTime.Sub(time.Now()) * time.Millisecond
	})

	s.C = collector
}

//

//GetListingURLS comment
func GetListingURLS(stateCodes []string, con websocket.Conn) {

	for i := range stateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", stateCodes[i])
		clQuery := fmt.Sprintf("%s/search/sof", stateOrg)
		percentComplete := fmt.Sprintf("%f", (float64(i)/float64(len(stateCodes)))*100)
		con.WriteJSON(SocketMessage{MessageType: "listingPercentComplete", Payload: percentComplete})
		c.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
			listingURL := e.Attr("href")
			con.WriteJSON(SocketMessage{MessageType: "listingURLs", Payload: listingURL})

		})
		c.Visit(clQuery)

	}

}

//GetContactInfoURLS comment
//func GetContactInfoURLS(link string) string {
//	fmt.Println(link)
//	var emailLink string
//
//}

//func GetContactInfoURLS(listings []Listing) []Listing {
//
//
//	for i := range listings {
//		c.OnHTML("body > section > section > header > div.reply-button-row > button", func(e *colly.HTMLElement) {
//			listings[i].GetContactInfoURLS = e.Attr("data-href")
//		})
//		c.Visit(listings[i].ListingUrl)
//		c.Wait()
//	}
//	return listings
//}
