package auth0

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"mrsydar/apiserver/configs/constants/envnames"
)

var TokenFetchURL string
var CallbackEndpoint string
var Secret string
var GetDataForTokenFetchWithCode func(code string) url.Values

func Init() error {
	log.Print("Initializing Auth0 variables")

	TokenFetchURL = fmt.Sprintf("https://%s/oauth/token", os.Getenv(envnames.Auth0Domain))
	CallbackEndpoint = os.Getenv(envnames.Auth0CallbackEndpoint)
	Secret = os.Getenv(envnames.Auth0ClientSecret)

	defaultData := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {os.Getenv(envnames.Auth0ClientID)},
		"client_secret": {os.Getenv(envnames.Auth0ClientSecret)},
		"redirect_uri":  {os.Getenv(envnames.Auth0CallbackURL)},
	}

	GetDataForTokenFetchWithCode = func(code string) url.Values {
		data := defaultData
		data["code"] = []string{code}
		return data
	}

	return nil
}
