package iomanager

type IOMananager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
