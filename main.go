package main

import (
	"net/http"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func main() {
	nogizaka := "http://www.nogizaka46.com/schedule/"
	resp, _ := http.PostForm(nogizaka, url.Values{"member": {"manatsu-akimoto"}})
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	result := doc.Find("div.scheduleTableList").Text()
	replaced := strings.Replace(result, "\n", "", 1)
	println(replaced)
}
