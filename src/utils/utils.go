package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func closestDistanceToC(note string) int {
	if note == "C" {
		return 0
	}
	notes := []string{"Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	positions := []int{-1, -1}
	for i, n := range notes {
		if n == note {
			if i > 11 {
				positions[1] = i - 11
			} else {
				positions[0] = 11 - i
			}
		}
	}
	if positions[0] < positions[1] {
		return -1 * positions[0]
	} else if positions[1] > positions[0] {
		return positions[1]
	}
	return positions[0]
}

var chromatic = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var chromaticSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

func midiToNote(midiNum uint8) string {
	midiNumF := float64(midiNum)
	return fmt.Sprintf("%s%1.0f", chromatic[int(math.Mod(midiNumF, 12))], math.Floor(midiNumF/12.0-1))
}

func noteToMidi(note string) (midiNum uint8, err error) {
	octave := 4
	if len(note) == 3 {
		octave, err = strconv.Atoi(string(note[2]))
		note = note[:2]
	} else if len(note) == 2 {
		octave, err = strconv.Atoi(string(note[1]))
		note = string(note[0])
	}
	if err != nil {
		return
	}
	note = normalizeNote(note)
	midiNum = uint8(indexOf(note, chromatic) + 12*(octave+1))
	return
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func chordToNote(chord string) (note string) {
	return strings.Replace(strings.ToUpper(chord), "M", "", -1)
}

func normalizeNote(note string) string {
	for i, n := range chromaticSharp {
		if n == note {
			return chromatic[i]
		}
	}
	return note
}
