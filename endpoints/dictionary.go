package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) TypeAhead(searchMap map[string]string) TypeAheadList {
	typeAheadList := TypeAheadList{}
	url := c.BaseUrl + "/dictionary/typeahead/?"
	for key, value := range searchMap {
		if key == "field" {
			url += key + "=" + value + "&"
		}
	}
	for key, value := range searchMap {
		for _, field := range QueryFields {
			if key == field {
				if key != "field" {
					url += key + "=" + value + "&"
				}
			}
		}
	}
	c.Request, _ = http.NewRequest("GET", url, nil)
	c.addHeaders()
	resp, _ := c.Client.Do(c.Request)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		println("Response Code: " + resp.Status)
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Message: " + string(body))
		return typeAheadList
	}
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &typeAheadList)
	if err != nil {
		println(err.Error())
		return typeAheadList
	}

	return typeAheadList
}

// Unimplented
func (c Client) NewDictRecord(searchMap map[string]string) {
}

func (c Client) DeleteDictRecord(searchMap map[string]string) {
}
