package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
	"strconv"
	"time"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func getForce(s string) uint8 {
	x, _ := strconv.ParseUint(s, 10, 8)
	if x == 0 {
		fmt.Printf("got 0 with %q\n", s)
		return 0
	}
	return uint8(60 + x/16)
}

var (
	ta time.Time
	tb time.Time
	tc time.Time
	td time.Time
	te time.Time
	tf time.Time
)

var (
	drumA *writer.Writer // 36 bass drum
	drumB *writer.Writer // 40 electric snare
	drumC *writer.Writer // 59 ride symbol2
	drumD *writer.Writer // 41 low floor tom
	drumE *writer.Writer // 42 closed high-hat
	drumF *writer.Writer // 46 open hi-hat
//drumG *writer.Writer // 49 crash symbol1
)

var velocity uint8
var t time.Duration

func playDrum(s string) {
	//fmt.Println(s)
	t = time.Duration(0)
	velocity = 0
	d := s[0]
	switch d {
	case 'a':
		{
			t = time.Since(ta)
			if t.Milliseconds() < 250 {
				return
			}
			velocity = getForce(s[1:])
			//fmt.Println(velocity)
			if velocity > 0 {
				writer.NoteOff(drumA, 36)
				writer.NoteOn(drumA, 36, velocity)
				ta = time.Now()
			}
			return
		}
	case 'b':
		{
			t = time.Since(tb)
			if t.Milliseconds() < 250 {
				return
			}
			velocity = getForce(s[1:])
			if velocity > 0 {
				writer.NoteOff(drumB, 40)
				writer.NoteOn(drumB, 40, velocity)
				tb = time.Now()
			}
			return
		}
	case 'c':
		{
			t = time.Since(tc)
			if t.Milliseconds() < 250 {
				return
			}
			velocity = getForce(s[1:])
			if velocity > 0 {
				writer.NoteOff(drumC, 59)
				writer.NoteOn(drumC, 59, velocity)
				tc = time.Now()
			}
			return
		}
	case 'd':
		{
			t = time.Since(td)
			if t.Milliseconds() < 250 {
				return
			}
			velocity = getForce(s[1:])
			if velocity > 0 {
				writer.NoteOff(drumD, 41)
				writer.NoteOn(drumD, 41, velocity)
				td = time.Now()
			}
			return
		}
	case 'e':
		{
			t = time.Since(te)
			if t.Milliseconds() < 250 {
				return
			}
			velocity = getForce(s[1:])
			if velocity > 0 {
				writer.NoteOff(drumE, 42)
				writer.NoteOn(drumE, 42, velocity)
				te = time.Now()
			}
			return
		}
	case 'f':
		{
			t = time.Since(tf)
			if t.Milliseconds() < 250 {
				return
			}
			velocity = getForce(s[1:])
			if velocity > 0 {
				writer.NoteOff(drumF, 46)
				writer.NoteOn(drumF, 46, velocity)
				tf = time.Now()
			}
			return
		}
	}
}

func main() {
	var midiNo int
	var com string
	flag.StringVar(&com, "s", "", "serial name, string")
	flag.IntVar(&midiNo, "p", -1, "midi synth output no, int")
	flag.Parse()
	if com == "" {
		panic("serial name must be specified")
	}
	if midiNo < 0 {
		panic("output midi device number can't be negative")
	}
	ta = time.Now()
	tb = ta
	tc = ta
	td = ta
	te = ta
	tf = ta
	c := &serial.Config{Name: com, Baud: 9600}
	s, err := serial.OpenPort(c)
	must(err)
	drv, err := driver.New()
	must(err)
	defer drv.Close()
	outs, err := drv.Outs()
	must(err)
	out := outs[midiNo]
	must(out.Open())
	drumA = writer.New(out)
	drumA.SetChannel(9)
	drumB = writer.New(out)
	drumB.SetChannel(9)
	drumC = writer.New(out)
	drumC.SetChannel(9)
	drumD = writer.New(out)
	drumD.SetChannel(9)
	drumE = writer.New(out)
	drumE.SetChannel(9)
	drumF = writer.New(out)
	drumF.SetChannel(9)

	scanner := bufio.NewScanner(s)
	msg := ""
	for {
		for scanner.Scan() {
			msg = scanner.Text()
			playDrum(msg)
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
