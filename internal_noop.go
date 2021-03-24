package jsonio

type internalNoOp struct {}

func (internalNoOp) Close() error {
	return nil
}

func (internalNoOp) ReadJSON(interface{}) error {
	return nil
}

func (internalNoOp) WriteJSON(interface{}) error {
	return nil
}
