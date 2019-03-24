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
		GetRandomPhrase(phrases, "D", true, true, 8, 40),
		GetRandomPhrase(phrases, "D", true, true, 8, 40),
		GetRandomPhrase(phrases, "F", false, true, 8, 40),
		GetRandomPhrase(phrases, "F", false, true, 8, 40),
		GetRandomPhrase(phrases, "D", true, true, 8, 44),
		GetRandomPhrase(phrases, "D", true, true, 8, 44),
		GetRandomPhrase(phrases, "A", true, true, 8, 40),
		GetRandomPhrase(phrases, "A", true, true, 8, 40),
	}
	fmt.Println(phraseList)
	err = Combine(phraseList, "songRH.mid")
	assert.Nil(t, err)
	phraseList = []string{
		GetRandomPhrase(phrases, "D", true, false, 0, 4),
		GetRandomPhrase(phrases, "D", true, false, 0, 4),
		GetRandomPhrase(phrases, "F", false, false, 0, 4),
		GetRandomPhrase(phrases, "F", false, false, 0, 4),
		GetRandomPhrase(phrases, "D", true, false, 0, 4),
		GetRandomPhrase(phrases, "D", true, false, 0, 4),
		GetRandomPhrase(phrases, "A", true, false, 0, 4),
		GetRandomPhrase(phrases, "A", true, false, 0, 4),
	}
	fmt.Println(phraseList)
	err = Combine(phraseList, "songLH.mid")
	assert.Nil(t, err)
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func GetRandomPhrase(phrases []Phrase, note string, minor bool, rh bool, minNotes, maxNotes int) string {
	for _, i := range r.Perm(len(phrases)) {
		if phrases[i].IsMinor != minor || phrases[i].RH != rh || phrases[i].TotalNotes < minNotes || phrases[i].TotalNotes > maxNotes {
			continue
		}
		return phrases[i].Filename + ".tr" + note + ".mid"
	}
	return ""
}
