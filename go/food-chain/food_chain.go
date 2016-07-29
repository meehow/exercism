package foodchain

import "strings"

const testVersion = 2

var swallowed = []string{"",
	"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}

var uniqLine = []string{"", "",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
}
var toCatch = func() []string {
	lines := make([]string, 9)
	for i := 8; i >= 4; i-- {
		lines[i] = "She swallowed the " + swallowed[i] + " to catch the " + swallowed[i-1] + "."
	}
	return lines
}()

func Verse(v int) string {
	lines := make([]string, 1, 9)
	lines[0] = "I know an old lady who swallowed a " + swallowed[v] + "."
	if v < 8 {
		if v >= 2 {
			lines = append(lines, uniqLine[v])
			if v >= 3 {
				for i := v; i >= 4; i-- {
					lines = append(lines, toCatch[i])
				}
				lines = append(lines, "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.")
			}
			lines = append(lines, "She swallowed the spider to catch the fly.")
		}
		lines = append(lines, "I don't know why she swallowed the fly. Perhaps she'll die.")
	} else {
		lines = append(lines, "She's dead, of course!")
	}
	return strings.Join(lines, "\n")

}

func Verses(v1, v2 int) string {
	return Verse(v1) + "\n\n" + Verse(v2)
}

func Song() string {
	verses := make([]string, 8)
	for i := 0; i < 8; i++ {
		verses[i] = Verse(i + 1)
	}
	return strings.Join(verses, "\n\n")
}
