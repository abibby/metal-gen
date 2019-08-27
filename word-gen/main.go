package main

import (
	"compress/bzip2"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type Redirect struct {
	Title string `xml:"title,attr"`
}

type Page struct {
	Title string `xml:"title"`
	Text  string `xml:"revision>text"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r, err := http.Get("https://dumps.wikimedia.freemirror.org/enwiktionary/20190820/enwiktionary-20190820-pages-articles-multistream.xml.bz2")
	check(err)

	defer r.Body.Close()

	decoder := xml.NewDecoder(bzip2.NewReader(r.Body))

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			// ...and its name is "page"
			if se.Name.Local == "page" {
				var p Page
				// decode a whole chunk of following XML into the
				// variable p which is a Page (se above)
				decoder.DecodeElement(&p, &se)

				if strings.Contains(p.Text, "==English==") {

					if strings.Contains(p.Text, "====Noun====") {
						fmt.Println(p.Title)
					}
					//  else if strings.Contains(p.Text, "====Verb====") {
					// 	fmt.Println(p.Title)
					// }
				}
			}
		}
	}
}
