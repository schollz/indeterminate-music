package main

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

func main() {
	fmt.Println(midi2())

	f, err := os.Open("Test.mid")
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
	addMidi("phrase3.mid", &wr)
	addMidi("phrase3.mid", &wr)
	addMidi("phrase3.mid", &wr)
	wr.Write(meta.EndOfTrack)
	ioutil.WriteFile("combine.mid", bf.Bytes(), 0644)

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

	if err != smf.ErrFinished {
		panic("error: " + err.Error())
	}

	return
}

func midi2() (err error) {
	f, err := os.Open("MidiPieces.mid")
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

	var bf bytes.Buffer
	var wr smf.Writer
	wr = smfwriter.New(&bf, smfwriter.TimeFormat(header.TimeFormat), smfwriter.Format(smf.SMF0))
	wr.Write(meta.TimeSig{
		Numerator:                4,
		Denominator:              4,
		ClocksPerClick:           24,
		DemiSemiQuaverPerQuarter: 8,
	})
	wr.Write(meta.BPM(90))
	phraseNum := 0
	var totalTicks uint32
	for {
		m, err = rd.Read()

		// at the end, smf.ErrFinished will be returned
		if err != nil {
			break
		}
		switch v := m.(type) {
		case channel.ControlChange:
			fmt.Println(v)
			wr.Write(v)
		case channel.NoteOn:
			fmt.Printf("%d\ton key: %v velocity: %v\n", rd.Delta(), v.Key(), v.Velocity())
			delta := rd.Delta()
			if rd.Delta() >= ticksPerQuarterNote*8 {
				fmt.Println("new phrase")
				fmt.Printf("total ticks: %d\n", totalTicks)
				if totalTicks > ticksPerQuarterNote*4 {
					err = fmt.Errorf("phrase %d is too long: %d", phraseNum, totalTicks)
				}
				if phraseNum > 0 {
					if totalTicks < ticksPerQuarterNote*4 {
						wr.SetDelta(ticksPerQuarterNote*4 - totalTicks)
						wr.Write(channel.Channel0.NoteOff(60))
					}
					wr.Write(meta.EndOfTrack)
					ioutil.WriteFile(fmt.Sprintf("phrase%d.mid", phraseNum), bf.Bytes(), 0644)
					bf.Reset()
				}
				totalTicks = 0
				phraseNum++
				wr = smfwriter.New(&bf, smfwriter.TimeFormat(header.TimeFormat), smfwriter.Format(smf.SMF0))
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
			fmt.Printf("%d\toff key: %v\n", rd.Delta(), v.Key())
			wr.SetDelta(rd.Delta())
			wr.Write(channel.Channel0.NoteOff(v.Key()))
			totalTicks += rd.Delta()
		}

	}
	fmt.Println("new phrase")
	fmt.Printf("total ticks: %d\n", totalTicks)
	if totalTicks < ticksPerQuarterNote*4 {
		wr.SetDelta(ticksPerQuarterNote*4 - totalTicks)
		wr.Write(channel.Channel0.NoteOff(60))
	}
	wr.Write(meta.EndOfTrack)
	ioutil.WriteFile(fmt.Sprintf("phrase%d.mid", phraseNum), bf.Bytes(), 0644)

	if err != smf.ErrFinished {
		panic("error: " + err.Error())
	}

	return
}

// func midi1() {
// 	f, err := os.Open("MidiPieces.mid")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	decoder := midi.NewDecoder(f)
// 	decoder.Debug = true
// 	if err := decoder.Decode(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("format:", decoder.Format)
// 	fmt.Println(decoder.TicksPerQuarterNote, "ticks per quarter")
// 	fmt.Println("Debugger on:", decoder.Debug)

// 	for _, tr := range decoder.Tracks {
// 		for _, ev := range tr.Events {
// 			if ev.Velocity > 0 {
// 				fmt.Println(ev)
// 			}
// 			// if ev.MsgType == 8 || ev.MsgType == 9 {
// 			// 	if float64(ev.TimeDelta) >= float64(decoder.TicksPerQuarterNote)*6 {
// 			// 		fmt.Println("new phrase")
// 			// 	}
// 			// 	fmt.Println(ev)
// 			// }
// 		}
// 	}

// }
