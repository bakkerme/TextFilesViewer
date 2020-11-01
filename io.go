package main

import (
	"io/ioutil"
	"encoding/json"
)

var assetDir string = "./assets/"

func loadTextFile(filePath string) (string) {
	textContent, err := ioutil.ReadFile(assetDir + filePath)
	if err != nil {
		LogOut.Println(err)
	}

	return string(textContent)
}

func loadIndex(filePath string) (*[]IndexItem, error) {
	indexJSON, err := ioutil.ReadFile(assetDir + filePath)
	if err != nil {
		return nil, err
	}

	var index []IndexItem
	marshalErr := json.Unmarshal([]byte(indexJSON), &index)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return &index, nil
}
