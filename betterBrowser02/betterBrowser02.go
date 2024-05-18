package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var validSites = [10]string{"docs.python.org", "reddit.com", "stackoverflow.com", "stackexchange.com", "geeksforgeeks.org", "medium.com", "github.com", "w3schools.com", "go.dev", "rust-lang.org"}
var baseUrl string = "https://www.google.com/search?q="

func makeQuery() string {
	arg := os.Args[1:]
	var argsString string = strings.Join(arg, " ")
	var searchQuery string = baseUrl + argsString
	return searchQuery
}

func createFilter(siteList []string) string {
	var searchFilter string = "("
	for i := 0; i < len(siteList); i++ {
		searchFilter += "site:" + siteList[i]
		if i == (len(siteList) - 1) {
			searchFilter += ")"
		} else {
			searchFilter += "OR"
		}
	}
	return searchFilter
}

func createURL(base string, query string, filter string) string {
	var search string = base + query + filter
	return search
}

func makeSearch(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	}

	if err != nil {
		fmt.Println("You have encountered an error. Your OS might not be supported.")
	}
}

func main() {
	makeSearch(createURL(baseUrl, makeQuery(), createFilter(validSites[:])))
}
