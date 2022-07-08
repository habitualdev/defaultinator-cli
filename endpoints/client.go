package endpoints

import (
	"net/http"
	"runtime"
)

const Version = "1.0"
const BaseUrl = "https://api.defaultinator.com"

var QueryFields = []string{"vendor", "product", "version", "username", "password", "part", "field"}

type Client struct {
	ApiKey  string
	Client  *http.Client
	Request *http.Request
	BaseUrl string
}

type CPE struct {
	Cpe string `json:"cpe"`
}

type TypeAheadEntry struct {
	Id    string `json:"_id"`
	Count int    `json:"count"`
}

type TypeAheadList []TypeAheadEntry

type CredentialDocumentList struct {
	Docs  []CredentialDocument `json:"docs"`
	Total int                  `json:"total"`
	Limit int                  `json:"limit"`
	Page  string               `json:"page"`
	Pages int                  `json:"pages"`
}

type CredentialDocument struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Protocol   string   `json:"protocol"`
	IsVerified bool     `json:"isVerified"`
	References []string `json:"references"`
	Cpe        struct {
		Part     string `json:"part"`
		Vendor   string `json:"vendor"`
		Product  string `json:"product"`
		Version  string `json:"version"`
		Language string `json:"language"`
		Update   string `json:"update"`
		Edition  string `json:"edition"`
	} `json:"cpe"`
	Edits []struct {
		ApiKey    string `json:"apiKey"`
		Timestamp int    `json:"timestamp"`
		Edit      struct {
			Username   string   `json:"username"`
			Password   string   `json:"password"`
			Protocol   string   `json:"protocol"`
			References []string `json:"references"`
			Cpe        struct {
				Part     string `json:"part"`
				Vendor   string `json:"vendor"`
				Product  string `json:"product"`
				Version  string `json:"version"`
				Language string `json:"language"`
				Update   string `json:"update"`
				Edition  string `json:"edition"`
			} `json:"cpe"`
		} `json:"edit"`
	} `json:"edits"`
}

func New(apiKey string) Client {

	newClient := Client{
		ApiKey:  apiKey,
		Client:  &http.Client{},
		BaseUrl: BaseUrl,
	}

	return newClient
}

func (c Client) ChangeBaseUrl(url string) {
	c.BaseUrl = url
}

func (c Client) addHeaders() {
	c.Request.Header.Set("X-Api-Key", c.ApiKey)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("User-Agent", "Defaultinator-client/"+Version+";Golang/"+runtime.Version())
}
