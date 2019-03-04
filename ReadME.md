A Go Crawler Utility API (GoCU)
Simply make a post request using JSON, with the following format:
 {
     "Item": [word you are searching for],
     "URL": [web address you wish to search for it in]
 }
And GoCU will return all pages containing that word, and the links to them:
    {
        "Item List": [
            "[Page Title]",
            "[URL Linking to Page]"
        ]
    }
