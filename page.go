package main

type Page interface {
	ShowPage(*IndexItem)
	SetupPage()
}
