package interfaces

type Controller interface {
	GetOne(id int) []byte
	Get() []byte
	Write(a []byte) []byte
}
