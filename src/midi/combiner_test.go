package midi

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	fnames, _ := filepath.Glob("phrases/phrase*.mid")
	phrases, err := Analyze(fnames)
	assert.Nil(t, err)
	fmt.Println(phrases)
	phraseList := []string{
		GetRandomPhrase(phrases, "D", true),
		GetRandomPhrase(phrases, "D", true),
		GetRandomPhrase(phrases, "F", false),
		GetRandomPhrase(phrases, "F", false),
		GetRandomPhrase(phrases, "A", true),
		GetRandomPhrase(phrases, "A", true),
		GetRandomPhrase(phrases, "C", false),
		GetRandomPhrase(phrases, "C", false),
		GetRandomPhrase(phrases, "D", true),
		GetRandomPhrase(phrases, "D", true),
		GetRandomPhrase(phrases, "F", false),
		GetRandomPhrase(phrases, "F", false),
		GetRandomPhrase(phrases, "A", true),
		GetRandomPhrase(phrases, "A", true),
		GetRandomPhrase(phrases, "C", false),
		GetRandomPhrase(phrases, "C", false),
	}
	fmt.Println(phraseList)
	err = Combine(phraseList, "song1.mid")
	assert.Nil(t, err)
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func GetRandomPhrase(phrases []Phrase, note string, minor bool) string {
	for _, i := range r.Perm(len(phrases)) {
		if phrases[i].IsMinor != minor {
			continue
		}
		return phrases[i].Filename + ".tr" + note + ".mid"
	}
	return ""
}
