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
type Item struct {
    Item string `json:"item"`
    URL string `json:"URL"`
}
func scraper(c echo.Context) error {
    u := Item{}
    err := c.Bind(&u)
    fmt.Printf(u.Item)
    fmt.Printf(u.URL)
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
    articles := scrape.FindAll(root, matcher)
    var s [][]string
	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
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
func main(){
    e := echo.New()
    e.POST("/", scraper)
    e.Start(":5000")
}
