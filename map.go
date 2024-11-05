package gds

import "slices"

type Map[K comparable, V comparable] struct {
	keyIndex map[K]int

	keys   []K
	values []V
}

func NewMap[K comparable, V comparable]() *Map[K, V] {
	return NewMapFrom[K, V](map[K]V{})
}

func NewMapFrom[K comparable, V comparable](val map[K]V) *Map[K, V] {
	keys := []K{}
	values := []V{}
	keyIndex := map[K]int{}

	i := 0

	for k, v := range val {
		keyIndex[k] = i
		keys = append(keys, k)
		values = append(values, v)

		i++
	}

	return &Map[K, V]{
		keyIndex: keyIndex,
		keys:     keys,
		values:   values,
	}
}

func (m *Map[K, V]) Clone() *Map[K, V] {
	return &Map[K, V]{
		keyIndex: m.keyIndex,
		keys:     m.keys,
		values:   m.values,
	}
}

func (m *Map[K, V]) Has(key K) bool {
	_, has := m.keyIndex[key]
	return has
}

func (m *Map[K, V]) List() []V {
	return m.values
}

func (m *Map[K, V]) Keys() []K {
	return m.keys
}

func (m *Map[K, V]) Set(key K, val V) {
	id, has := m.keyIndex[key]
	if !has {
		m.keyIndex[key] = len(m.values)
		m.keys = append(m.keys, key)
		m.values = append(m.values, val)

		return
	}

	m.keys[id] = key
	m.values[id] = val
}

func (m *Map[K, V]) Len() int {
	return len(m.values)
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	id, has := m.keyIndex[key]
	if !has {
		return (any(0)).(V), false
	}

	return m.values[id], true
}

func (m *Map[K, V]) Equal(that *Map[K, V]) bool {
	if m.Len() != that.Len() {
		return false
	}

	for k, id := range m.keyIndex {
		v, has := that.Get(k)
		if !has {
			return false
		}

		if v != m.values[id] {
			return false
		}
	}

	return true
}

func (m *Map[K, V]) Walk(callback func(key K, val V) bool) {
	for k, id := range m.keyIndex {
		callback(k, m.values[id])
	}
}

func (m *Map[K, V]) Delete(key K) {
	id, has := m.keyIndex[key]
	if !has {
		return
	}

	deletingKey := m.keys[id]

	values := m.values[0:id:id]
	keys := m.keys[0:id:id]

	delete(m.keyIndex, deletingKey)

	for i := id + 1; i < len(m.values); i++ {
		values = append(values, m.values[i])
		keys = append(keys, m.keys[i])

		m.keyIndex[m.keys[i]]--
	}

	m.values = values
	m.keys = keys
}

func (m *Map[K, V]) IsEmpty() bool {
	return len(m.values) == 0
}

func (m *Map[K, V]) DeleteMany(delkeys []K) {
	if m.IsEmpty() || len(delkeys) == 0 {
		return
	}

	ids := make([]int, 0, len(delkeys))
	for _, key := range delkeys {
		id, ok := m.keyIndex[key]
		if ok {
			ids = append(ids, id)

			delete(m.keyIndex, key)
		}
	}

	slices.Sort(ids)

	values := m.values[0:ids[0]:ids[0]]
	keys := m.keys[0:ids[0]:ids[0]]

	for i := ids[0]; i < len(m.values); i++ {
		if len(ids) > 0 && ids[0] == i {
			ids = ids[1:]

			continue
		}

		values = append(values, m.values[i])
		keys = append(keys, m.keys[i])

		m.keyIndex[m.keys[i]]--
	}

	m.values = values
	m.keys = keys
}
