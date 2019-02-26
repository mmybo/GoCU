package main
//
// // importing dependencies
// import (
//     "encoding/json"
//     "log"
//     "os"
//     "fmt"
//     "strings"
//     "github.com/gocolly/colly"
//     "github.com/fatih/color"
//     "database/sql"
//     _ "github.com/lib/pq"
//     "strconv"
// )
//
// // setting up a datastruture to represent a form of Clothing
// type Clothing struct {
//     Name                    string
//     Code                    string
//     Description     string
//     Price                   float64
// }
//
// // setting up a function to write to our db
// func dbWrite(product Clothing) {
//     const (
//       host     = "localhost"
//       port     = 5432
//       user     = "user"
//       // password = ""
//       dbname   = "lucas_db"
//     )
//
//     psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//     "dbname=%s sslmode=disable",
//     host, port, user, dbname)
//
//     db, err := sql.Open("postgres", psqlInfo)
//   if err != nil {
//     panic(err)
//   }
//   defer db.Close()
//
//   err = db.Ping()
//   if err != nil {
//     panic(err)
//   }
//
//   // some debug print logs
//   log.Print("Successfully connected!")
//     fmt.Printf("%s, %s, %s, %f", product.Name, product.Code, product.Description, product.Price)
//     sqlStatement := `
//     INSERT INTO floryday (product, code, description, price)
//     VALUES ($1, $2, $3, $4)`
//     _, err = db.Exec(sqlStatement, product.Name, product.Code, product.Description, product.Price)
//     if err != nil {
//       panic(err)
//     }
// }
//
// // our main function - using a colly collector
// func main() {
//         // creating our new colly collector with a localised cache
//     c := colly.NewCollector(
//         // colly.AllowedDomains("https://www.clotheswebsite.com/"),
//         colly.CacheDir(".floryday_cache"),
//     // colly.MaxDepth(5), // keeping crawling limited for our initial experiments
//   )
//
//     // clothing detail scraping collector
//     detailCollector := c.Clone()
//
//         // setting our array of clothing to size 200
//     clothes := make([]Clothing, 0, 200)
//
//     // Find and visit all links
//     c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//
//         link := e.Attr("href")
//
//         // hardcoded urls to skip -> these arent relevant for products
//         if !strings.HasPrefix(link, "/?country_code") || strings.Index(link, "/cart.php") > -1 ||
//         strings.Index(link, "/login.php") > -1 || strings.Index(link, "/cart.php") > -1 ||
//         strings.Index(link, "/account") > -1 || strings.Index(link, "/privacy-policy.html") > -1 {
//             return
//         }
//
//         // scrape the page
//         e.Request.Visit(link)
//     })
//
//     // printing visiting message for debug purposes
//     c.OnRequest(func(r *colly.Request) {
//         log.Println("Visiting", r.URL.String(), "\n")
//     })
//
//     // visting any href links -> this can be optimised later
//     c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
//
//         clothingURL := e.Request.AbsoluteURL(e.Attr("href"))
//
//         // this was a way to determine the page was definitely a product
//                 // if it contained -Dress- we were good to scrape
//         if strings.Contains(clothingURL, "-Dress-"){
//             // Activate detailCollector
//             color.Green("Crawling Link Validated -> Commencing Crawl for %s", clothingURL)
//             detailCollector.Visit(clothingURL)
//         } else {
//             color.Red("Validation Failed -> Cancelling Crawl for %s", clothingURL)
//             return
//         }
//
//     })
//
//     // Extract details of the clothing
//     detailCollector.OnHTML(`div[class=prod-right-in]`, func(e *colly.HTMLElement) {
//         // some html parsing to get the exact values we want
//         title := e.ChildText(".prod-name")
//         code := strings.Split(e.ChildText(".prod-item-code"), "#")[1]
//         stringPrice := strings.TrimPrefix(e.ChildText(".prod-price"),"€ ")
//         price, err := strconv.ParseFloat(stringPrice, 64) // conversion to float64
//         color.Red("err in parsing price -> %s", err)
//         description := e.ChildText(".grid-uniform")
//
//         clothing := Clothing{
//             Name:                   title,
//             Code:                   code,
//             Description:        description,
//             Price:                  price,
//         }
//
//         // writing as we go to DB
//         // TODO optiize to handle bulk array uploads instead of one at a time
//         dbWrite(clothing)
//
//         // appending to our output array...
//         clothes = append(clothes, clothing)
//     })
//
//     // start scraping at our seed address
//     c.Visit("https://ebay.com/")
//
//     enc := json.NewEncoder(os.Stdout)
//     enc.SetIndent("", "  ")
//
//     // Dump json to the standard output
//     enc.Encode(clothes)
//
// }
