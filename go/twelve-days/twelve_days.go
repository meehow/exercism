package twelve

import (
	"fmt"
	"strings"
)

const (
	testVersion = 1
	template    = "On the %s day of Christmas my true love gave to me, %s."
)

type christmasDay struct {
	number string
	gift   string
}

var days = []christmasDay{
	{"first", "and a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Verse(number int) string {
	if number == 1 {
		return fmt.Sprintf(template, days[0].number, days[0].gift[4:])
	}
	gifts := make([]string, 0, number)
	for i := number - 1; i >= 0; i-- {
		gifts = append(gifts, days[i].gift)
	}
	return fmt.Sprintf(template, days[number-1].number, strings.Join(gifts, ", "))
}

func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n") + "\n"
}
