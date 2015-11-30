package nzb

import (
	"encoding/xml"
	"html"
)

//NZB is a go struct representation of an nzb file with the appropriate tags
//and methods for XML unmarshalling.
type NZB struct {
	XMLName xml.Name `xml:"nzb"`
	Meta    Meta     `xml:"head>meta"`
	Files   []File   `xml:"file"`

	size uint64
}

//The Meta type is simply a map[string]string that implements UnmarshalXML to
//appropriately unmarshal the xml metadata tags.
type Meta map[string]string

//UnmarshalXML implements xml.Unmarshaler
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

//Size returns the sum of the file sizes within the nzb.
func (nzb *NZB) Size() uint64 {
	if nzb.size == 0 {
		for _, f := range nzb.Files {
			for _, seg := range f.Segments {
				nzb.size += uint64(seg.Bytes)
			}
		}
	}

	return nzb.size
}
