package main
import (
    "fmt"
    "net/http"
    "github.com/yhat/scrape"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
    // "fmt"
	// "net/http"
    // "time"
    // "io/ioutil"
    // "log"
    //"encoding/json"
    "github.com/labstack/echo"
    //"github.com/labstack/echo/middleware"
)
type Item struct {
    URL string `json:"URL"`
}
type ItemArray struct {
    items string
    urls string
}
// func crawler(u string) string {
//     resp, err := http.Get(u)
//     if err != nil {
//         panic(err)
//     }
//     root, err := html.Parse(resp.Body)
//     if err != nil {
//         panic(err)
//     }
//     matcher := func(n *html.Node) bool {
//             if n.DataAtom == atom.A && n.Parent != nil {
//                 return scrape.Attr(n.Parent, "class") == "result-info"
//             }
//             return false
//         }
//     articles := scrape.FindAll(root, matcher) //<--change to find, with item matching
//         for i, article := range articles {
//             fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
//             //go deeper, in future
//             // for n := range scrape.Attr(article,"href") {
//             //     fmt.Printf("%2d %s\n", n, scrape.Text(article))
//             // }
//     }
//     return "stringio"
// }
func scraper(c echo.Context) error {
    u := Item{}
    err := c.Bind(&u)
    //u := c.Param("URL")
    fmt.Printf(u.URL)
    // dataType := c.Param("data")
    // if dataType == "json" {
    //     return c.JSON(http.StatusOK, crawler(u))
    // } else {
    //     // return c.JSON(http.StatusBadRequest, map[string]string{
    //     //     "error": "this is wrong",
    //     // })
    // }
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
	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
        sa := []string{scrape.Text(article),scrape.Attr(article, "href")}
        s = append(s, sa)
        //iarray[i] = ItemArray{scrape.Text(article),scrape.Attr(article, "href")}
		//go deeper, in future
        // for n := range scrape.Attr(article,"href") {
        //     fmt.Printf("%2d %s\n", n, scrape.Text(article))
        // }
    }
    //ia := ItemArray{"Stringy","Stringycheese"}
    return c.JSON(http.StatusOK, map[string][][]string{
            "Item List": s,
    })
	//return c.String(http.StatusOK, "Cookie with your meal, Sir?")
}
func main(){
    e := echo.New()
    e.POST("/", scraper)
    e.Start(":5000")
}
