package domh

import (
	"errors"
	"net/http"

	"golang.org/x/net/html"
)

// GetTitle fetch a website title
func GetTitle(url string) (string, error) {

	node, err := getBody(url)
	if err != nil {
		return "", err
	}

	// find title
	title, ok := traverse(node)
	if !ok {
		return "", errors.New("Can not fetch title")
	}

	return title, nil
}

// GetBody get a website content
func getBody(url string) (*html.Node, error) {
	// fetch url
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("can not get the url")
	}

	// parse website body
	body, err := html.Parse(res.Body)
	if err != nil {
		return nil, errors.New("Parse error")
	}

	return body, nil
}

// check weather the element is title
func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

// traverse all over the elements to find title
func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}
