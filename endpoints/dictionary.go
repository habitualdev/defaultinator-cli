package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c Client) TypeAhead(searchMap map[string]string) (TypeAheadList, error) {
	typeAheadList := TypeAheadList{}
	url := c.BaseUrl + "/dictionary/typeahead/?"
	for key, value := range searchMap {
		if key == "field" {
			url += key + "=" + value + "&"
		}
	}
	for key, value := range searchMap {
		if value == "" {
			continue
		}
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
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return typeAheadList, errors.New("Response Code: " + resp.Status + " Message: " + string(body))
	}
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &typeAheadList)
	if err != nil {
		return typeAheadList, err
	}

	return typeAheadList, nil
}

// Unimplented
func (c Client) NewDictRecord(searchMap map[string]string) {
}

func (c Client) DeleteDictRecord(searchMap map[string]string) {
}
