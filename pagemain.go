package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
)

type PageMain struct {
	mainList *tview.List
	currentIndex *[]IndexItem

	listIndexes []int
}

func (page *PageMain) loadPageIndex(file string) {
	page.mainList.Clear()
	LogOut.Printf("load category %s", file)

	// Switch current dir to whatever this is
	dirStack.pushToWorkingDirStack(file)

	currDir := dirStack.getCurrentDirPath()
	LogOut.Printf("current dir is %s", currDir)

	var categoryIndexPath = currDir + "index.json"
	LogOut.Printf("load category file %s", categoryIndexPath)
	index, err := loadIndex(categoryIndexPath)

	if err != nil {
		LogOut.Printf("Could not load page index %s", categoryIndexPath)
		LogOut.Println(err)
		exit()
		return
	}

	page.currentIndex = index

	for _, v := range *index {
		page.mainList.AddItem(v.File, v.Description, 0, nil)
	}
}

func (page *PageMain) loadMainIndex() {
	index, err := loadIndex("index.json")

	if err != nil {
		LogOut.Printf("Could not load main index /assets/index.json")
		LogOut.Println(err)
		exit()
		return
	}

	page.mainList.Clear()
	for _, v := range *index {
		page.mainList.AddItem(v.File, v.Description, 0, nil)
	}
}

func (page *PageMain) hasSelectedItem(index int, mainText string, secondaryText string, shortcut rune) {
	LogOut.Printf("selected %d %s - %s", index, mainText, secondaryText)

	page.listIndexes = append(page.listIndexes, index)

	LogOut.Printf("Pushing index %d, state is %v", index, page.listIndexes)

	selectedItem := (*page.currentIndex)[index]
	switch selectedItem.Type {
	case "directory":
		page.loadPageIndex(selectedItem.File)
	case "file":
		var items []IndexItem
		items = append(items, selectedItem)
		textPage.SetPageData(&items)
		textPage.ShowPage()
	}
}

func (page *PageMain) listInputHandler(key *tcell.EventKey) *tcell.EventKey {
	if(key.Key() == tcell.KeyCtrlLeftSq) { // Got escape, jump back through the stack
		LogOut.Println("Got Escape on list")

		_, err := dirStack.popFromWorkingDirStack()

		// If we are at the end of the stack, the user is hitting
		// Esc on the list page, so they want to quit 
		if err == ErrEndOfStack {
			exit();
			return key
		}

		indexToLoad, err := dirStack.getFinalDirEntry()
		if( err == ErrEndOfStack) {
			page.loadMainIndex()
			return key
		}

		page.loadPageIndex(indexToLoad)
	}

	return key
}

func (page *PageMain) SetPageData(textFiles *[]IndexItem) {
	page.currentIndex = textFiles

	for _, v := range *page.currentIndex {
		page.mainList.AddItem(v.File, v.Description, 0, nil)
	}
}

func (page *PageMain) ShowPage() {
	pages.SwitchToPage(PAGE_MAIN)
}

func (page *PageMain) SetupPage() {
	page.listIndexes = make([]int, 0)

	page.mainList = tview.NewList()
	page.mainList.SetSelectedFunc(page.hasSelectedItem)
	page.mainList.SetInputCapture(page.listInputHandler)

	pages.AddPage(PAGE_MAIN, page.mainList, true, true)
}
