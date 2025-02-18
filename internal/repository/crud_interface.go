package repository

type CrudInterface[T any] interface {
	Create(entity *T) error
	GetById(id uint) (*T, error)
	Update(entity *T) error
	Delete(entity *T) error
}
