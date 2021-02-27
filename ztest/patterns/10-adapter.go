package patterns

import (
	"fmt"
)

type iplayer interface {
	playmusic(str string)
}

func play(p iplayer, str string) {
	p.playmusic(str)
}

type musicplayer struct {
}

func (m *musicplayer) playmusic(str string) {
	fmt.Println("music player play music: ", str)
}

type gameplayer struct {
}

func (g *gameplayer) playsound(str string) {
	fmt.Println("game player play sound:", str)
}

type gameadatper struct {
	game *gameplayer
}

func (a *gameadatper) playmusic(str string) {
	a.game.playsound(str)
}
