package main

import (
	"fmt"
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

func exit() {
	app.Stop()
}

var dirStack DirStack

var textPage Page
var mainPage Page
func main() {
	InitLogger();

	tview.Styles.PrimaryTextColor = tcell.ColorGreen
	tview.Styles.TertiaryTextColor = tcell.ColorWhite

	dirStack = DirStack{}

	app = tview.NewApplication()
	pages = tview.NewPages()

	textPage = &PageText{}
	textPage.SetupPage()

	currentIndex, err := loadIndex("index.json")
	if err != nil {
		fmt.Println("Could not load base page index, './assets/index.json'")
		fmt.Println(err)

		fmt.Println("Could not load base page index, './assets/index.json'")
		LogOut.Println(err)
		exit()
	}

	mainPage = &PageMain{}
	mainPage.SetupPage()
	mainPage.SetPageData(currentIndex)
	mainPage.ShowPage()

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
