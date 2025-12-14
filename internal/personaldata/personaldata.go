package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %f.2\nРост: %f.2\n", p.Name, p.Weight, p.Height)
}
