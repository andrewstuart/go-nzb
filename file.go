package nzb

import (
	"fmt"
	"strings"
)

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
	Id     string `xml:",innerxml"`
}

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
