package main

import (
	"defaultinator-cli/endpoints"
	"encoding/json"
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

var tuiVendor string
var tuiProduct string
var tuiVersion string
var tuiUsername string
var tuiPassword string
var tuiPart string
var tuiField string
var queryWindow = true

func StartTui(apiKey string) {
	textUpdate := make(chan string, 1024)
	queryUpdate := make(chan bool, 1024)
	a := tview.NewApplication().EnableMouse(true)

	// Set up the main application window
	mainWindow := tview.NewGrid().SetColumns(0, 0)

	// Build the individual parameter windows
	switchButton := tview.NewButton("|| Switch Query Type ||").SetSelectedFunc(func() {
		queryWindow = !queryWindow
		queryUpdate <- queryWindow
	})
	switchButton.SetBorder(true).SetBackgroundColor(tcell.ColorDarkRed)

	fieldDropdown := tview.NewDropDown().SetLabel("Field (Required): ").SetOptions([]string{"vendor", "product", "version", "username", "password", "part"}, func(text string, index int) {
		tuiField = text
	})

	vendorEntry := tview.NewInputField().SetLabel("Vendor: ").SetFieldWidth(20)
	vendorEntry.SetDoneFunc(func(key tcell.Key) {
		tuiVendor = vendorEntry.GetText()
	})

	productEntry := tview.NewInputField().SetLabel("Product: ").SetFieldWidth(20)
	productEntry.SetDoneFunc(func(key tcell.Key) {
		tuiProduct = productEntry.GetText()
	})

	versionEntry := tview.NewInputField().SetLabel("Version: ").SetFieldWidth(20)
	versionEntry.SetDoneFunc(func(key tcell.Key) {
		tuiVersion = versionEntry.GetText()
	})

	usernameEntry := tview.NewInputField().SetLabel("Username: ").SetFieldWidth(20)
	usernameEntry.SetDoneFunc(func(key tcell.Key) {
		tuiUsername = usernameEntry.GetText()
	})

	passwordEntry := tview.NewInputField().SetLabel("Password: ").SetFieldWidth(20)
	passwordEntry.SetDoneFunc(func(key tcell.Key) {
		tuiPassword = passwordEntry.GetText()
	})

	partEntry := tview.NewInputField().SetLabel("Part: ").SetFieldWidth(20)
	partEntry.SetDoneFunc(func(key tcell.Key) {
		tuiPart = partEntry.GetText()
	})

	paramList := tview.NewTextView()
	paramList.SetBorder(true).SetTitle("Parameters")
	paramList.SetText("")

	typeAheadButton := tview.NewButton("Search").SetSelectedFunc(func() {
		go func() {
			c := endpoints.New(apiKey)
			stringMap := make(map[string]string)
			stringMap["vendor"] = tuiVendor
			stringMap["product"] = tuiProduct
			stringMap["version"] = tuiVersion
			stringMap["username"] = tuiUsername
			stringMap["password"] = tuiPassword
			stringMap["part"] = tuiPart
			stringMap["field"] = tuiField
			list, err := c.TypeAhead(stringMap)
			if err != nil {
				textUpdate <- err.Error()
				return
			}
			returnData, _ := json.MarshalIndent(list, "", "    ")
			textUpdate <- string(returnData)
		}()
	})
	typeAheadButton.SetBorder(true).SetBackgroundColor(tcell.ColorDarkRed)

	searchButton := tview.NewButton("Search").SetSelectedFunc(func() {
		go func() {
			c := endpoints.New(apiKey)
			stringMap := make(map[string]string)
			stringMap["vendor"] = tuiVendor
			stringMap["product"] = tuiProduct
			stringMap["version"] = tuiVersion
			stringMap["username"] = tuiUsername
			stringMap["password"] = tuiPassword
			stringMap["part"] = tuiPart
			stringMap["field"] = tuiField
			list, err := c.SearchCredentials(stringMap)
			if err != nil {
				textUpdate <- err.Error()
				return
			}
			returnData, _ := json.MarshalIndent(list, "", "    ")
			textUpdate <- string(returnData)
		}()
	})
	searchButton.SetBorder(true).SetBackgroundColor(tcell.ColorDarkRed)

	// Construct typeahead window
	typeAheadWindow := tview.NewGrid().SetColumns(0, 0, 0).SetRows(-2, -2, -1)
	typeAheadWindow.AddItem(fieldDropdown, 1, 2, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(vendorEntry, 0, 0, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(productEntry, 0, 1, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(versionEntry, 0, 2, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(usernameEntry, 0, 3, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(passwordEntry, 1, 0, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(passwordEntry, 1, 0, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(paramList, 1, 3, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(partEntry, 1, 1, 1, 1, 0, 0, false)
	typeAheadWindow.AddItem(typeAheadButton, 2, 0, 1, 4, 0, 0, false)

	// Construct search window
	searchWindow := tview.NewGrid().SetColumns(0, 0, 0).SetRows(-2, -2, -1)
	searchWindow.AddItem(vendorEntry, 0, 0, 1, 1, 0, 0, false)
	searchWindow.AddItem(productEntry, 0, 1, 1, 1, 0, 0, false)
	searchWindow.AddItem(versionEntry, 0, 2, 1, 1, 0, 0, false)
	searchWindow.AddItem(usernameEntry, 0, 3, 1, 1, 0, 0, false)
	searchWindow.AddItem(passwordEntry, 1, 0, 1, 1, 0, 0, false)
	searchWindow.AddItem(partEntry, 1, 1, 1, 1, 0, 0, false)
	searchWindow.AddItem(tview.NewBox(), 1, 2, 1, 1, 0, 0, false)
	searchWindow.AddItem(paramList, 1, 3, 1, 1, 0, 0, false)
	searchWindow.AddItem(searchButton, 2, 0, 1, 4, 0, 0, false)

	// Left Pane - Search Parameters
	leftPane := tview.NewGrid().SetRows(3, 1, 0)
	leftPane.SetBorder(true).SetTitle("Query Options")
	leftPane.AddItem(switchButton, 0, 0, 1, 1, 0, 0, false)
	leftPane.AddItem(typeAheadWindow, 2, 0, 1, 1, 0, 0, false)

	// Right Pane - Text display
	rightSide := tview.NewGrid().SetRows(0, 3)

	rightPane := tview.NewTextView().SetText("Defaultinator-CLI")
	rightPane.SetBorder(true).SetTitle("Results")

	copyButton := tview.NewButton("Copy to Clipboard").SetSelectedFunc(func() {
		clipboard.WriteAll(rightPane.GetText(true))
	})
	copyButton.SetBorder(true).SetBackgroundColor(tcell.ColorDarkRed)

	rightSide.AddItem(rightPane, 0, 0, 1, 1, 0, 0, false)
	rightSide.AddItem(copyButton, 1, 0, 1, 1, 0, 0, false)

	mainWindow.AddItem(leftPane, 0, 0, 1, 1, 0, 0, false)
	mainWindow.AddItem(rightSide, 0, 1, 1, 1, 0, 0, false)

	// Start the goroutine that updates the panes.
	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			select {
			case text := <-textUpdate:
				a.QueueUpdateDraw(
					func() {
						rightPane.SetText(text)
					})
			case query := <-queryUpdate:

				if query {
					a.QueueUpdateDraw(func() {
						leftPane.Clear()
						leftPane.SetTitle("TYPE-AHEAD")
						leftPane.AddItem(switchButton, 0, 0, 1, 1, 0, 0, false)
						leftPane.AddItem(typeAheadWindow, 2, 0, 1, 1, 0, 0, false)
						textUpdate <- "Type ahead query"
					})
				} else {
					a.QueueUpdateDraw(func() {
						leftPane.Clear()
						leftPane.SetTitle("SEARCH")
						leftPane.AddItem(switchButton, 0, 0, 1, 1, 0, 0, false)
						leftPane.AddItem(searchWindow, 2, 0, 1, 1, 0, 0, false)
						textUpdate <- "Search query"
					})
				}

			default:
				a.QueueUpdateDraw(func() {
					paramString := "vendor: " + tuiVendor + "\n" + "product: " + tuiProduct + "\n" + "version: " + tuiVersion + "\n" + "username: " + tuiUsername + "\n" + "password: " + tuiPassword + "\n" + "part: " + tuiPart + "\n" + "field: " + tuiField + "\n"
					paramList.SetText(paramString)
				})
				continue
			}

		}
	}()

	a.SetRoot(mainWindow, true).SetFocus(rightPane).Run()
}
