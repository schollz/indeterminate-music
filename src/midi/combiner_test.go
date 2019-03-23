package midi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	err := Combine([]string{
		"phrases/phrase1.mid.trF.mid",
		"phrases/phrase3.mid.trF.mid",
		"phrases/phrase2.mid.trD.mid",
		"phrases/phrase4.mid.trD.mid",
	}, "song1.mid")
	assert.Nil(t, err)
}
