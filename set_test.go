package gds

import (
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	set := NewSet[string]()

	set.Add("1")
	set.Add("1")
	set.Add("2")
	set.Add("1")
	set.Add("3")
	set.Add("3")

	assert.Equal(t, []string{
		"1", "2", "3",
	}, set.List())
}

func TestEqual(t *testing.T) {
	cases := []struct {
		Title string

		One *Set[int]
		Two *Set[int]

		Expected bool
	}{
		{
			Title:    "check empty sets",
			One:      NewSet[int](),
			Two:      NewSet[int](),
			Expected: true,
		},
		{
			Title:    "check identical sets",
			One:      NewSet[int](1, 2),
			Two:      NewSet[int](1, 2),
			Expected: true,
		},
		{
			Title:    "check identical sets with different orders",
			One:      NewSet[int](1, 2),
			Two:      NewSet[int](2, 1),
			Expected: true,
		},
		{
			Title:    "check different sets with different lengths",
			One:      NewSet[int](1, 2),
			Two:      NewSet[int](2),
			Expected: false,
		},
		{
			Title:    "check different sets",
			One:      NewSet[int](1, 2),
			Two:      NewSet[int](2, 3),
			Expected: false,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Title, func(t *testing.T) {
			assert.Equal(t, tCase.Expected, tCase.One.Equal(tCase.Two))
		})
	}
}

func TestUnmarshalYAML(t *testing.T) {
	t.Run("scalar", func(t *testing.T) {
		spec := struct {
			IDs *Set[int] `yaml:"ids"`
		}{}

		specVal := `
ids: [1,2,3,4,5]
`

		err := yaml.Unmarshal([]byte(specVal), &spec)
		require.NoError(t, err)

		assert.Equal(t, NewSet[int](1, 2, 3, 4, 5), spec.IDs)
	})
}
