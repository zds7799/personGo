package chapter

import (
	"encoding/xml"
	"io"
	"os"
)

func BaseEncodingParse(xmlFile string) {
	f, err := os.Open(xmlFile)
	if err != nil {
		return
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:

		}
	}
}

func BaseETreeParse() {

}
