package main

import (
	"craigslist-auto-apply/internal"
	"craigslist-auto-apply/internal/scrape"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
	"google.golang.org/api/gmail/v1"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"github.com/googollee/go-socket.io"

)

var (
	listings   []string
	newline    = []byte{'\n'}
	space      = []byte{' '}
	stateCodes = []string{"auburn", "bham", "dothan", "shoals", "gadsden", "huntsville", "mobile", "montgomery", "tuscaloosa", "anchorage", "fairbanks", "kenai", "juneau", "flagstaff", "mohave", "phoenix", "prescott", "showlow", "sierravista", "tucson", "yuma", "fayar", "fortsmith", "jonesboro", "littlerock", "texarkana", "bakersfield", "chico", "fresno", "goldcountry", "hanford", "humboldt", "imperial", "inlandempire", "losangeles", "mendocino", "merced", "modesto", "monterey", "orangecounty", "palmsprings", "redding", "sacramento", "sandiego", "sfbay", "slo", "santabarbara", "santamaria", "siskiyou", "stockton", "susanville", "ventura", "visalia", "yubasutter", "boulder", "cosprings", "denver", "eastco", "fortcollins", "rockies", "pueblo", "westslope", "newlondon", "hartford", "newhaven", "nwct", "delaware", "washingtondc", "miami", "daytona", "keys", "fortlauderdale", "fortmyers", "gainesville", "cfl", "jacksonville", "lakeland", "miami", "lakecity", "ocala", "okaloosa", "orlando", "panamacity", "pensacola", "sarasota", "miami", "spacecoast", "staugustine", "tallahassee", "tampa", "treasure", "miami", "albanyga", "athensga", "atlanta", "augusta", "brunswick", "columbusga", "macon", "nwga", "savannah", "statesboro", "valdosta", "honolulu", "boise", "eastidaho", "lewiston", "twinfalls", "bn", "chambana", "chicago", "decatur", "lasalle", "mattoon", "peoria", "rockford", "carbondale", "springfieldil", "quincy", "bloomington", "evansville", "fortwayne", "indianapolis", "kokomo", "tippecanoe", "muncie", "richmondin", "southbend", "terrehaute", "ames", "cedarrapids", "desmoines", "dubuque", "fortdodge", "iowacity", "masoncity", "quadcities", "siouxcity", "ottumwa", "waterloo", "lawrence", "ksu", "nwks", "salina", "seks", "swks", "topeka", "wichita", "bgky", "eastky", "lexington", "louisville", "owensboro", "westky", "batonrouge", "cenla", "houma", "lafayette", "lakecharles", "monroe", "neworleans", "shreveport", "maine", "annapolis", "baltimore", "easternshore", "frederick", "smd", "westmd", "boston", "capecod", "southcoast", "westernmass", "worcester", "annarbor", "battlecreek", "centralmich", "detroit", "flint", "grandrapids", "holland", "jxn", "kalamazoo", "lansing", "monroemi", "muskegon", "nmi", "porthuron", "saginaw", "swmi", "thumb", "up", "bemidji", "brainerd", "duluth", "mankato", "minneapolis", "rmn", "marshall", "stcloud", "gulfport", "hattiesburg", "jackson", "meridian", "northmiss", "natchez", "columbiamo", "joplin", "kansascity", "kirksville", "loz", "semo", "springfield", "stjoseph", "stlouis", "billings", "bozeman", "butte", "greatfalls", "helena", "kalispell", "missoula", "montana", "grandisland", "lincoln", "northplatte", "omaha", "scottsbluff", "elko", "lasvegas", "reno", "nh", "cnj", "jerseyshore", "newjersey", "southjersey", "albuquerque", "clovis", "farmington", "lascruces", "roswell", "santafe", "albany", "binghamton", "buffalo", "catskills", "chautauqua", "elmira", "fingerlakes", "glensfalls", "hudsonvalley", "ithaca", "longisland", "newyork", "oneonta", "plattsburgh", "potsdam", "rochester", "syracuse", "twintiers", "utica", "watertown", "asheville", "boone", "charlotte", "eastnc", "fayetteville", "greensboro", "hickory", "onslow", "outerbanks", "raleigh", "wilmington", "winstonsalem", "bismarck", "fargo", "grandforks", "nd", "akroncanton", "ashtabula", "athensohio", "chillicothe", "cincinnati", "cleveland", "columbus", "dayton", "limaohio", "mansfield", "sandusky", "toledo", "tuscarawas", "youngstown", "zanesville", "lawton", "enid", "oklahomacity", "stillwater", "tulsa", "bend", "corvallis", "eastoregon", "eugene", "klamath", "medford", "oregoncoast", "portland", "roseburg", "salem", "altoona", "chambersburg", "erie", "harrisburg", "lancaster", "allentown", "meadville", "philadelphia", "pittsburgh", "poconos", "reading", "scranton", "pennstate", "williamsport", "york", "providence", "charleston", "columbia", "florencesc", "greenville", "hiltonhead", "myrtlebeach", "nesd", "csd", "rapidcity", "siouxfalls", "sd", "chattanooga", "clarksville", "cookeville", "jacksontn", "knoxville", "memphis", "nashville", "tricities", "abilene", "amarillo", "austin", "beaumont", "brownsville", "collegestation", "corpuschristi", "dallas", "nacogdoches", "delrio", "elpaso", "galveston", "houston", "killeen", "laredo", "lubbock", "mcallen", "odessa", "sanangelo", "sanantonio", "sanmarcos", "bigbend", "texoma", "easttexas", "victoriatx", "waco", "wichitafalls", "logan", "ogden", "provo", "saltlakecity", "stgeorge", "vermont", "charlottesville", "danville", "fredericksburg", "norfolk", "harrisonburg", "lynchburg", "blacksburg", "richmond", "roanoke", "swva", "winchester", "bellingham", "kpr", "moseslake", "olympic", "pullman", "seattle", "skagit", "spokane", "wenatchee", "yakima", "charlestonwv", "martinsburg", "huntington", "morgantown", "wheeling", "parkersburg", "swv", "wv", "appleton", "eauclaire", "greenbay", "janesville", "racine", "lacrosse", "madison", "milwaukee", "northernwi", "sheboygan", "wausau", "wyoming", "micronesia", "puertorico", "virgin"}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getRequest(rawUrl string, c *colly.Context) *colly.Request {
	h := http.Header{}
	url, _ := url.Parse(rawUrl)
	request := colly.Request{URL: url, Headers: &h, Ctx: c}
	return &request

}
func scrapeTestCL( w http.ResponseWriter, r *http.Request) {


	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	con.WriteJSON(scrape.SocketMessage{MessageType: "listingPercentComplete", Payload: "100"})
	for _,l := range listings{
		con.WriteJSON(scrape.SocketMessage{MessageType: "listings", Payload: l})
	}
	//for i, _ := range internal.StateCodes {
	//	percentComplete := fmt.Sprintf("%f", (float64(i)/float64(len(stateCodes)))*100)
	//}

}
func scrapeCL( w http.ResponseWriter, r *http.Request) {
	cScrape := scrape.CollyScrape{
		ActiveRequestMap:    make(map[string]time.Time),
		CompletedRequestMap: make(map[string]time.Duration),
	}
	cScrape.BuildCollector()

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	cScrape.U = con
	startTime := time.Now()
	//	emailString := r.FindString(string(htmlData))
	//	listing.Email = strings.Trim(emailString, ":")
	for i, state := range internal.StateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
		cScrape.C.Visit(fmt.Sprintf("%s/d/software-qa-dba-etc/search/sof", stateOrg))
		cScrape.C.Wait()
		cScrape.C.Visit(fmt.Sprintf("%s/search/sof?employment_type=3", stateOrg))
		cScrape.C.Wait()
		percentComplete := fmt.Sprintf("%f", (float64(i)/float64(len(stateCodes)))*100)
		con.WriteJSON(scrape.SocketMessage{MessageType: "listingPercentComplete", Payload: percentComplete})
	}
	//contactInfo := scrape.GetContactInfoURL(listing)
	//if contactInfo == "" {
	//	body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),"statusText":http.StatusText(http.StatusBadRequest),"error": "no contact info for link"}
	//	c.JSON(http.StatusBadRequest,body)
	//	return
	//}
	//contactInfos := scrape.GetContactInfos()
	//fmt.Println(len(contactInfos))
	//for _, contactURL := range contactInfos {
	//	infoRESP, err := http.Get(contactURL)
	//	htmlData, err := ioutil.ReadAll(infoRESP.Body)
	//	if err != nil {
	//		jsonListing, _ := json.Marshal(listing)
	//		body := map[string]string{"status": fmt.Sprintf("%d", http.StatusBadRequest),
	//			"statusText": http.StatusText(http.StatusBadRequest),
	//			"body":       string(jsonListing)}
	//		c.JSON(http.StatusBadRequest, body)
	//		return
	//	}
	//
	//}
	//listing.ContactInfoUrl = contactInfo
	//r, _ := regexp.Compile(":([a-zA-Z0-9])+@job.craigslist.org")
	//
	////
	////
	//if err != nil {
	//	jsonListing, _ := json.Marshal(listing)
	//	body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
	//		"statusText":http.StatusText(http.StatusBadRequest),
	//		"body": string(jsonListing)}
	//	c.JSON(http.StatusBadRequest,body)
	//	return
	//}
	totalTime := time.Now().Sub(startTime)
	fmt.Println(totalTime.Seconds())
	fmt.Println(len(cScrape.ListingURLS))


}

func buildEmail(email string, url string) gmail.Message {
	var message gmail.Message

	temp := []byte("From: 'me'\r\n" +
		fmt.Sprintf("To: %s \r\n", email) +
		"Subject: Software Position \r\n" +
		"\r\nHey!\r\n" + " My name is Joe Jackson and I'm interested in applying for the position you posted on craigs list." +
		" This is a link to my most up to date resume https://docs.google.com/document/d/1ugz6WqXaWEj2s4CLRC5ecz40RiRUmfC9XxmvW-TSXwA/edit?usp=sharing " +
		fmt.Sprintf("%s", url) +
		"\r\nBest," + "\r\nJoe Jackson")

	message.Raw = base64.StdEncoding.EncodeToString(temp)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "=", "", -1)
	return message
}

//https://sandiego.craigslist.org/nsd/sof/d/carlsbad-full-stack-web-developer/6955927244.html
//"https://sandiego.craigslist.org/contactinfo/sdo/sof/6955927244

func main() {



	f, err := os.Open("/Users/joejackson/dev/craigslist-auto-apply/config/listings.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&listings)
	if err != nil {
		panic(err)
	}

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/test_scrape", "test_scrape", func(s socketio.Conn) string {
		s.SetContext("test_scrape")
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8080", nil))



	//gin.SetMode(gin.DebugMode)
	//r := gin.New()
	//r.Use(cors.Default())
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//r.GET("/scrape", func(c *gin.Context) {scrapeCL(c.Writer, c.Request)})
	//r.GET("/scrape_test", func(c *gin.Context) {scrapeTestCL(c.Writer, c.Request)})
	//r.Run()


	//})
	//c := scrape.BuildCollector()
	//
	//for _, state := range stateCodes {
	//	stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
	//	time.Sleep(10 * time.Millisecond)
	//	c.Visit(fmt.Sprintf("%s/d/software-qa-dba-etc/search/sof", stateOrg))
	//	//c.Wait()
	//	//time.Sleep(10 * time.Millisecond)
	//	//c.Visit(fmt.Sprintf("%s/search/sof?employment_type=3", stateOrg))
	//	//c.Wait()
	//}



	//
	//if htmlData == nil {
	//	jsonListing, _ := json.Marshal(listing)
	//	body := map[string]string{"status": fmt.Sprintf("%d", http.StatusBadRequest),
	//		"statusText": http.StatusText(http.StatusBadRequest),
	//		"body":       string(jsonListing)}
	//	c.JSON(http.StatusBadRequest, body)
	//	return
	//}

	//listingMAP := scrape.GetListingMap()
	//for k, _ := range listingMAP.Keys() {
	//
	//}
	//for key,val := range {
	//	fmt.Println("%s",key)
	//	for link := range val {
	//		fmt.Println("  %s",link)
	//	}
	//
	//}

}

//c.Wait()
//



//r.POST("/sendEmail", func(c *gin.Context) {
//
//
//	listing := scrape.Listing{}
//	err := json.NewDecoder(c.Request.Body).Decode(&listing)
//	if err != nil {
//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),"statusText":http.StatusText(http.StatusBadRequest),"error":err.Error()}
//		c.JSON(http.StatusBadRequest,body)
//		return
//	}
//
//

//
//
//	if infoRESP == nil {
//		jsonListing, _ := json.Marshal(listing)
//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
//			"statusText":http.StatusText(http.StatusBadRequest),
//			"body": string(jsonListing)}
//		c.JSON(http.StatusBadRequest,body)
//		return
//	}

//
//
//	if err != nil {
//		jsonListing, _ := json.Marshal(listing)
//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
//			"statusText":http.StatusText(http.StatusBadRequest),
//			"body": string(jsonListing)}
//		c.JSON(http.StatusBadRequest,body)
//		return
//	}
//

//
//
//	if listing.Email  != ""{
//
//		queryString := fmt.Sprintf("in:sent %s ",listing.Url)
//
//		messages,err := srv.Users.Messages.List("me").Q(queryString).MaxResults(10000).Do()
//
//		if len(messages.Messages) > 0 {
//			jsonListing, _ := json.Marshal(listing)
//			body := map[string]string{
//				"status": fmt.Sprintf("%d",http.StatusAccepted),
//				"statusText":http.StatusText(http.StatusAccepted),
//				"body": string(jsonListing)}
//			c.JSON(http.StatusAccepted,body)
//			return
//		}
//
//		clEmail := buildEmail(listing.Email,listing.Url)
//
//		emailResponse, err := srv.Users.Messages.Send("me",&clEmail).Do()
//
//		if err != nil {
//			body := map[string]string{
//				"status":fmt.Sprintf("%d",http.StatusInternalServerError),
//				"statusText":http.StatusText(http.StatusInternalServerError),
//				"error": err.Error()}
//			c.JSON(http.StatusInternalServerError,body)
//			return
//
//		}else {
//			listing.EmailResponse =  emailResponse.Raw
//			jsonListing, _ := json.Marshal(listing)
//			body := map[string]string{"status":fmt.Sprintf("%d", http.StatusCreated),
//				"statusText":http.StatusText(http.StatusCreated),
//				"body": string(jsonListing)}
//			c.JSON(http.StatusCreated,body)
//			return
//
//		}
//	}else {
//		jsonListing, _ := json.Marshal(listing)
//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
//			"statusText":http.StatusText(http.StatusBadRequest),
//			"body": string(jsonListing)}
//		c.JSON(http.StatusBadRequest,body)
//		return
//	}
//})
//r.Run()
