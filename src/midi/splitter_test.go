package midi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitter(t *testing.T) {
	err := Split("../../testing/MidiPieces.mid")
	assert.Nil(t, err)
}
