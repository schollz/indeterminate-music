package midi

import (
	"os"
	"strings"

	"github.com/gomidi/midi"
	"github.com/gomidi/midi/midimessage/channel"
	"github.com/gomidi/midi/smf"
	"github.com/gomidi/midi/smf/smfreader"
	"github.com/schollz/musical-keyboard/src/utils"
)

type Phrase struct {
	Filename   string
	TotalNotes int
	IsMinor    bool
}

func Analyze(fnames []string) (phrases []Phrase, err error) {
	phrases = make([]Phrase, len(fnames))
	i := 0
	for _, fname := range fnames {
		if strings.Contains(fname, ".tr") {
			continue
		}
		phrases[i], err = analyze(fname)
		if err != nil {
			return
		}
		i++
	}
	phrases = phrases[:i]
	return
}

func analyze(fname string) (p Phrase, err error) {
	p.Filename = fname

	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	rd := smfreader.New(f)

	var m midi.Message
	for {
		m, err = rd.Read()
		if err != nil {
			break
		}

		switch v := m.(type) {
		case channel.NoteOn:
			p.TotalNotes++
			if strings.Contains(utils.MidiToNote(v.Key()), "Eb") {
				p.IsMinor = true
			}
		}
	}

	if err == smf.ErrFinished {
		err = nil
	}
	return
}
