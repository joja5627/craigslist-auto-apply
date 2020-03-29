package email

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
)

// smtpServer data to smtp server.
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server.
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
// EmailServiceImpl holds objects and functions for gmail API calls
type EmailServiceImpl struct {
	Srv  *gmail.Service
	User string
}

// NewEmailService creates a new Provider
func NewEmailService(user string) *EmailServiceImpl {
	s,err := newService()
	if err != nil {
		panic(err)
	}
	return &EmailServiceImpl{
		Srv:  s,
		User: user,
	}
}
//func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
//	cacheFile, err := tokenCacheFile()
//	if err != nil {
//		log.Fatalf("Unable to get path to cached credential file. %v", err)
//	}
//	tok, err := tokenFromFile(cacheFile)
//	if err != nil {
//		tok = getTokenFromWeb(config)
//		saveToken(cacheFile, tok)
//	}
//	return config.Client(ctx, tok)
//}

func newService() (service *gmail.Service, err error){
	ctx := context.Background()
	json, err := ioutil.ReadFile("../../cmd/craigslistAPI/credentials.json")
	if err != nil {
		log.Printf("[ERROR] Failed to process read file: %s", err)
		return nil, err
	}
	return gmail.NewService(ctx, option.WithCredentialsJSON(json))
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}


// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
//func GetClient(ctx context.Context, config *oauth2.Config) *http.Client {
//	cacheFile, err := tokenCacheFile()
//	if err != nil {
//		log.Fatalf("Unable to get path to cached credential file. %v", err)
//	}
//	tok, err := tokenFromFile(cacheFile)
//	if err != nil {
//		tok = getTokenFromWeb(config)
//		saveToken(cacheFile, tok)
//	}
//	return config.Client(ctx, tok)
//}


// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
//func tokenCacheFile() (string, error) {
//	usr, err := user.Current()
//	if err != nil {
//		return "", err
//	}
//	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
//	os.MkdirAll(tokenCacheDir, 0700)
//	return filepath.Join(tokenCacheDir,
//		url.QueryEscape("gmail-go-quickstart.json")), err
//}

