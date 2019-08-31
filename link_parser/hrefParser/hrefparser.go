package hrefparser

import (
	"strings"

	"golang.org/x/net/html"
)

// Link represents Href associated with given text
type Link struct {
	Href string
	Text string
}

func generateText(root *html.Node) string {
	if root == nil || root.Type == html.CommentNode {
		return ""
	}

	var text []string
	currentText := ""
	if root.Type == html.TextNode {
		currentText = strings.TrimSpace(root.Data)
		text = append(text, currentText)
	}

	childText := generateText(root.FirstChild)
	if len(childText) > 0 {
		text = append(text, childText)
	}
	siblingText := generateText(root.NextSibling)
	if len(siblingText) > 0 {
		text = append(text, siblingText)
	}

	return strings.Join(text, " ")
}

func generateLinks(root *html.Node, links *[]Link) {
	if root == nil {
		return
	}
	if root.DataAtom == 0x1 {
		val := generateText(root.FirstChild)
		href := ""
		for _, attribute := range root.Attr {
			if attribute.Key == "href" {
				href = attribute.Val
				break
			}
		}
		*links = append(*links, Link{href, val})
	} else {
		// DFS
		generateLinks(root.FirstChild, links)
	}
	// Traverse to next sibling
	generateLinks(root.NextSibling, links)
}

// Parse the HTML code and generates Links
func Parse(htmlContent string) []Link {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		panic(err)
	}
	var links []Link
	generateLinks(doc, &links)
	return links
}
