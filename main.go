package main

import "fmt"

type Slime struct{}
type dragon struct{}

// Parameters and Return must be the same
type ActionHandler interface {
	Attack()
	Move()
}

func NewSlimeAction() ActionHandler {
	return &Slime{}
}

func NewDragonAction() ActionHandler {
	return &dragon{}
}

func (s *Slime) Attack() {
	fmt.Println("Slime Attack")
}

func (s *Slime) Move() {
	fmt.Println("Slime Move")
}

func (s *Slime) Hidden() {
	fmt.Println("Slime Hidden")
}

func (d *dragon) Attack() {
	fmt.Println("Dragon Attack")
}

func (d *dragon) Move() {
	fmt.Println("Dragon Move")
}

func main() {
	slime := NewSlimeAction()
	slime.Attack()
	slime.Move()
	//slime.Hidden()

	dragon := NewDragonAction()
	dragon.Attack()
	dragon.Move()
}
