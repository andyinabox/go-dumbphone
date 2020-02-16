// download articles to read

package reader

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mauidude/go-readability"
)

// Article stores data from a given URL
type Article struct {
	URL   string
	Title string
	Body  string
}

// New creates a new article
func New(url string) (Article, error) {

	var a Article

	if url == "" {
		return a, errors.New("No URL provided")
	}

	a = Article{
		URL: url,
	}

	err := a.fetch()
	if err != nil {
		return a, err
	}

	return a, nil
}

// Render writes rendered html to given `io.Writer`
func (a *Article) Render(wr io.Writer, tplData []byte) error {

	funcMap := template.FuncMap{
		"unescape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	tpl, err := template.New("reader.html").Funcs(funcMap).Parse(string(tplData))
	if err != nil {
		return err
	}

	err = tpl.Execute(wr, a)
	if err != nil {
		return err
	}

	return nil
}

func (a *Article) fetch() error {

	res, err := http.Get(a.URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the original html doc as goquery doc
	qdoc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	// Get original html doc as string
	html, err := goquery.OuterHtml(qdoc.Selection)
	if err != nil {
		return err
	}

	// Create readable doc from html string
	rdoc, err := readability.NewDocument(html)
	if err != nil {
		return err
	}

	// Create goquery doc from readable doc
	rqdoc, err := goquery.NewDocumentFromReader(strings.NewReader(rdoc.Content()))
	if err != nil {
		return err
	}

	qdoc.Find("title").Each(func(i int, s *goquery.Selection) {
		a.Title = s.Text()
	})

	rqdoc.Find("body").Each(func(i int, s *goquery.Selection) {
		var h string
		h, err = s.Html()
		a.Body = h
	})
	if err != nil {
		return err
	}

	return nil
}
