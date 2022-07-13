package main

import (
	"defaultinator-cli/endpoints"
	"defaultinator-cli/utils"
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
	flag.StringVar(&searchType, "type", "", "Choose Typeahead or Search")
	flag.StringVar(&vendor, "vendor", "", "search vendor")
	flag.StringVar(&product, "product", "", "search product")
	flag.StringVar(&username, "username", "", "search username")
	flag.StringVar(&password, "password", "", "search password")
	flag.StringVar(&part, "part", "", "search part")
	flag.StringVar(&apiKey, "apiKey", "", "api key")
	flag.StringVar(&output, "output", "", "set output file")
	flag.StringVar(&baseUrl, "baseUrl", "", "Change to use a different base url")

}

var Usage = func() {
	fmt.Println(Splash)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if apiKey == "" {
		var err error
		apiKey, err = utils.GetCreds()
		if tui && err != nil {
			fmt.Println("You must provide an API key")
			fmt.Print("API Key: ")
			fmt.Scanln(&apiKey)

		} else if err != nil {
			fmt.Println("Please provide an api key")
			return
		} else {
			fmt.Println("API Key Retrieved")
		}
	} else {
		err := utils.SaveCreds(apiKey)
		if err != nil {
			fmt.Println("Error saving api key")
			fmt.Println(err.Error())
			return
		}
	}
	c := endpoints.New(apiKey)
	err := c.CheckKey()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if !tui {
		if searchType == "" {
			fmt.Println("Please provide a search type")
			return
		}
		if searchType == "Search" {
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
			list, err := c.SearchCredentials(search)
			if err != nil {
				fmt.Println(err)
				return
			}
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

			typeahead, err := c.TypeAhead(search)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
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
		fmt.Println("Starting TUI")
		StartTui(apiKey)
	}
}
