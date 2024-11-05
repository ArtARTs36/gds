package gds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_List(t *testing.T) {
	m := NewMap[string, string]()

	m.Set("1", "2")
	m.Set("3", "4")

	m.Set("1", "5")

	assert.Equal(t, []string{"5", "4"}, m.List())
}

func TestMap_Delete(t *testing.T) {
	type kv struct {
		key   string
		value string
	}

	cases := []struct {
		Title        string
		Data         []kv
		DeletingKey  string
		ExpectedList []string
	}{
		{
			Title: "remove first item",
			Data: []kv{
				{key: "1", value: "a"},
				{key: "2", value: "b"},
				{key: "3", value: "c"},
				{key: "4", value: "d"},
			},
			DeletingKey:  "1",
			ExpectedList: []string{"b", "c", "d"},
		},
		{
			Title: "remove middle item",
			Data: []kv{
				{key: "1", value: "a"},
				{key: "2", value: "b"},
				{key: "3", value: "c"},
				{key: "4", value: "d"},
			},
			DeletingKey:  "2",
			ExpectedList: []string{"a", "c", "d"},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Title, func(t *testing.T) {
			m := NewMap[string, string]()

			for _, v := range tCase.Data {
				m.Set(v.key, v.value)
			}

			m.Delete(tCase.DeletingKey)

			assert.Equal(t, tCase.ExpectedList, m.List())
		})
	}

	t.Run("complex test", func(t *testing.T) {
		m := NewMap[string, string]()

		m.Set("a", "1")
		m.Set("b", "2")
		m.Set("c", "3")
		m.Set("d", "4")

		m.Delete("b")

		assert.Equal(t, []string{"1", "3", "4"}, m.List())

		m.Set("f", "5")

		assert.Equal(t, []string{"1", "3", "4", "5"}, m.List())

		m.Delete("d")

		assert.Equal(t, []string{"1", "3", "5"}, m.List())

		m.Delete("c")

		assert.Equal(t, []string{"1", "5"}, m.List())
	})
}

func TestMap_DeleteMany(t *testing.T) {
	type kv struct {
		key   string
		value string
	}

	cases := []struct {
		Title        string
		Data         []kv
		DeletingKeys []string
		ExpectedList []string
	}{
		{
			Title: "remove first item",
			Data: []kv{
				{key: "1", value: "a"},
				{key: "2", value: "b"},
				{key: "3", value: "c"},
				{key: "4", value: "d"},
			},
			DeletingKeys: []string{"1"},
			ExpectedList: []string{"b", "c", "d"},
		},
		{
			Title: "remove first items",
			Data: []kv{
				{key: "1", value: "a"},
				{key: "2", value: "b"},
				{key: "3", value: "c"},
				{key: "4", value: "d"},
			},
			DeletingKeys: []string{"1", "2"},
			ExpectedList: []string{"c", "d"},
		},
		{
			Title: "remove middle item",
			Data: []kv{
				{key: "1", value: "a"},
				{key: "2", value: "b"},
				{key: "3", value: "c"},
				{key: "4", value: "d"},
				{key: "5", value: "f"},
			},
			DeletingKeys: []string{"3"},
			ExpectedList: []string{"a", "b", "d", "f"},
		},
		{
			Title: "remove middle items",
			Data: []kv{
				{key: "1", value: "a"},
				{key: "2", value: "b"},
				{key: "3", value: "c"},
				{key: "4", value: "d"},
				{key: "5", value: "f"},
				{key: "6", value: "g"},
			},
			DeletingKeys: []string{"3", "4"},
			ExpectedList: []string{"a", "b", "f", "g"},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Title, func(t *testing.T) {
			m := NewMap[string, string]()

			for _, v := range tCase.Data {
				m.Set(v.key, v.value)
			}

			m.DeleteMany(tCase.DeletingKeys)

			assert.Equal(t, tCase.ExpectedList, m.List())
		})
	}

	t.Run("complex test", func(t *testing.T) {
		m := NewMap[string, string]()

		m.Set("a", "1")
		m.Set("b", "2")
		m.Set("c", "3")
		m.Set("d", "4")

		m.DeleteMany([]string{"b"})

		assert.Equal(t, []string{"1", "3", "4"}, m.List())

		m.Set("f", "5")

		assert.Equal(t, []string{"1", "3", "4", "5"}, m.List())

		m.DeleteMany([]string{"d"})

		assert.Equal(t, []string{"1", "3", "5"}, m.List())

		m.DeleteMany([]string{"c"})

		assert.Equal(t, []string{"1", "5"}, m.List())
	})
}
