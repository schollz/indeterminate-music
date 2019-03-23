package midi

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyze(t *testing.T) {
	fnames, _ := filepath.Glob("phrase*.mid")
	ps, err := Analyze(fnames)
	assert.Nil(t, err)
	fmt.Println(ps)
}
