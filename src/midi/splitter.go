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
)

// Split will take a midi file separated by rests and split
// it into separate midi files
func Split(fname string) (err error) {
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
	var totalTicks uint32
	var phraseNum int

	// initiate the first phrase
	bf.Reset()
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
			log.Debug(v)
			wr.Write(v)
		case channel.NoteOn:
			log.Debugf("%d\ton key: %v velocity: %v\n", rd.Delta(), v.Key(), v.Velocity())
			delta := rd.Delta()
			if rd.Delta() >= ticksPerQuarterNote*8 {
				log.Debug("new phrase")
				log.Debugf("total ticks: %d\n", totalTicks)
				if totalTicks > ticksPerQuarterNote*4 {
					err = fmt.Errorf("phrase %d is too long: %d", phraseNum, totalTicks)
					return
				}
				if totalTicks < ticksPerQuarterNote*4 {
					wr.SetDelta(ticksPerQuarterNote*4 - totalTicks)
					wr.Write(channel.Channel0.NoteOff(60))
				}
				wr.Write(meta.EndOfTrack)
				if phraseNum > 0 {
					// phrase 0 is empty
					ioutil.WriteFile(fmt.Sprintf("phrases/phrase%d.mid", phraseNum), bf.Bytes(), 0644)
					log.Debugf("wrote phrase%d.mid", phraseNum)
				}
				bf.Reset()
				totalTicks = 0
				phraseNum++
				wr = smfwriter.New(&bf, smfwriter.TimeFormat(header.TimeFormat), smfwriter.Format(smf.SMF0))
				wr.WriteHeader()
				wr.Write(meta.TimeSig{
					Numerator:                4,
					Denominator:              4,
					ClocksPerClick:           24,
					DemiSemiQuaverPerQuarter: 8,
				})
				wr.Write(meta.BPM(90))
				delta = 0
			}

			wr.SetDelta(delta)
			wr.Write(channel.Channel0.NoteOn(v.Key(), v.Velocity()))
			totalTicks += delta
		case channel.NoteOff:
			log.Debugf("%d\toff key: %v\n", rd.Delta(), v.Key())
			wr.SetDelta(rd.Delta())
			wr.Write(channel.Channel0.NoteOff(v.Key()))
			totalTicks += rd.Delta()
		}

	}
	log.Debug("new phrase")
	log.Debugf("total ticks: %d\n", totalTicks)
	if totalTicks < ticksPerQuarterNote*4 {
		wr.SetDelta(ticksPerQuarterNote*4 - totalTicks)
		wr.Write(channel.Channel0.NoteOff(60))
	}
	wr.Write(meta.EndOfTrack)

	err = ioutil.WriteFile(fmt.Sprintf("phrases/phrase%d.mid", phraseNum), bf.Bytes(), 0644)
	return err
}
