package reader

import (
	"github.com/andyinabox/go-dumbphone/internal/utils"
	"testing"
)

const (
	url = "http://magazine.art21.org/2012/01/31/5-questions-for-contemporary-practice-with-claire-pentecost/#.XaomsCV7mL4"
)

func TestReader(t *testing.T) {
	_, err := GetPage(url)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReaderToFile(t *testing.T) {
	f, err := utils.CreateTempFile("")
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}

	html, err := GetPage(url)
	if err != nil {
		t.Fatal(err)
	}

	_, err = f.WriteString(html)
	if err != nil {
		t.Fatal(err)
	}

	if testing.Verbose() {
		utils.BrowserSend(f)
	}
}
