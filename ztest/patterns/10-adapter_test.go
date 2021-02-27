package patterns

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	m := new(musicplayer)
	g := new(gameplayer)

	a := &gameadatper{game: g}

	play(m, "aaa")
	play(a, "bbb")
}
