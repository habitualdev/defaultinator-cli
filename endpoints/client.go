package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"runtime"
)

const Version = "1.2.2"
const BaseUrl = "https://api.defaultinator.com"

var QueryFields = []string{"vendor", "product", "version", "username", "password", "part", "field"}

type Client struct {
	ApiKey     string
	ApiKeyInfo *ApiKeyInfo
	Client     *http.Client
	Request    *http.Request
	BaseUrl    string
}

type ApiKeyInfo struct {
	Id        string `json:"_id"`
	ApiKey    string `json:"apiKey"`
	Email     string `json:"email"`
	Notes     string `json:"notes"`
	IsAdmin   bool   `json:"isAdmin"`
	IsRootKey bool   `json:"isRootKey"`
	V         int    `json:"__v"`
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
		ApiKey:     apiKey,
		Client:     &http.Client{},
		ApiKeyInfo: &ApiKeyInfo{},
		Request:    &http.Request{},
		BaseUrl:    BaseUrl,
	}

	return newClient
}

func (c *Client) ChangeBaseUrl(url string) {
	c.BaseUrl = url
}

func (c *Client) addHeaders() {
	c.Request.Header.Set("X-Api-Key", c.ApiKey)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("User-Agent", "Defaultinator-client/"+Version+";Golang/"+runtime.Version())
}

func (c *Client) CheckKey() error {
	c.Request, _ = http.NewRequest("GET", c.BaseUrl+"/apikeys/keyinfo", nil)
	c.addHeaders()
	resp, err := c.Client.Do(c.Request)
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	tempKeyInfo := ApiKeyInfo{}
	err = json.Unmarshal(body, &tempKeyInfo)
	if err != nil {
		println(err.Error())
		return err
	}
	c.ApiKeyInfo = &tempKeyInfo

	return nil
}
