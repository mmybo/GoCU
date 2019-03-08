[![Go Report Card](https://goreportcard.com/badge/github.com/mmybo/GoCU)](https://goreportcard.com/report/github.com/mmybo/GoCU)

[![GoDoc](https://godoc.org/github.com/mmybo/GoCU?status.svg)](https://godoc.org/github.com/mmybo/GoCU)

A Go Crawler Utility API (GoCU)

Simply make a post request using JSON, with the following format:

     {
         "Item": [word you are searching for],
         "URL": [web address you wish to search for it in]
     }
 
And GoCU will return all pages containing that word, and the links to them:

    {
        "Item List": [
            "[Page Title]",
            "[URL Linking to Page]"
        ]
    }

Docs page can be found here: https://mmybo.github.io/GoCU/
Deployed: https://gocu.herokuapp.com/

GoCU is open source, so you can clone it and make it even more powerful!

To start off, simply clone the repo, using 'git clone https://github.com/mmybo/GoCU.git' into your terminal.

Next install the dependencies, using the following commands:
go get github.com/yhat/scrape
go get golang.org/x/net/html
go get golang.org/x/net/html/atom
go get github.com/labstack/echo

And that's it! After that, just build and run! (go build, and go run main.go)

Run it locally, or deploy your own!
