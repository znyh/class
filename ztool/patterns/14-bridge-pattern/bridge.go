package bridge

import "fmt"

/*
	桥接模式分离抽象部分和实现部分，使得两部分可以独立扩展
*/

//创建桥接接口
type SoftWare interface {
	Run()
}

type Cpu struct {
}

func (c *Cpu) Run() {
	fmt.Println("this is cpu run")
}

type Storage struct {
}

func (s *Storage) Run() {
	fmt.Println("this is storage run")
}

type Shape struct {
	software SoftWare //创建Shape struct（抽象部分）
}

func (s *Shape) SetSoftWare(soft SoftWare) {
	s.software = soft
}

type Phone struct {
	shape Shape
}

func (p *Phone) SetShape(soft SoftWare) {
	p.shape.SetSoftWare(soft)
}

func (p *Phone) Print() {
	p.shape.software.Run()
}
