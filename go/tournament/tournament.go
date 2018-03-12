package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
)

const header = "Team                           | MP |  W |  D |  L |  P"

type result struct {
	Name          string
	MatchesPlayed int
	MatchesWon    int
	MatchesDrawn  int
	MatchesLost   int
	Points        int
}

type table []*result

func (tab table) Len() int {
	return len(tab)
}

func (tab table) Less(i, j int) bool {
	if tab[i].Points > tab[j].Points {
		return true
	} else if tab[i].Points < tab[j].Points {
		return false
	}
	return tab[i].Name < tab[j].Name
}

func (tab table) Swap(i, j int) {
	tab[i], tab[j] = tab[j], tab[i]
}

func (tab table) WriteTo(output io.Writer) (int64, error) {
	var written int64
	n, err := fmt.Fprintln(output, header)
	written += int64(n)
	if err != nil {
		return written, err
	}
	for _, r := range tab {
		n, err := fmt.Fprintf(output, "%-31s|%3d |%3d |%3d |%3d |%3d\n",
			r.Name,
			r.MatchesPlayed,
			r.MatchesWon, r.MatchesDrawn, r.MatchesLost,
			r.Points)
		written += int64(n)
		if err != nil {
			return written, err
		}
	}
	return written, nil
}

// Tally writes score table to a file
func Tally(input io.Reader, output io.Writer) error {
	results := make(map[string]*result)
	reader := csv.NewReader(input)
	reader.Comma = ';'
	reader.Comment = '#'
	reader.ReuseRecord = true
	reader.FieldsPerRecord = 3
	for {
		game, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		home, exists := results[game[0]]
		if !exists {
			home = &result{Name: game[0]}
			results[game[0]] = home
		}
		guest, exists := results[game[1]]
		if !exists {
			guest = &result{Name: game[1]}
			results[game[1]] = guest
		}
		home.MatchesPlayed++
		guest.MatchesPlayed++
		switch game[2] {
		case "win":
			guest.MatchesLost++
			home.MatchesWon++
			home.Points += 3
		case "loss":
			home.MatchesLost++
			guest.MatchesWon++
			guest.Points += 3
		case "draw":
			home.MatchesDrawn++
			guest.MatchesDrawn++
			home.Points++
			guest.Points++
		default:
			return errors.New("Unknown result")
		}
	}
	tab := make(table, len(results))
	i := 0
	for _, res := range results {
		tab[i] = res
		i++
	}
	sort.Sort(tab)
	_, err := tab.WriteTo(output)
	return err
}
