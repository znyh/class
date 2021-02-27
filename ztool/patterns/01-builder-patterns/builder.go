package builder

/*
	建造者模式：将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示

	设计思想：
		*Builder interface (包含1.biuld_XX method 返回的是biulder接口，2.get_XX 返回对象)
		*父struct
		*Director struct, 属性为Builder, 实现Construct()和SetBuilder()方法
		*不同的子struct组合，实现接口builder
*/

type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	GetVehicle() Vehicle
}

type Vehicle struct {
	Wheels int
	Seats  int
}

type Car struct {
	vehicle Vehicle
}

func (car *Car) SetWheels() Builder {
	car.vehicle.Wheels = 4
	return car
}

func (car *Car) SetSeats() Builder {
	car.vehicle.Seats = 4
	return car
}

func (car *Car) GetVehicle() Vehicle {
	return car.vehicle
}

type Director struct {
	builder Builder
}

func (director *Director) Construct(b Builder) {
	b.SetWheels().SetSeats() //链式调用
	director.builder = b
}
