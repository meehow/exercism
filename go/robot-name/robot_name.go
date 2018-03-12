package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = newName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = newName()
}

var usedNames = make(map[string]bool)

func newName() string {
	var letters [2]byte
	letters[0] = byte(rand.Int31n('Z'-'A') + 'A')
	letters[1] = byte(rand.Int31n('Z'-'A') + 'A')
	name := fmt.Sprintf("%s%d", letters, rand.Int31n(900)+100)
	if usedNames[name] == true {
		return newName()
	}
	usedNames[name] = true
	return name
}
