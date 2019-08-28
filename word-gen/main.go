package main

import (
	"compress/bzip2"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
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

	nounFile, err := os.OpenFile("nouns.mjs", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	check(err)
	adjectiveFile, err := os.OpenFile("adjectives.mjs", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	check(err)

	firstNoun := true
	firstAdjective := true
	_, err = fmt.Fprintf(nounFile, "export const Nouns = [\n")
	check(err)
	_, err = fmt.Fprintf(adjectiveFile, "export const Adjectives = [\n")
	check(err)

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
				check(decoder.DecodeElement(&p, &se))

				if strings.Contains(p.Text, "==English==") {

					if strings.Contains(p.Text, "====Noun====") {
						writeReccord(nounFile, p.Title, firstNoun)
						firstNoun = false
					} else if strings.Contains(p.Text, "====Adjective====") {
						writeReccord(adjectiveFile, p.Title, firstAdjective)
						firstAdjective = false
					}
				}
			}
		}
	}

	_, err = fmt.Fprintf(nounFile, "]\n")
	check(err)
	_, err = fmt.Fprintf(adjectiveFile, "]\n")
	check(err)
}

func writeReccord(file io.Writer, word string, isFirst bool) {
	b, err := json.Marshal(word)
	check(err)

	if !isFirst {
		_, err = fmt.Fprintf(file, ",\n")
		check(err)
	}

	_, err = fmt.Fprintf(file, "    %s", b)
	check(err)
}
