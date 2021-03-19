package jsonio

type JSONReader interface {
	ReadJSON(interface{}) error
}
