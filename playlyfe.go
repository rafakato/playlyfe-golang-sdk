package playlyfe

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"log"
	"os"
)

const authorizationUrl string = "https://api.playlyfe.com/v1/auth"
const tokenUrl string = "https://playlyfe.com/auth/token"

type PlaylyfeClientConfiguration struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
	Type         string
	CacheFile    string
}

func Client(configuration PlaylyfeClientConfiguration) (client PlaylyfeClient, err error) {
	if configuration.Type == "" {
		configuration.Type = "code"
	}

	config := &oauth.Config{
		ClientId:     configuration.ClientId,
		ClientSecret: configuration.ClientSecret,
		RedirectURL:  configuration.RedirectUrl,
		AuthURL:      authorizationUrl,
		TokenURL:     tokenUrl,
		TokenCache:   oauth.CacheFile(configuration.CacheFile),
		AccessType:   configuration.Type,
	}
	transport := &oauth.Transport{Config: config}

	token, err := config.TokenCache.Token()
	if err != nil {
		token, err = getToken(configuration, transport, config)
	}

	if token.Expired() {
		token, err = getToken(configuration, transport, config)
	}

	transport.Token = token

	httpClient := transport.Client()

	client = PlaylyfeClient{http: *httpClient}

	return client, nil
}

func getToken(configuration PlaylyfeClientConfiguration, transport *oauth.Transport, config *oauth.Config) (token *oauth.Token, err error) {
	if configuration.ClientId == "" || configuration.ClientSecret == "" {
		fmt.Fprint(os.Stderr)
		os.Exit(2)
	}

	switch configuration.Type {
	case "code":
		//TODO: Needs implementing of code based authentication
		/*if *code == "" {
		      url := config.AuthCodeURL("")
		      fmt.Print("Visit this URL to get a code, then run again with -code=YOUR_CODE\n\n")
		      fmt.Println(url)
		      return
		  }

		  token, err = transport.Exchange(*code)
		  if err != nil {
		      log.Fatal("Exchange:", err)
		  }*/
	case "client":
		err = transport.AuthenticateClient()
		if err != nil {
			log.Fatal("Authenticate Client:", err)
		}

		token = transport.Token
	default:
		log.Fatal("You must set either 'code' or 'client'")
	}

	err = config.TokenCache.PutToken(token)
	if err != nil {
		log.Fatal("Put Token:", err)
	}
	return
}
