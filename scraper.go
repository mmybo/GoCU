package main
import (
	"fmt"
	"net/http"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)
func main() {
	resp, err := http.Get("https://sfbay.craigslist.org/d/for-sale/search/sss")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
    matcher := func(n *html.Node) bool {
    		if n.DataAtom == atom.A && n.Parent != nil {
    			return scrape.Attr(n.Parent, "class") == "result-info"
    		}
    		return false
    	}
        articles := scrape.FindAll(root, matcher)
        	for i, article := range articles {
        		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
        	}
}