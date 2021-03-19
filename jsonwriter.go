package jsonio

type JSONWriter interface {
	WriteJSON(interface{}) error
}
