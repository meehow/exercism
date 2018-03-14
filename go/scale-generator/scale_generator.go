package scale

import (
	"strings"
)

var (
	withSharps = [...]string{"C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
	sharps     = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	flats      = [12]string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
)

func useSharps(tonic string) bool {
	for _, s := range withSharps {
		if tonic == s {
			return true
		}
	}
	return false
}

func fullScale(tonic string, scaleSet [12]string) []string {
	if tonic[0] >= 'a' {
		tonic = strings.Title(tonic)
	}
	scale := make([]string, 12)
	i := 0
	j := 0
	start := false
	for j < 12 {
		if !start && scaleSet[i] == tonic {
			start = true
		}
		if start {
			scale[j] = scaleSet[i%12]
			j++
		}
		i++
	}
	return scale
}

// Scale returns a scale for a given tonic and interval
func Scale(tonic, interval string) []string {
	var fs []string
	if useSharps(tonic) {
		fs = fullScale(tonic, sharps)
	} else {
		fs = fullScale(tonic, flats)
	}
	if interval == "" {
		return fs
	}
	ss := make([]string, 0, 12)
	offset := 0
	for i, in := range interval {
		ss = append(ss, fs[i+offset])
		switch in {
		case 'M':
			offset++
		case 'A':
			offset += 2
		}
	}
	return ss
}
