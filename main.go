package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	"github.com/labstack/echo"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
	_ "github.com/heroku/x/hmetrics/onload"
)
type Item struct {
    Item string `json:"item"`
    URL string `json:"URL"`
}
func scraper(c echo.Context) error {
	b := bluemonday.UGCPolicy()
    u := Item{}
    err := c.Bind(&u)
	u.Item, u.URL = b.Sanitize(u.Item), b.Sanitize(u.URL)
	resp, error := http.Get(u.URL)
	if error != nil {
		panic(err)
	}
	root, error := html.Parse(resp.Body)
	if error != nil {
		panic(err)
	}
    matcher := func(n *html.Node) bool {
    		if n.DataAtom == atom.A && n.Parent != nil {
    			return scrape.Attr(n.Parent, "class") == "result-info"
    		}
    		return false
    }
    articles := scrape.FindAll(root, matcher) //<--change to find, with item matching
    //iarray := [10]ItemArray
    var s [][]string
	for _, article := range articles {
        sps := " " + u.Item + " "
        if strings.Contains(scrape.Text(article),sps) {
            sa := []string{scrape.Text(article),scrape.Attr(article, "href")}
            s = append(s, sa)
        }
    }
    return c.JSON(http.StatusOK, map[string][][]string{
            "Item List": s,
    })
}
func main() {
	b := bluemonday.UGCPolicy()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	e := echo.New()
    e.POST("/searchy", scraper)
	router.POST("/search", func(c *gin.Context) {
		u := Item{}
		err := c.Bind(&u)
		u.Item, u.URL = b.Sanitize(u.Item), b.Sanitize(u.URL)
		fmt.Printf(u.Item)
		resp, error := http.Get(u.URL)
		if error != nil {
			panic(err)
		}
		root, error := html.Parse(resp.Body)
		if error != nil {
			panic(err)
		}
		matcher := func(n *html.Node) bool {
				if n.DataAtom == atom.A && n.Parent != nil {
					return scrape.Attr(n.Parent, "class") == "result-info"
				}
				return false
		}
		articles := scrape.FindAll(root, matcher) //<--change to find, with item matching
		var s [][]string
		for i, article := range articles {
			fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
			sps := strings.ToLower(u.Item)
			ls := strings.ToLower(scrape.Text(article))
			if strings.Contains(ls,sps) {
				sa := []string{scrape.Text(article),scrape.Attr(article, "href")}
				s = append(s, sa)
			}
		}
		c.JSON(http.StatusOK, map[string][][]string{
	            "Item List": s,
	    })
	})
	router.GET("/mark", func(c *gin.Context) {
  		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
	})
	//e.Logger.Fatal(e.Start(":" + port))
	router.Run(":" + port)
	//e.Start(":" + port)
}
