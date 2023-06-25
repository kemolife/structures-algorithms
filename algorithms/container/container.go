package container

// Container interface implement standard key-value storage methods where key is type Int and
type Container[T any] interface {
	Get() T
	Insert(value T)
	Update(index int)
	Delete(index int)
	Clear()
	Count() int
}
