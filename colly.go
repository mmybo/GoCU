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
    Delay: 1 * time.Second
    //delay at random, so cyber Judge Dredd doesn't try to take down GoCU
    RandomDelay: 1 * time.Second,
})
