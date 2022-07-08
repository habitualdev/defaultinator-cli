package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c Client) SearchCredentials(searchMap map[string]string) CredentialDocumentList {
	tempCredentialDocumentList := CredentialDocumentList{}
	morePages := true
	url := c.BaseUrl + "/credentials/search?"
	query := ""
	for key, value := range searchMap {
		for _, field := range QueryFields {
			if key == field {
				query += key + "=" + value + "&"
			}
		}
	}
	page := 1
	for morePages {
		bodyList := CredentialDocumentList{}
		requrl := url + "page=" + strconv.Itoa(page) + "&" + query
		c.Request, _ = http.NewRequest("GET", requrl, nil)
		c.addHeaders()
		resp, err := c.Client.Do(c.Request)
		if page == 1 && resp.StatusCode != 200 {
			println(resp.Status)
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
			return tempCredentialDocumentList
		}
		if err != nil || resp.StatusCode != 200 {
			morePages = false
			continue
		}
		println("Page: " + strconv.Itoa(page))
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(body, &bodyList)
		if err != nil {
			println(err.Error())
		}
		if page == bodyList.Pages+1 {
			println("No more pages")
			morePages = false
			continue
		}

		tempCredentialDocumentList.Docs = append(tempCredentialDocumentList.Docs, bodyList.Docs...)
		page++
	}
	return tempCredentialDocumentList
}

// Unimplented
func (c Client) SetCredential(id string) {
}

func (c Client) DeleteCredential(id string) {
}
