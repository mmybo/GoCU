package main
import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)
func main(){
    foundUrls := make(map[string]bool) //any clues as to what this is?
    seedUrls := os.Args[1:] //good ol' args
    //the channels
    chUrls := make(chan string) //channel dict
    chFinished := make(chan bool) //the channel to tell us it is finished
    //starting to crawl
    for _, url := range seedUrls {
        go crawl(url, chUrls,chFinished)
    }
    //subscribe to my channel!
    for c := 0; c < len(seedUrls); {
        select {
        case url := <-chUrls:
            foundUrls[url] = true
        case <-chFinished:
            c++
        }
    }
    //time to print!
    fmt.Println("\nFound", len(foundUrls), "unique urls:\n")
    for url, _ := range foundUrls {
        fmt.Println(" - " + url)
    }
    //it is time for this channel... to end
    close(chUrls)

    //ignore this stuff
    //resp, _ := http.Get(url)
    //bytes, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("HTML:\n\n",string(bytes))
    //resp.Body.Close()

}
// Get href from token
func getHref(t html.Token) (ok bool, href string) {
    for _, a := range t.Attr {
        if a.Key == "href" {
            href = a.Val
            ok = true
        }
    }
	return
}

//extract all https from page
func crawl(url string, ch chan string, chFinished chan bool){
    resp, err := http.Get(url)
    defer func(){
        chFinished <- true
    }()
    if err != nil {
        fmt.Println("ERROR: The crawl failed, daddy! WAAAAH! \"" + url + "\"")
        return
    }
    b := resp.Body
    defer b.Close() //close dat Body, once you return dat function, gurl!
    z := html.NewTokenizer(b) //Need dem html tokens, yessir!
    for {
        tt := z.Next()
        switch {
        case tt == html.ErrorToken:
            //it is time for this document parsing... to end
            return
        case tt == html.StartTagToken:
            t := z.Token() //dat zee token, boi!
            //is this an a tag?
            isAnchor := t.Data == "a"
            if isAnchor {
                fmt.Println("We found it!")
                fmt.Println("Praise the lord! This link is working!")
            } else {
                fmt.Println("We didn't find anything!")
                fmt.Println("Damn you! You gave us a faulty link!")
                continue
            }
            //getting dat href
            ok, url := getHref(t)
            if !ok {
                continue
            }
            //my href don't want nothin' less you got an http, hun!
            hasProto := strings.Index(url, "http") == 0
            if hasProto {
                ch <- url
            }
    }
    }
}
