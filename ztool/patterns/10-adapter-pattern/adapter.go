package adapter

import "fmt"

/*
	适配器模式:将一个类的接口转换成客户希望的另外一个接口。
	适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
*/

type Player interface {
	PlayMusic()
}

func Play(player Player) {
	player.PlayMusic()
}

type MusicPlayer struct {
	Src string
}

func (music *MusicPlayer) PlayMusic() {
	fmt.Println("play music: " + music.Src)
}

type GamePlayer struct {
	Src string
}

func (game *GamePlayer) PlaySound() {
	fmt.Println("play sound: " + game.Src)
}

//通过组合的方式，声明一个Game的适配器
type GamePlayerAdapter struct {
	Game GamePlayer
}

//继承Player interface, 调用GamePlayer的方法
func (game *GamePlayerAdapter) PlayMusic() {
	game.Game.PlaySound()
}
