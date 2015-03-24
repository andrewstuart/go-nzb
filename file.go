package nzb

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
