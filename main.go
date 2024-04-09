package main

import "github.com/MarkTBSS/go-interface2/monsters"

type Actionner interface {
	Attack()
	Move()
}

func NewSlime() Actionner {
	return &monsters.Slime{}
}

func NewDragon() Actionner {
	return &monsters.Dragon{}
}

func main() {
	slime := NewSlime()
	slime.Attack()
	slime.Move()
	//slime.Hidden()

	dragon := NewDragon()
	dragon.Attack()
	dragon.Move()
}
