package flyweight

/*
	享元模式:把多个实例对象共同需要的数据，独立出一个享元，从而减少对象数量和节省内存
	经典案例：数据库连接池，避免重新开启数据库链接的开销
*/

type Shape interface {
	SetRadius(radius int)
	SetColor(color string)
}

type Circle struct {
	color  string
	radius int
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

type ShapeFactory struct {
	circleMap map[string]Shape
}

func (sh *ShapeFactory) GetCircle(color string) Shape {
	if sh.circleMap == nil {
		sh.circleMap = make(map[string]Shape)
	}
	if shape, ok := sh.circleMap[color]; ok {
		return shape
	}
	circle := new(Circle)
	circle.SetColor(color)
	sh.circleMap[color] = circle
	return circle
}
