package nzb

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestNzb(t *testing.T) {
	r := strings.NewReader(testNzb)

	dec := xml.NewDecoder(r)

	n := &NZB{}

	dec.Decode(n)

	if n.Meta["category"] != "TV > HD" {
		t.Errorf("Wrong category: %s", n.Meta["category"])
	}

	if len(n.Files) != 41 {
		t.Fatalf("Wrong number of files: %d", len(n.Files))
	}

	f := n.Files[0]

	if len(f.Groups) != 1 {
		t.Errorf("Wrong number of groups for file 1: %d", len(f.Groups))
	}

	if len(f.Segments) != 3 {
		t.Errorf("Wrong number of segments for file 1: %d", len(f.Segments))
	}
}
