package midi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gomidi/midi"
	"github.com/gomidi/midi/midimessage/channel"
	"github.com/gomidi/midi/midimessage/meta"
	"github.com/gomidi/midi/smf"
	"github.com/gomidi/midi/smf/smfreader"
	"github.com/gomidi/midi/smf/smfwriter"
	"github.com/schollz/musical-keyboard/src/utils"
)

// Transpose takes a file and transposes and writes it back
// as fname.mid.trA where it has been transposed to A.
func Transpose(fname string, note string) (err error) {
	transposeAmount := utils.ClosestDistanceToC(note)
	log.Debugf("note %s transposed %d", note, transposeAmount)

	// open the file
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	// read in the midi file and header
	rd := smfreader.New(f)
	err = rd.ReadHeader()
	if err != nil {
		return
	}

	// process the header
	header := rd.Header()
	foo, err := strconv.Atoi(strings.Fields(header.TimeFormat.String())[0])
	if err != nil {
		return
	}
	ticksPerQuarterNote := uint32(foo)
	log.Debug(ticksPerQuarterNote)

	var m midi.Message
	var bf bytes.Buffer
	var wr smf.Writer
	var errDone error

	// initiate the first phrase
	wr = smfwriter.New(&bf, smfwriter.TimeFormat(header.TimeFormat), smfwriter.Format(smf.SMF0))
	wr.WriteHeader()
	wr.Write(meta.TimeSig{
		Numerator:                4,
		Denominator:              4,
		ClocksPerClick:           24,
		DemiSemiQuaverPerQuarter: 8,
	})
	wr.Write(meta.BPM(90))

	for {
		// read in the data
		m, errDone = rd.Read()

		// at the end, smf.ErrFinished will be returned
		if errDone != nil {
			break
		}
		switch v := m.(type) {
		case channel.ControlChange:
			wr.Write(v)
		case channel.NoteOn:
			wr.SetDelta(rd.Delta())
			wr.Write(channel.Channel0.NoteOn(uint8(int(v.Key())+transposeAmount), v.Velocity()))
		case channel.NoteOff:
			wr.SetDelta(rd.Delta())
			wr.Write(channel.Channel0.NoteOff(uint8(int(v.Key()) + transposeAmount)))
		}

	}
	wr.Write(meta.EndOfTrack)

	err = ioutil.WriteFile(fmt.Sprintf("%s.tr%s.mid", fname, note), bf.Bytes(), 0644)
	return
}
