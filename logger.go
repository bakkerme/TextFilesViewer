package main

import (
	"os"
	"log"
)

var LogOut *log.Logger

func InitLogger() {
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	LogOut = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
