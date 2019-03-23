package utils

import "fmt"

func ExampleClosestDistanceToC() {
	fmt.Println(ClosestDistanceToC("Bb"))
	fmt.Println(ClosestDistanceToC("E"))
	// Output: 2
	// -4
}

func ExampleMidiToNote() {
	fmt.Println(MidiToNote(60))
	// Output: C4
}

func ExampleNoteToMidi() {
	fmt.Println(NoteToMidi("D#3"))
	// Output: 51 <nil>
}

func ExampleChordToNote() {
	fmt.Println(ChordToNote("Bbm"))
	// Output: Bb
}
