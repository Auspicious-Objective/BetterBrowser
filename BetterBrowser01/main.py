import webbrowser
import sys

validSites = ["docs.python.org", "reddit.com", "stackoverflow.com", "stackexchange.com", "geeksforgeeks.org", "medium.com", "github.com"]

baseUrl = "https://www.google.com/search?q="

def makeQuery():
    userInput = sys.argv[1:]
    global searchQuery
    searchQuery = "+".join(userInput)

def creatFilter():
    global searchFilter
    searchFilter = "("

    for index, wesite in enumerate(validSites):
        searchFilter += "site:" + wesite
        if index == len(validSites) - 1:
            searchFilter += ")"
        else:
            searchFilter += " OR "
    
def creatUrl(query, filters):
    global search
    search = baseUrl + query + filters
    print(search)

def conductSearch(search):
    webbrowser.open_new(search)

def main():
    makeQuery()
    creatFilter()
    creatUrl(searchQuery, searchFilter)
    conductSearch(search)

main()
