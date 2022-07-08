package main

import (
	"defaultinator-cli/endpoints"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var tui bool
var field string
var searchType string
var vendor string
var product string
var username string
var password string
var part string
var apiKey string
var output string
var baseUrl string

const Splash = `
    ____       ____            ____  _             __                   ________    ____
   / __ \___  / __/___ ___  __/ / /_(_)___  ____ _/ /_____  _____      / ____/ /   /  _/
  / / / / _ \/ /_/ __  / / / / / __/ / __ \/ __  / __/ __ \/ ___/_____/ /   / /    / /  
 / /_/ /  __/ __/ /_/ / /_/ / / /_/ / / / / /_/ / /_/ /_/ / /  /_____/ /___/ /____/ /   
/_____/\___/_/  \__,_/\__,_/_/\__/_/_/ /_/\__,_/\__/\____/_/         \____/_____/___/ 

by habitual`

func init() {
	flag.BoolVar(&tui, "tui", false, "launch the tui")
	flag.StringVar(&field, "field", "", "search field")
	flag.StringVar(&searchType, "type", "Search", "Choose Typeahead or Search")
	flag.StringVar(&vendor, "vendor", "", "search vendor")
	flag.StringVar(&product, "product", "", "search product")
	flag.StringVar(&username, "username", "", "search username")
	flag.StringVar(&password, "password", "", "search password")
	flag.StringVar(&part, "part", "", "search part")
	flag.StringVar(&apiKey, "apiKey", "", "api key")
	flag.StringVar(&output, "output", "", "set output file")
	flag.StringVar(&baseUrl, "baseUrl", "", "Change to use a different base url")

}

func main() {
	flag.Parse()
	fmt.Println(Splash)

	if apiKey == "" {
		fmt.Println("Please provide an api key")
		return
	}
	if !tui {
		if searchType == "Search" {
			c := endpoints.New(apiKey)
			if baseUrl != "" {
				c.ChangeBaseUrl(baseUrl)
			}
			search := map[string]string{}
			if vendor != "" {
				search["vendor"] = vendor
			}
			if product != "" {
				search["product"] = product
			}
			if username != "" {
				search["username"] = username
			}
			if password != "" {
				search["password"] = password
			}
			if part != "" {
				search["part"] = part
			}
			list := c.SearchCredentials(search)
			data, err := json.Marshal(list.Docs)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				if output != "" {
					fmt.Println("Writing to file")
					err := os.WriteFile(output, data, 0644)
					if err != nil {
						fmt.Println(err.Error())
					}
				} else {
					fmt.Println(string(data))
				}
			}
		} else if searchType == "Typeahead" {
			c := endpoints.New(apiKey)
			if baseUrl != "" {
				c.ChangeBaseUrl(baseUrl)
			}
			search := map[string]string{}
			if field == "" {
				fmt.Println("Please specify a field to search")
				return
			} else {
				search["field"] = field
			}
			if vendor != "" {
				search["vendor"] = vendor
			}
			if product != "" {
				search["product"] = product
			}
			if part != "" {
				search["part"] = part
			}

			typeahead := c.TypeAhead(search)
			data, err := json.Marshal(typeahead)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				if output != "" {
					fmt.Println("Writing to file")
					err := os.WriteFile(output, data, 0644)
					if err != nil {
						fmt.Println(err.Error())
					}
				} else {
					fmt.Println(string(data))
				}
			}
		}
	} else {
		fmt.Println("TUI not implemented yet")
	}
}
