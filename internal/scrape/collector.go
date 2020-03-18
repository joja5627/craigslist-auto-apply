package scrape

import (
	"craigslist-auto-apply/internal/socks5"
	"fmt"
	"github.com/corpix/uarand"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/prometheus/client_golang/prometheus"
	"regexp"
	"strings"
	"sync"
	"time"
)




type CollyScraperInterface interface {
	baseCollector() *colly.Collector
}

// Exporter represents an instance of the Netgear cable modem exporter.
type CollyExporter struct {
	url, authHeaderValue string

	mu sync.Mutex

	// Exporter metrics.
	totalScrapes prometheus.Counter
	scrapeErrors prometheus.Counter

	// Downstream metrics.
	dsChannelSNR               *prometheus.Desc
	dsChannelPower             *prometheus.Desc
	dsChannelCorrectableErrs   *prometheus.Desc
	dsChannelUncorrectableErrs *prometheus.Desc

	// Upstream metrics.
	usChannelPower      *prometheus.Desc
	usChannelSymbolRate *prometheus.Desc
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


//BuildCollector ...
func BuildCollector() *colly.Collector {

	//socks5Service := &socks5.Service{RotatingServers: false}
	//
	//proxyFunc, err := getProxyFunc(socks5Service)
	//if err != nil {
	//	fmt.Println("can't set proxy")
	//}
	//collector.SetProxyFunc(proxyFunc)
	collector := baseCollector()

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("error: ", err.Error())
		fmt.Println("errors: ", len(errors))
		errors = append(errors, err.Error())
		time.Sleep(5 * time.Second)
		c.Visit(r.Request.URL.String())
	})
	collector.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
		listingURL := e.Attr("href")
		c.Visit(listingURL)
	})
	collector.OnHTML("button.reply-button.js-only", func(e *colly.HTMLElement) {
		info := e.Attr("data-href")
		r, _ := regexp.Compile("https://([a-z]+).craigslist.org")
		contactInfoTop := fmt.Sprintf("%s/contactinfo/", r.FindString(e.Request.URL.String()))
		info = strings.Replace(info, "/__SERVICE_ID__/", contactInfoTop, 1)
		contactInfos = append(contactInfos, info)

	})
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Del("User-Agent")
		r.Headers.Add("User-Agent", uarand.GetRandom())

	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("listing urls: ", len(links))
		//fmt.Println(r.Request)
		//fmt.Println(r.StatusCode)
		//fmt.Println(r.Body)
		//fmt.Println("======")
	})

	return collector
}
