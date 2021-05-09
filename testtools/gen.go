package testtools

import (
	"archive/tar"
	"io"
	"math/rand"
	"strconv"
	"sync/atomic"
)

var counter int32

// StrideByteReader allows for customizable "strides" of
// random bytes. Creates an infinite stream.
type StrideByteReader struct {
	stride    int
	counter   int
	randBytes []byte
}

// NewStrideByteReader creates a new random byte
// stride generator with a seed of 0.
func NewStrideByteReader(s int) *StrideByteReader {
	sb := StrideByteReader{
		stride:    s,
		randBytes: make([]byte, s),
	}

	rand.Seed(0)
	// rand.Seed(time.Now().UTC().UnixNano())
	rand.Read(sb.randBytes)
	return &sb
}

// Read creates randomly generated bytes of 'stride' length.
func (sb *StrideByteReader) Read(p []byte) (int, error) {
	l := len(sb.randBytes)

	n := 0
	for start := 0; start < len(p); n = copy(p[start:], sb.randBytes[sb.counter:]) {
		sb.counter = (sb.counter + n) % l
		start += n
	}

	return len(p), nil
}

// CreateTar creates a new tarball from the passed in reader
// and writes to a destination writer.
func CreateTar(w io.Writer, r *io.LimitedReader) {
	// defer TimeTrack(time.Now(), "CREATE TAR")
	tmp := atomic.AddInt32(&counter, 1)
	_ = tmp
	tw := tar.NewWriter(w)

	hdr := &tar.Header{
		Name: strconv.Itoa(int(counter)),
		Size: r.N,
		Mode: 0600,
	}

	if err := tw.WriteHeader(hdr); err != nil {
		panic(err)
	}

	if _, err := io.Copy(tw, r); err != nil {
		panic(err)
	}

	if err := tw.Close(); err != nil {
		panic(err)
	}
}
