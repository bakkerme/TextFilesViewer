package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"io/ioutil"
)


var textView *tview.TextView
var text = ""

func loadTextFile(filePath string) (string) {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		LogOut.Println(err)
	}

	return string(text)
}

func pageTextViewLoad(textFile *IndexItem) {
	LogOut.Printf("load text file %s", textFile.File)

	text = loadTextFile(getCurrentDirPath() + textFile.File)
	textView.SetText(text)

	pages.SwitchToPage(PAGE_TEXT)
}

func textViewInputHandler(key *tcell.EventKey) *tcell.EventKey {
	if(key.Key() == tcell.KeyCtrlLeftSq) { // Got escape, jump back through the stack
		pageMainLoad()
	}

	return key
}

func createTextPage() {
	textView = tview.NewTextView()
	textView.SetInputCapture(textViewInputHandler)

	pages.AddPage(PAGE_TEXT, textView, true, true)
}
