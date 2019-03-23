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

func Combine(fnames []string, finishedName string) (err error) {
	f, err := os.Open("phrases/phrase1.mid")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := smfreader.New(f)
	err = rd.ReadHeader()
	if err != nil {
		panic(err)
	}
	header := rd.Header()
	var bf bytes.Buffer
	wr := smfwriter.New(&bf, smfwriter.TimeFormat(header.TimeFormat), smfwriter.Format(smf.SMF0))
	wr.Write(meta.TimeSig{
		Numerator:                4,
		Denominator:              4,
		ClocksPerClick:           24,
		DemiSemiQuaverPerQuarter: 8,
	})
	wr.Write(meta.BPM(90))
	for _, fname := range fnames {
		err = addMidi(fname, &wr)
		if err != nil {
			return
		}
	}
	wr.Write(meta.EndOfTrack)
	err = ioutil.WriteFile(finishedName, bf.Bytes(), 0644)
	return
}

func addMidi(fname string, wr *smf.Writer) (err error) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := smfreader.New(f)
	err = rd.ReadHeader()
	if err != nil {
		return
	}
	header := rd.Header()
	foo, err := strconv.Atoi(strings.Fields(header.TimeFormat.String())[0])
	if err != nil {
		return
	}
	ticksPerQuarterNote := uint32(foo)
	fmt.Println(ticksPerQuarterNote)
	var m midi.Message

	for {
		m, err = rd.Read()
		if err != nil {
			break
		}

		switch v := m.(type) {
		case channel.ControlChange:
			fmt.Println(v)
			(*wr).Write(v)
		case channel.NoteOn:
			fmt.Printf("%d\ton key: %v velocity: %v\n", rd.Delta(), v.Key(), v.Velocity())
			(*wr).SetDelta(rd.Delta())
			(*wr).Write(channel.Channel0.NoteOn(v.Key(), v.Velocity()))
		case channel.NoteOff:
			fmt.Printf("%d\toff key: %v\n", rd.Delta(), v.Key())
			(*wr).SetDelta(rd.Delta())
			(*wr).Write(channel.Channel0.NoteOff(v.Key()))
		}

	}

	if err == smf.ErrFinished {
		err = nil
	}

	return
}
