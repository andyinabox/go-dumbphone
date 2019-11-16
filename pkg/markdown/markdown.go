package markdown

import (
	"gopkg.in/russross/blackfriday.v2"
)

// Parse markdown bytes into html bytes
func Parse(b []byte) ([]byte, error) {

	// why no error handling blackfriday?
	html := blackfriday.Run(b)

	return html, nil
}
