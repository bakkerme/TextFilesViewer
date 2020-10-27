package main

type Page interface {
	ShowPage()
	SetupPage()
	SetPageData(*[]IndexItem)
}
