# Welcome to the Gerrymandering API!

> A Go Crawler Utility API (GoCU) to search websites for matching items

# Getting started with API calls
>Simply make a JSON Post request to 'https://gocu.herokuapp.com/search', with the following format:
>
    {
    "Item": "Name of item you are searching for",
    "URL": "Website you are searching for it in"
    }
>And GoCU will return all pages containing that word, and the links to them:
>
    {
    "Item Names": [
    "Array of page titles associated with word",
    "URL of web page it's from"
        ]
    }
>It's that simple!

# Cloning the API
>GoCU is open source, so you can clone it and make it even more powerful!

>To start off, simply clone the repo, using 'git clone https://github.com/mmybo/GoCU.git' into your terminal

>But you probably already know that!

>Next install the dependencies:

>go get github.com/yhat/scrape

>go get golang.org/x/net/html

>go get golang.org/x/net/html/atom

>go get github.com/labstack/echo

>And that's it! After that, just build and run! (go build, and go run main.go)

>Run it locally, or deploy your own!

# More info

>GoCU is deployed here: https://gocu.herokuapp.com/
>Github repo: https://github.com/mmybo/GoCU
>My LinkedIn: https://www.linkedin.com/mynetwork/
>My Portfolio: https://www.makeschool.com/portfolio/Jaeson-Booker
