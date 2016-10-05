// phantomjs_test.go
package gophantomjs

import (
	"fmt"
	"testing"
)

func TestGetPageContent(t *testing.T) {
	s, _ := Get("http://www.oschina.net").
		SetUserAgent("PhantomJsServer").SetLoadImages(false).PageContent()
	fmt.Println(s)
}
func TestGetCookies(t *testing.T) {
	s, _ := Get("http://www.oschina.net").
		SetUserAgent("PhantomJsServer").SetLoadImages(false).GetCookies()
	fmt.Println(s)
}
