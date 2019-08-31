package main

import (
	hrefparser "ayushbhargav/link_parser/hrefParser"
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("test.htm")
	if err != nil {
		panic(err)
	}

	fmt.Println(hrefparser.Parse(string(content)))
}
