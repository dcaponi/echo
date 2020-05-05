package interfaces

type Repository interface {
	Get() (interface{}, error)
	GetOne(id int) (interface{}, error)
	Write(ent interface{}) (interface{}, error)
}
