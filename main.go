package main
import (
    "fmt"
    "net/http"
    "github.com/yhat/scrape"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
    "strings"
    "github.com/labstack/echo"
)
//create struct that will be used to get javascript objects
type Item struct {
    Item string `json:"item"`
    URL string `json:"URL"`
}
//create scrapet function, passing in context
func scraper(c echo.Context) error {
    //set u to Item struct
    u := Item{}
    //bind u struct with echo context
    err := c.Bind(&u)
    fmt.Printf(u.Item)
    fmt.Printf(u.URL)
	resp, error := http.Get(u.URL)
	//check for errors when getting, if so... panic!
	if error != nil {
		panic(err)
	}
	root, error := html.Parse(resp.Body)
	//do the same thing again, except for parsing. Don't Panic... unless you have to
	if error != nil {
		panic(err)
	}
    //utilizing awesome html/atom to scrape attributes of site we're searching
    matcher := func(n *html.Node) bool {
	    	//check to make sure attribute is node data, and that it has a parent
    		if n.DataAtom == atom.A && n.Parent != nil {
			//scrape it
    			return scrape.Attr(n.Parent, "class") == "result-info"
    		}
    		return false
    }
    //get all the pages from url
    articles := scrape.FindAll(root, matcher)
    //create string array for returning JSON
    var s [][]string
	//iterate through index and articles (webpage routes)
	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	//add spaces between items, to make sure it's a word and not part of a word.
	//Example: If we search "read", we want articles with "read", not "thread"
        sps := " " + u.Item + " "
	//check if article title contains item
        if strings.Contains(scrape.Text(article),sps) {
	    //if matching, create array of strings, containing title of webpage and URL
            sa := []string{scrape.Text(article),scrape.Attr(article, "href")}
	    //append it to JSON array
            s = append(s, sa)
        }
    }
     //use echo context to return javascript object dictionary, containing array of items returned
    return c.JSON(http.StatusOK, map[string][][]string{
            "Item List": s,
    })
}
func main(){
    //create new echo
    e := echo.New()
    //call scraper function on post request
    e.POST("/search", scraper)
    //run locally
    e.Start(":5000")
}
