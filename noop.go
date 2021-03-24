package jsonio

var NoOp JSONReadWriteCloser

func init() {
	NoOp = internalNoOp{}
}
