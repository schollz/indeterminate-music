package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Constants
var notes = []string{"Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var chromatic = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var chromaticSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

// ClosestDistanceToC returns the transposition nessecary to get to the closest C
func ClosestDistanceToC(note string) int {
	if note == "C" {
		return 0
	}
	note = normalizeNote(note)
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
	} else if positions[1] < positions[0] {
		return positions[1]
	}
	return positions[0]
}

// MidiToNote returns a name for the specified midi note
func MidiToNote(midiNum uint8) string {
	midiNumF := float64(midiNum)
	return fmt.Sprintf("%s%1.0f", chromatic[int(math.Mod(midiNumF, 12))], math.Floor(midiNumF/12.0-1))
}

// NoteToMidi takes a note name, like "C4" and tries to return
// a midi value, like 60.
func NoteToMidi(note string) (midiNum uint8, err error) {
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

// ChordToNote takes a chord like Cm and returns the note, C
func ChordToNote(chord string) (note string) {
	return strings.Replace(strings.Replace(chord, "M", "", -1), "m", "", -1)
}

func normalizeNote(note string) string {
	for i, n := range chromaticSharp {
		if n == note {
			return chromatic[i]
		}
	}
	return note
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
