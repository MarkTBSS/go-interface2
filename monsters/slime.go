package monsters

import "fmt"

type Slime struct{}

func (s *Slime) Attack() {
	fmt.Println("Slime Attack")
}

func (s *Slime) Move() {
	fmt.Println("Slime Move")
}

func (s *Slime) Hidden() {
	fmt.Println("Slime Hidden")
}
