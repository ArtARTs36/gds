package gds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings_Wrap(t *testing.T) {
	strs := NewStrings("a", "b", "c")

	assert.Equal(t, []string{"'a'", "'b'", "'c'"}, strs.Wrap("'").items)
}
