package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	unknownChordMessage = "Unknown chord"
	blankChordMessage   = "You didn't said chord"
)

var (
	chord, parallelChord string
	err                  error
	majorToMinor         bool
	chords               = map[string]string{
		"C": "Am",
		"D": "Hm",
		"E": "C#m",
		"F": "Dm",
		"G": "Em",
		"A": "F#m",
		"H": "G#m",
	}
)

func main() {
	fmt.Printf("Say chord: ")
	fmt.Scanf("%s", &chord)
	if strings.ContainsAny(chord, "m") {
		parallelChord, err = getMajor(chord)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		parallelChord, err = getMinor(chord)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("%s chord have %s parallel chord\n", chord, parallelChord)
}

func getMinor(major string) (string, error) {
	if major == "" {
		return "", errors.New(blankChordMessage)
	}
	if _, ok := chords[major]; !ok {
		return "", errors.New(unknownChordMessage)
	}
	return chords[major], nil
}

func getMajor(minor string) (string, error) {
	if minor == "" {
		return "", errors.New(blankChordMessage)
	}
	for k, v := range chords {
		if v == minor {
			return k, nil
		}
	}
	return "", errors.New(unknownChordMessage)
}
