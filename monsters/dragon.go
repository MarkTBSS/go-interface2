package monsters

import "fmt"

type Dragon struct{}

func (d *Dragon) Attack() {
	fmt.Println("Dragon Attack")
}

func (d *Dragon) Move() {
	fmt.Println("Dragon Move")
}
