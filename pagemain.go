package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
)

var categoriesGlobal *[]IndexItem

var mainList *tview.List
var currentIndex *[]IndexItem

var listIndexes []int
var listDirs []string

func pageIndexLoad(file string) {
	mainList.Clear()
	LogOut.Printf("load category %s", file)

	currDir := getCurrentDirPath()
	var categoryDirPath = currDir + file;
	var categoryIndexPath =  categoryDirPath + "index.json"
	LogOut.Printf("load category file %s", categoryIndexPath)
	var index = loadIndex(categoryIndexPath)

	currentIndex = index

	// Switch current dir to whatever this is
	currDir = categoryDirPath
	listDirs = append(listDirs, file)

	LogOut.Printf("Dir stack is %v", listDirs)

	for _, v := range *index {
		mainList.AddItem(v.File, v.Description, 0, nil)
	}
}

func hasSelectedItem(index int, mainText string, secondaryText string, shortcut rune) {
	LogOut.Printf("selected %d %s - %s", index, mainText, secondaryText)

	listIndexes = append(listIndexes, index)

	LogOut.Printf("Pusing index %d, state is %v", index, listIndexes)

	selectedItem := (*currentIndex)[index]
	switch selectedItem.Type {
	case "directory":
		pageIndexLoad(selectedItem.File)
	case "file":
		pageTextViewLoad(&selectedItem)
	}
}

func listInputHandler(key *tcell.EventKey) *tcell.EventKey {
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

			pageIndexLoad(indexToLoad)
		} else {
			exit();
		}
	}

	return key
}

func pageMainLoad() {
	pages.SwitchToPage(PAGE_MAIN)
}

func createMainPage(categories *[]IndexItem) {
	currentIndex = categories

	listIndexes = make([]int, 0)
	listDirs = append(listDirs, "./assets/")

	mainList = tview.NewList()
	mainList.SetSelectedFunc(hasSelectedItem)
	mainList.SetInputCapture(listInputHandler)

	categoriesGlobal = categories

	for _, v := range *categories {
		mainList.AddItem(v.File, v.Description, 0, nil)
	}

	pages.AddPage(PAGE_MAIN, mainList, true, true)
}
