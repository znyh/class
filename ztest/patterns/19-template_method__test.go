package patterns

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {

	base := new(component)
	temp := new(template)
	temp.component = base

	base.setname("abc")
	fmt.Println("base name:", base.getname())
	fmt.Println("sub name:", temp.getname())

	temp.setname("123")
	fmt.Println("base name:", base.getname())
	fmt.Println("sub name:", temp.getname())

}
