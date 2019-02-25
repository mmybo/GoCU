package main

c := colly.NewCollector(
    //limit crawing to craigslist
    colly.AllowedDomains("craigslist.com"),
    //allow repeated visits to site
    colly.AllowURLRevisit(),
    //do parrel crawling, GoCU Style!
    colly.Async(true),
)
c.Limit(&colly.LimitRule{
    //filter websites according to limitations
    DomainGlob: ()"craigslist.com"),
    //Delay for requests, so as not to get in trouble with the Domain Lords
    Delay: 1 * time.Second,
    //delay at random, so cyber Frieza doesn't try to take down GoCU
    RandomDelay: 1 * time.Second,
})
c.Visit("https://craigslist.com")
c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    //get the link from anchor HTMLElement
    link := e.Attr("href")
    //get Colly collector to go to link
    c.Visit(e.Request.AbsoluteURL(link))
})
c.OnHTML("title", func(e *colly.HTMLElement) {
    fmt.Println(e.Text)
})
numVisited := 0
c.OnRequest(func(r *colly.Request) {
    if numVisited > 100 {
        r.Abort()
    }
    numVisited++
})
c.OnResponse(func(r *colly.Response) {
    fmt.Println(r.Body)
})
c.OnHTML("#items", func(e *colly.HTMLElement) {
    fmt.Println(e.ChildText("item"))
})
c.OnHTML("#items", func(e *colly.HTMLElement) {
    e.ForEach("item", func(_ int, elem *colly.HTMLElement) {
        if strings.Contains(elem.Text, "item") {
            fmt.Println(elem.Text)
        }
    })
})
dom, _ := goquery.NewDocument(htmlData)
dom.Find("a").Siblings().Each(func (i int, s *goquery.Selection)  {
    fmt.Printf("%d, Sibling text: %s\n", i, s.Test())
})
anchor.ParentsUntil("~").Find("item").Text()
c.OnHTML("div",func(e *colly.HTMLElement){
    // Goquery selection of the HTMLElement is in e.DOM
    goquerySelection := e.DOM
    fmt.Println(goquerySelection.Find(" space").Children().Text())
})
