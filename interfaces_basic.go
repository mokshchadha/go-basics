package main

import "fmt"

type bot interface {
	getGretting() string // what different functions and return type can a bot have
}

type englishBot struct {
}

type spanishBot struct {
}

func (englishBot) getGretting() string {
	return "Hello !!"
}

func (spanishBot) getGretting() string {
	return "Hola !!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGretting())
}
