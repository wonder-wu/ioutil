package ioutil

import (
	"io"
	"strconv"
)

type intReader struct {
	reader io.Reader
	buffer [1]byte
}

//NewIntReader receieve a io.Reader interface to spwan a integer reader
func NewIntReader(r io.Reader) *intReader {

	return &intReader{r, [1]byte{}}
}

//ReadInt reads a integer each time called
func (r *intReader) ReadInt() (i int, err error) {
	var bt []byte

	parseStart := false
	for {
		n, rerr := r.reader.Read(r.buffer[:])
		if rerr != nil {

			if rerr == io.EOF {
				err = rerr
				break
			} else {
				return 0, rerr
			}
		}
		if n > 0 {

			if r.buffer[0] == '0' || r.buffer[0] == '1' || r.buffer[0] == '2' || r.buffer[0] == '3' || r.buffer[0] == '4' || r.buffer[0] == '5' || r.buffer[0] == '6' || r.buffer[0] == '7' || r.buffer[0] == '8' || r.buffer[0] == '9' {
				parseStart = true
				bt = append(bt, r.buffer[0])
			} else {
				if parseStart {
					//illegal charater found
					break
				}
			}

		} else {
			break
		}
	}

	if parseStart == false && err == io.EOF {
		return 0, err
	}
	i, err = strconv.Atoi(string(bt))
	return i, err

}
