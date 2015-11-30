# nzb
--
    import "github.com/andrewstuart/go-nzb"


## Usage

#### type File

```go
type File struct {
	Poster   string    `xml:"poster,attr"`
	Date     int       `xml:"date,attr"`
	Subject  string    `xml:"subject,attr"`
	Groups   []string  `xml:"groups>group,internalxml"`
	Segments []Segment `xml:"segments>segment"`
}
```

File is a go struct representation of the nzb file elements. It contains the
appropriate field tags and methods for deserialization

#### func (*File) Name

```go
func (f *File) Name() (string, error)
```
Name returns the estimated filename that an NZB represents.

#### type Meta

```go
type Meta map[string]string
```

The Meta type is simply a map[string]string that implements UnmarshalXML to
appropriately unmarshal the xml metadata tags.

#### func (*Meta) UnmarshalXML

```go
func (m *Meta) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```
UnmarshalXML implements xml.Unmarshaler

#### type NZB

```go
type NZB struct {
	XMLName xml.Name `xml:"nzb"`
	Meta    Meta     `xml:"head>meta"`
	Files   []File   `xml:"file"`
}
```

NZB is a go struct representation of an nzb file with the appropriate tags and
methods for XML unmarshalling.

#### func (*NZB) Size

```go
func (nzb *NZB) Size() uint64
```
Size returns the sum of the file sizes within the nzb.

#### type Segment

```go
type Segment struct {
	Number int    `xml:"number,attr"`
	Bytes  int    `xml:"bytes,attr"`
	ID     string `xml:",innerxml"`
}
```

A Segment is a piece to be downloaded separately
