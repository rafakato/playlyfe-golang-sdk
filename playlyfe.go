package playlyfe

import (
  "code.google.com/p/goauth2/oauth"
  "crypto/rand"
  "errors"
  "fmt"
  "os"
)

const authorizationUrl string = "https://playlyfe.com/auth"
const tokenUrl string = "https://playlyfe.com/auth/token"

type PlaylyfeClientConfiguration struct {
  ClientId     string
  ClientSecret string
  RedirectUrl  string
  Type         string
  Code         string
  CacheFile    string
}

func Client(configuration PlaylyfeClientConfiguration) (client *PlaylyfeClient, accessUrl string, err error) {
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
  }
  transport := &oauth.Transport{Config: config}

  token, err := config.TokenCache.Token()
  if err != nil {
    token, accessUrl, err = getToken(configuration, transport, config)
    if err != nil {
      return nil, accessUrl, err
    } else if token == nil {
      return nil, "", errors.New("There was a error when trying to get the token")
    }
  }

  if token.Expired() {
    token, accessUrl, err = getToken(configuration, transport, config)
    if err != nil {
      return nil, accessUrl, err
    } else if token == nil {
      return nil, "", errors.New("There was a error when trying to get the token")
    }
  }

  transport.Token = token

  httpClient := transport.Client()

  client = &PlaylyfeClient{http: *httpClient}

  return client, "", nil
}

func getToken(configuration PlaylyfeClientConfiguration, transport *oauth.Transport, config *oauth.Config) (token *oauth.Token, authUrl string, err error) {
  if configuration.ClientId == "" || configuration.ClientSecret == "" {
    fmt.Fprint(os.Stderr)
    os.Exit(2)
  }

  switch configuration.Type {
  case "code":
    if configuration.Code == "" {
      url := config.AuthCodeURL("")
      return nil, url, errors.New("You must get and code to access the api")
    }

    token, err = transport.Exchange(configuration.Code)
    if err != nil {
      return nil, "", err
    }
  case "client":
    err = transport.AuthenticateClient()
    if err != nil {
      return nil, "", err
    }

    token = transport.Token
  default:
    return nil, "", errors.New("You must set either 'code' or 'client'")
  }

  err = config.TokenCache.PutToken(token)
  if err != nil {
    return nil, "", err
  }
  return token, "", nil
}
