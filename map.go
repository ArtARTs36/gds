package gds

import (
	"reflect"
	"slices"
)

type Map[K comparable, V any] struct {
	keyIndex map[K]int

	keys   []K
	values []V
	mapped map[K]V

	nilVal V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return NewMapFrom[K, V](map[K]V{})
}

func NewMapFrom[K comparable, V any](val map[K]V) *Map[K, V] {
	keys := []K{}
	values := []V{}
	keyIndex := map[K]int{}
	mapped := map[K]V{}

	i := 0

	for k, v := range val {
		keyIndex[k] = i
		keys = append(keys, k)
		values = append(values, v)
		mapped[k] = v

		i++
	}

	return &Map[K, V]{
		keyIndex: keyIndex,
		keys:     keys,
		values:   values,
		mapped:   mapped,
	}
}

func (m *Map[K, V]) First() V {
	if m.IsEmpty() {
		return m.nilVal
	}
	return m.values[0]
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
		m.mapped[key] = val

		return
	}

	m.keys[id] = key
	m.values[id] = val
	m.mapped[key] = val
}

func (m *Map[K, V]) Len() int {
	return len(m.values)
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	id, has := m.keyIndex[key]
	if !has {
		return m.nilVal, false
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

		if !reflect.DeepEqual(v, m.values[id]) {
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
	delete(m.mapped, deletingKey)

	for i := id + 1; i < len(m.values); i++ {
		values = append(values, m.values[i])
		keys = append(keys, m.keys[i])

		m.keyIndex[m.keys[i]]--
	}

	m.values = values
	m.keys = keys
}

func (m *Map[K, V]) ToMap() map[K]V {
	return m.mapped
}

func (m *Map[K, V]) IsEmpty() bool {
	return len(m.values) == 0
}

func (m *Map[K, V]) IsNotEmpty() bool {
	return len(m.values) > 0
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
			delete(m.mapped, key)
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

func (m *Map[K, V]) Keep(keys ...K) {
	newMap := m.CloneAndKeep(keys...)

	m.keyIndex = newMap.keyIndex
	m.keys = newMap.keys
	m.values = newMap.values
	m.mapped = newMap.mapped
}

func (m *Map[K, V]) CloneAndKeep(keys ...K) *Map[K, V] {
	newMap := NewMap[K, V]()

	for _, key := range keys {
		v, has := m.Get(key)
		if !has {
			continue
		}
		newMap.Set(key, v)
	}

	return newMap
}
