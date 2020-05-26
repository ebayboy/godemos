package Sumx

import "fmt"

type Sumx struct {
	a int
	b int
}

func NewSumx(a int, b int) *Sumx {
	return &Sumx{a: a, b: b}
}

func init() {
	fmt.Println("Sumx init...")
}

func (s Sumx) Add() int {
	return s.a + s.b
}
