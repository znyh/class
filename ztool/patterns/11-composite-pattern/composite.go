package composite

import "fmt"

/*
	组合模式有助于表达数据结构, 将对象组合成树形结构以表示"部分-整体"的层次结构, 常用于树状的结构
*/

//抽离共同属性部分
type MenuComponent interface {
	Price() float32
	Print()
}

type MenuDesc struct {
	name        string
	description string
}

func (desc *MenuDesc) Name() string {
	return desc.name
}
func (desc *MenuDesc) Description() string {
	return desc.description
}

type MenuItem struct {
	MenuDesc
	price float32
}

func NewMenuItem(name, description string, price float32) *MenuItem {
	return &MenuItem{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
		price: price,
	}
}

func (item *MenuItem) Price() float32 {
	return item.price
}
func (item *MenuItem) Print() {
	fmt.Printf("	%s, %0.2f\n", item.name, item.price)
	fmt.Printf("	-- %s\n", item.description)
}

type MenuGroup struct {
	child []MenuComponent
}

func (group *MenuGroup) Add(component MenuComponent) {
	group.child = append(group.child, component)
}
func (group *MenuGroup) Remove(id int) {
	group.child = append(group.child[:id], group.child[id+1:]...)
}
func (group *MenuGroup) Find(id int) MenuComponent {
	return group.child[id]
}

type Menu struct {
	MenuDesc
	MenuGroup
}

func NewMenu(name, description string) *Menu {
	return &Menu{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
	}
}

func (m *Menu) Price() (price float32) {
	for _, v := range m.child {
		price += v.Price()
	}
	return price
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s, ¥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.child {
		v.Print()
	}
	fmt.Println("结束")
}
