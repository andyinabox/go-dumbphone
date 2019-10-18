// download articles to read

package reader

import (
	"fmt"
	"github.com/mauidude/go-readability"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Url string

	title string
	content string
}

func (p *Page) Fetch() error {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	html := fmt.Sprintf("%s\n", bytes)

	doc, err := readability.NewDocument(html)
	if err != nil {
		return "", err
	}

	content = 

	return doc.Content(), nil
}

func (p *Page) Content() string {
	return p.content
}


func (p *Page) Title() string {
	return p.title
}


// GetPage get url using readability and return HTML
func GetPage(url string) (string, error) {

}
