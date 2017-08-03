package main

import (
	"golang-web-testcase/router"
)

func main() {
	router := router.Load()
	router.LoadHTMLGlob("templates/*")
	router.Run()
}
