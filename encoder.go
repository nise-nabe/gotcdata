// see http://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
package gotcdata

import (
	"io"
	"bytes"
	"encoding/xml"
)

type charsetISO88591er struct {
	r   io.ByteReader
	buf *bytes.Buffer
}

func newCharsetISO88591(r io.Reader) *charsetISO88591er {
	buf := bytes.Buffer{}
	return &charsetISO88591er{r.(io.ByteReader), &buf}
}

func (cs *charsetISO88591er) Read(p []byte) (n int, err error) {
	for _ = range p {
		if r, err := cs.r.ReadByte(); err != nil {
			break
		} else {
			cs.buf.WriteRune(rune(r))
		}
	}
	return cs.buf.Read(p)
}

func charsetReaderISO88791(charset string, input io.Reader) (io.Reader, error) {
	return newCharsetISO88591(input), nil
}

func unmarshal(data []byte, v interface{}) error {
	d := xml.NewDecoder(bytes.NewReader(data))
	d.CharsetReader = charsetReaderISO88791

	return d.Decode(v)
}
