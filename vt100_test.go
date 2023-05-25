package vt100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResize(t *testing.T) {
	vt := NewVT100(1, 2)

	err := vt.Process(runeCommand(rune(65)))
	assert.NoError(t, err)

	vt.Resize(1, 1)
	err = vt.Process(runeCommand(rune(65)))
	assert.NoError(t, err)
}
