package main

import (
	"fmt"
	"os"
)

const c = "C"
const g = "G"
const d = "D"
const a = "A"
const e = "E"
const b = "B"
const f = "F"

const cm = "Cm"
const gm = "Gm"
const dm = "Dm"
const am = "Am"
const em = "Em"
const bm = "Bm"
const fm = "Fm"

// ChordCircle function
func ChordCircle(inChord string) (outChord string) {
	// var outChord string
	// if inChord == "C" {
	// 	outChord = "Am"
	// }
	return inChord + " chord have " + outChordFunc(inChord) + " parallel chord"
}

func outChordFunc(inChord string) (outChord string) {
	switch inChord {
	case c:
		outChord = am
	case g:
		outChord = em
	case d:
		outChord = bm
	case a:
		outChord = "F#m"
	case e:
		outChord = "C#m"
	case b:
		outChord = "G#m"
	case f:
		outChord = dm
	case am:
		outChord = c
	case em:
		outChord = g
	case bm:
		outChord = d
	case fm:
		outChord = "Ab"
	case cm:
		outChord = "Eb"
	case dm:
		outChord = f
	case gm:
		outChord = "Bb"
	}
	return
}

func main() {
	var chord string
	fmt.Print("Введите аккорд: ")
	fmt.Fscan(os.Stdin, &chord)
	fmt.Println(ChordCircle(chord))
}
