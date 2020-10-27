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

func (page *PageMain) pageIndexLoad(file string) {
	page.mainList.Clear()
	LogOut.Printf("load category %s", file)

	currDir := getCurrentDirPath()
	var categoryDirPath = currDir + file;
	var categoryIndexPath =  categoryDirPath + "index.json"
	LogOut.Printf("load category file %s", categoryIndexPath)
	var index = loadIndex(categoryIndexPath)

	page.currentIndex = index

	// Switch current dir to whatever this is
	currDir = categoryDirPath
	listDirs = append(listDirs, file)

	LogOut.Printf("Dir stack is %v", listDirs)

	for _, v := range *index {
		page.mainList.AddItem(v.File, v.Description, 0, nil)
	}
}

func (page *PageMain) hasSelectedItem(index int, mainText string, secondaryText string, shortcut rune) {
	LogOut.Printf("selected %d %s - %s", index, mainText, secondaryText)

	page.listIndexes = append(page.listIndexes, index)

	LogOut.Printf("Pusing index %d, state is %v", index, page.listIndexes)

	selectedItem := (*page.currentIndex)[index]
	switch selectedItem.Type {
	case "directory":
		page.pageIndexLoad(selectedItem.File)
	case "file":
		var items []IndexItem
		items = append(items, selectedItem)
		textPage.SetPageData(&items)
		textPage.ShowPage()
	}
}

func (page *PageMain) listInputHandler(key *tcell.EventKey) *tcell.EventKey {
	// LogOut.Printf("key is %d", key.Key())
	if(key.Key() == tcell.KeyCtrlLeftSq) { // Got escape, jump back through the stack
		if(len(listDirs) > 1) {
			LogOut.Println("Got Escape on list")
			LogOut.Printf("Dir stack was %v", listDirs)

			listDirs[len(listDirs)-1] = ""
			listDirs = listDirs[:len(listDirs)-1]

			indexToLoad := listDirs[len(listDirs) - 1]

			// Pop a second time, since we are reloading this item
			listDirs[len(listDirs)-1] = ""
			listDirs = listDirs[:len(listDirs)-1]

			LogOut.Printf("Dir stack is %v", listDirs)

			page.pageIndexLoad(indexToLoad)
		} else {
			exit();
		}
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
	listDirs = append(listDirs, "./assets/")

	page.mainList = tview.NewList()
	page.mainList.SetSelectedFunc(page.hasSelectedItem)
	page.mainList.SetInputCapture(page.listInputHandler)

	pages.AddPage(PAGE_MAIN, page.mainList, true, true)
}
