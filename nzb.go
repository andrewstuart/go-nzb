package nzb

import (
	"encoding/xml"
	"html"
)

type NZB struct {
	XMLName xml.Name `xml:"nzb"`
	Meta    Meta     `xml:"head>meta"`
	Files   []File   `xml:"file"`

	size int
}

type Meta map[string]string

func (m *Meta) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tag := struct {
		Type  string `xml:"type,attr"`
		Value string `xml:",innerxml"`
	}{}

	err := d.DecodeElement(&tag, &start)

	if err != nil {
		return err
	}

	if *m == nil {
		*m = make(map[string]string)
	}

	(*m)[tag.Type] = html.UnescapeString(tag.Value)

	return nil
}

func (nzb *NZB) Size() int {
	if nzb.size == 0 {
		for _, f := range nzb.Files {
			for _, seg := range f.Segments {
				nzb.size += seg.Bytes
			}
		}
	}

	return nzb.size
}
