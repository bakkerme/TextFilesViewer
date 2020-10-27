package main

import (
	"io/ioutil"
	"encoding/json"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
)

type IndexItem struct {
	File string `json:"file"`
	Description string `json:"description"`
	Type string `json:"type"`
	Size int `json:"size,omitempty"`
}

const PAGE_MAIN = "MAIN"
const PAGE_TEXT= "TEXT"

var app *tview.Application
var pages *tview.Pages


var listDirs []string
func getCurrentDirPath () (string) {
	tempPath := ""
	for _, value := range listDirs {
		tempPath += value
	}

	return tempPath
}

func loadIndex(filePath string) (*[]IndexItem) {
	indexJSON, err := ioutil.ReadFile(filePath)
	if err != nil {
		LogOut.Println(err)
	}

	var index []IndexItem
	marshalErr := json.Unmarshal([]byte(indexJSON), &index)
	if marshalErr != nil {
		LogOut.Println(marshalErr)
	}

	return &index
}

func exit() {
	app.Stop()
}

var textPage Page
var mainPage Page
func main() {
	InitLogger();

	tview.Styles.PrimaryTextColor = tcell.ColorGreen
	tview.Styles.TertiaryTextColor = tcell.ColorWhite

	app = tview.NewApplication()
	pages = tview.NewPages()

	textPage = &PageText{}
	textPage.SetupPage()

	currentIndex := loadIndex("./assets/index.json")
	mainPage = &PageMain{}
	mainPage.SetupPage()
	mainPage.SetPageData(currentIndex)
	mainPage.ShowPage()

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
