package dao

import "time"

type B struct {
	Name string
}

func NewB() *B {
	return &B{
		Name: time.Now().String(),
	}
}
