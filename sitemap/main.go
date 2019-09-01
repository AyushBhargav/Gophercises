package main

import (
	hrefparser "ayushbhargav/link_parser/hrefParser"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type queue struct {
	store []string
	front int
	rear  int
}

// NewQueue creates empty queue for BFS
func NewQueue() queue {
	return queue{store: make([]string, 20000), front: 0, rear: -1}
}

func (q *queue) add(s string) {
	q.rear++
	q.store[q.rear] = s
}

func (q *queue) remove() string {
	if q.front > q.rear {
		// Empty queue
		return ""
	}

	s := q.store[q.front]
	q.front++
	return s
}

func (q *queue) isEmpty() bool {
	if q.front > q.rear {
		return true
	}
	return false
}

type location struct {
	Loc string `xml:"loc"`
}

type siteMap struct {
	XMLName xml.Name   `xml:"urlset"`
	XMLNs   string     `xml:"xmlns,attr"`
	URLs    []location `xml:"url"`
}

func getSource(domain string) string {
	resp, err := http.Get(domain)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	httpStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(httpStr)
}

func getQualifiedURL(domain string, link hrefparser.Link) string {
	matched, err := regexp.Match(`^(/.*)|(#.*)`, []byte(link.Href))
	if err != nil {
		panic(err)
	}

	href := link.Href
	if matched {
		href = domain + href
	}

	return href
}

func mapToURLs(domain string, links []hrefparser.Link) []location {
	loc := make([]location, len(links))

	for index, link := range links {
		loc[index] = location{getQualifiedURL(domain, link)}
	}

	return loc
}

func validURL(domain string) bool {
	if strings.Contains(domain, "mailto") || len(domain) == 0 {
		return false
	}
	return true
}

func bfs(rootDomain string, domain string, links *[]hrefparser.Link, visited map[string]bool, level int) {
	if visited[domain] == true || !validURL(domain) || level == 0 {
		return
	}
	level--
	httpContent := getSource(domain)
	siblingLinks := hrefparser.Parse(string(httpContent))

	visited[domain] = true
	q := NewQueue()
	for _, siblingLink := range siblingLinks {
		q.add(getQualifiedURL(rootDomain, siblingLink))
	}

	*links = append(*links, siblingLinks...)

	for !q.isEmpty() {
		siblingDomain := q.remove()
		bfs(rootDomain, siblingDomain, links, visited, level)
	}
}

func main() {
	domain := flag.String("domain", "<html>", "Website domain")

	/*httpContent := getSource(*domain)
	links := hrefparser.Parse(string(httpContent))*/

	var links []hrefparser.Link
	bfs(*domain, *domain, &links, map[string]bool{}, 2)
	sm := siteMap{XMLNs: "http://www.sitemaps.org/schemas/sitemap/0.9", URLs: mapToURLs(*domain, links)}

	f, e := os.Create("sitemap.xml")
	if e != nil {
		panic(e)
	}

	_, e = f.WriteString(xml.Header)
	if e != nil {
		panic(e)
	}

	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	if err := enc.Encode(sm); err != nil {
		panic(err)
	}
}
