package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
	  <img src="xyx.jpg" width="104" height="142">
  </body>
  <script> console.log("Do something"); </script>
  <script> console.log("Do something  else"); </script>
</html>`

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse failed : %s\n", err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)
	fmt.Printf("Count of words %d and count of pics %d\n", words, pics)

}

func countWordsAndImages(doc *html.Node) (wCount int, pCount int) {

	visit(doc, &wCount, &pCount)
	return
}

func visit(n *html.Node, wCount, pCount *int) {
	// if its an element node, what tag does it have ?

	if n.Type == html.ElementNode && n.Data == "script" {
		return
	}

	if n.Type == html.TextNode {
		*wCount += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pCount++
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		visit(i, wCount, pCount)
	}

}
