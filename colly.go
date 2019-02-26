package main
//  import (
//      "fmt"
//      "strings"
//      "time"
//
//      "github.com/PuerkitoBio/goquery"
//      "github.com/gocolly/colly"
//  )
// // func main(){
// //     c := colly.NewCollector(
// //         //limit crawing to craigslist
// //         colly.AllowedDomains("craigslist.com"),
// //         //allow repeated visits to site
// //         colly.AllowURLRevisit(),
// //         //do parrel crawling, GoCU Style!
// //         colly.Async(true),
// //     )
// //     c.Limit(&colly.LimitRule{
// //         //filter websites according to limitations
// //         DomainGlob: ()"craigslist.com"),
// //         //Delay for requests, so as not to get in trouble with the Domain Lords
// //         Delay: 1 * time.Second,
// //         //delay at random, so cyber Frieza doesn't try to take down GoCU
// //         //in other words, so our crawler doesn't get removed for making too many patterned requests
// //         RandomDelay: 1 * time.Second,
// //     })
// //     c.Visit("https://craigslist.com")
// //     c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// //         //get the link from anchor HTMLElement
// //         link := e.Attr("href")
// //         //get Colly collector to go to link
// //         c.Visit(e.Request.AbsoluteURL(link))
// //     })
// //     numVisited := 0
// //     c.OnRequest(func(r *colly.Request) {
// //         if numVisited > 100 {
// //             r.Abort()
// //         }
// //         numVisited++
// //     })
// //     c.OnResponse(func(r *colly.Response) {
// //         fmt.Println(r.Body)
// //     })
// //     c.OnHTML("#items", func(e *colly.HTMLElement) {
// //         fmt.Println(e.ChildText("item"))
// //     })
// //     c.OnHTML("#items", func(e *colly.HTMLElement) {
// //         e.ForEach("item", func(_ int, elem *colly.HTMLElement) {
// //             if strings.Contains(elem.Text, "item") {
// //                 fmt.Println(elem.Text)
// //             }
// //         })
// //     })
// // }
// // c.Limit(&colly.LimitRule{
// //     //filter websites according to limitations
// //     DomainGlob: ()"craigslist.com"),
// //     //Delay for requests, so as not to get in trouble with the Domain Lords
// //     Delay: 1 * time.Second,
// //     //delay at random, so cyber Frieza doesn't try to take down GoCU
// //     //in other words, so our crawler doesn't get removed for making too many patterned requests
// //     RandomDelay: 1 * time.Second,
// // })
// // c.Visit("https://craigslist.com")
// // c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// //     //get the link from anchor HTMLElement
// //     link := e.Attr("href")
// //     //get Colly collector to go to link
// //     c.Visit(e.Request.AbsoluteURL(link))
// // })
// // c.OnHTML("title", func(e *colly.HTMLElement) {
// //     fmt.Println(e.Text)
// // })
// // numVisited := 0
// // c.OnRequest(func(r *colly.Request) {
// //     if numVisited > 100 {
// //         r.Abort()
// //     }
// //     numVisited++
// // })
// // c.OnResponse(func(r *colly.Response) {
// //     fmt.Println(r.Body)
// // })
// // c.OnHTML("#items", func(e *colly.HTMLElement) {
// //     fmt.Println(e.ChildText("item"))
// // })
// // c.OnHTML("#items", func(e *colly.HTMLElement) {
// //     e.ForEach("item", func(_ int, elem *colly.HTMLElement) {
// //         if strings.Contains(elem.Text, "item") {
// //             fmt.Println(elem.Text)
// //         }
// //     })
// // })
// // dom, _ := goquery.NewDocument(htmlData)
// // dom.Find("a").Siblings().Each(func (i int, s *goquery.Selection)  {
// //     fmt.Printf("%d, Sibling text: %s\n", i, s.Test())
// // })
// // anchor.ParentsUntil("~").Find("item").Text()
// // c.OnHTML("div",func(e *colly.HTMLElement){
// //     // Goquery selection of the HTMLElement is in e.DOM
// //     goquerySelection := e.DOM
// //     fmt.Println(goquerySelection.Find(" space").Children().Text())
// // })
// func main() {
//     c := colly.NewCollector(
//         colly.AllowedDomains("craigslist.org"),
//     )
//     Callback for when a scraped page contains an article element
//     c.OnHTML("article", func(e *colly.HTMLElement) {
//         isPage := false
//         // Extract meta tags from the document
//         metaTags := e.DOM.ParentsUntil("~").Find("meta")
//         metaTags.Each(func(_ int, s *goquery.Selection) {
//             // Search for og:type meta tags
//             property, _ := s.Attr("property")
//             if strings.EqualFold(property, "og:type") {
//                 content, _ := s.Attr("content")
//                 // check pages
//                 isPage = strings.EqualFold(content, "article")
//             }
//         })
//         if isPage {
//             // Find the page title
//             fmt.Println("Title: ", e.DOM.Find("h1").Text())
//             // Grab all the text from description
//             fmt.Println(
//                 "Description: ",
//                 e.DOM.Find(".description").Find("p").Text())
//         }
//     })
//     // Callback for links on scraped pages
//     c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//         // Extract the linked URL from the anchor tag
//         link := e.Attr("href")
//         // Have our crawler visit the linked URL
//         c.Visit(e.Request.AbsoluteURL(link))
//     })
//     c.Limit(&colly.LimitRule{
//         DomainGlob:  "*",
//         RandomDelay: 1 * time.Second,
//     })
//     c.OnRequest(func(r *colly.Request) {
//         fmt.Println("Visiting", r.URL.String())
//     })
//     c.Visit("https://craigslist.org")
// }
// // c.OnHTML("html", func(e *colly.HTMLElement) {
// //     if strings.EqualFold(e.ChildAttr(`meta[property="og:type"]`, "content"), "article") {
// //         // Find the page title
// //         fmt.Println("Title: ", e.ChildText("article h1"))
// //         // Grab all the text from the page's description
// //         fmt.Println("Description: ", e.ChildText("article .description p"))
// //     }
// // })
//
