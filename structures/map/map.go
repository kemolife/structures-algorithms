package _map

type (
	Map[K string, V any] struct {
		Size uint8
		Data []V
	}
)

func New[K string, V any](size uint8) *Map[K, V] {
	array := make([]V, size)
	return &Map[K, V]{Data: array, Size: size}
}

func (m *Map[K, V]) hash(k K) uint8 {
	var hash, i uint8
	hash = 0
	for i = 0; i < uint8(len(k))-1; i++ {
		hash = (hash + k[i]*i) % m.Size
	}

	return hash
}

func (m *Map[K, V]) Set(k K, v V) {
	hashKey := m.hash(k)
	m.Data[hashKey] = v
}

func (m *Map[K, V]) Get(k K) (v V) {
	hashKey := m.hash(k)
	return m.Data[hashKey]
}
