package nzb

import (
	"fmt"
	"strings"
)

//File is a go struct representation of the nzb file elements. It contains the
//appropriate field tags and methods for deserialization
type File struct {
	Poster   string    `xml:"poster,attr"`
	Date     int       `xml:"date,attr"`
	Subject  string    `xml:"subject,attr"`
	Groups   []string  `xml:"groups>group,internalxml"`
	Segments []Segment `xml:"segments>segment"`
}

//A Segment is a piece to be downloaded separately
type Segment struct {
	Number int    `xml:"number,attr"`
	Bytes  int    `xml:"bytes,attr"`
	ID     string `xml:",innerxml"`
}

//Name returns the estimated filename that an NZB represents.
func (f *File) Name() (string, error) {
	parts := strings.Split(f.Subject, `"`)

	n := ""
	if len(parts) > 1 {
		n = strings.Replace(parts[1], "/", "-", -1)
	} else {
		return "", fmt.Errorf("could not parse subject")
	}
	return n, nil
}
