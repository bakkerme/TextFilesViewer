package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"io/ioutil"
)

type PageText struct {
	textView *tview.TextView
	textContent string
}

func (page *PageText) loadTextFile(filePath string) (string) {
	textContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		LogOut.Println(err)
	}

	return string(textContent)
}

func (page *PageText) textViewInputHandler(key *tcell.EventKey) *tcell.EventKey {
	if(key.Key() == tcell.KeyCtrlLeftSq) { // Got escape, jump back through the stack
		mainPage.ShowPage()
	}

	return key
}

func (page *PageText) SetPageData(textFiles *[]IndexItem) {
	textFile := (*textFiles)[0]
	LogOut.Printf("load text file %s", textFile.File)
	page.textContent = page.loadTextFile(getCurrentDirPath() + textFile.File)
	page.textView.SetText(page.textContent)
}

func (page *PageText) ShowPage() {
	pages.SwitchToPage(PAGE_TEXT)
}

func (page *PageText) SetupPage () {
	page.textView = tview.NewTextView()
	page.textView.SetInputCapture(page.textViewInputHandler)

	page.textView.SetText("test text")

	pages.AddPage(PAGE_TEXT, page.textView, true, true)
}
