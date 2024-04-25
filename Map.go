package search

type Map[K comparable, V any] struct{ m map[K][]V }

func New[K comparable, V any]() *Map[K, V] { return &Map[K, V]{map[K][]V{}} }
func (m *Map[K, V]) Store(keys []K, value V) {
	for _, key := range keys {
		m.m[key] = append(m.m[key], value)
	}
}
func (m *Map[K, V]) Load(keys ...K) []V {
	var values = []V{}
	for _, key := range keys {
		values = append(values, m.m[key]...)
	}
	return values
}
func (i *Map[K, V]) Clear(value V) { clear(i.m) }
