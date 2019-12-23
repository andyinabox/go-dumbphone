package reader

import (
	"testing"

	"github.com/andyinabox/go-dumbphone/pkg/browser"
	"github.com/andyinabox/go-dumbphone/pkg/utils"
)

const (
	url = "https://hyperallergic.com/523358/activists-infiltrate-moma-party/"
)

func TestNew(t *testing.T) {
	a, err := New(url)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if a.Title == "" {
		t.Fatal("Title field is empty")
	} else {
		t.Logf("Title: %s", a.Title)
	}

	if a.Body == "" {
		t.Fatal("Body field is empty")
	} else {
		t.Logf("Body: %s", a.Body)
	}
}

func TestRender(t *testing.T) {
	a, err := New(url)
	if err != nil {
		t.Fatalf("%s", err)
	}

	f, err := utils.CreateTempFile("")
	defer f.Close()
	if err != nil {
		t.Fatalf("%s", err)
	}

	err = a.Render(f, "./reader.html")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if testing.Verbose() {
		err = browser.OpenFile(f)
		if err != nil {
			t.Logf("Error opening in browser: %s", err)
		}
	}

}
