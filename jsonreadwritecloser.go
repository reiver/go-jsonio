package jsonio

import (
	"io"
)

type JSONReadWriteCloser interface {
	JSONReader
	JSONWriter
	io.Closer
}
