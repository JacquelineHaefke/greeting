package main

import "fmt"

type person struct {
	name string
	age  int
}

type child struct{}

func (p person) greet() {
	fmt.Println("Hello person " + p.name + "!")
}

func (d child) greet() {
	fmt.Println("Hello child!")
}
