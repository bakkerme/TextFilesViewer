package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"io/ioutil"
)

type PageText struct {
	textView *tview.TextView
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
		pageMainLoad()
	}

	return key
}

func (page *PageText) ShowPage(textFile *IndexItem) {
	LogOut.Printf("load text file %s", textFile.File)

	textContent := page.loadTextFile(getCurrentDirPath() + textFile.File)

	// LogOut.Printf("text content %s", textContent)
	LogOut.Printf("text view %p", &page.textView)

	page.textView.SetText(textContent)

	pages.SwitchToPage(PAGE_TEXT)
}

func (page *PageText) SetupPage () {
	page.textView = tview.NewTextView()
	page.textView.SetInputCapture(page.textViewInputHandler)

	page.textView.SetText("test text")

	pages.AddPage(PAGE_TEXT, page.textView, true, true)
}
