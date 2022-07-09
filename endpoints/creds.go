package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func (c Client) SearchCredentials(searchMap map[string]string) (CredentialDocumentList, error) {
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
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			return tempCredentialDocumentList, errors.New("Response Code: " + resp.Status + " Message: " + string(body))
		}
		if err != nil || resp.StatusCode != 200 {
			morePages = false
			continue
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(body, &bodyList)
		if err != nil {
			return CredentialDocumentList{}, err
		}
		if page == bodyList.Pages+1 {
			println("No more pages")
			morePages = false
			continue
		}

		tempCredentialDocumentList.Docs = append(tempCredentialDocumentList.Docs, bodyList.Docs...)
		page++
	}
	return tempCredentialDocumentList, nil
}

// Unimplented
func (c Client) SetCredential(id string) {
}

func (c Client) DeleteCredential(id string) {
}
