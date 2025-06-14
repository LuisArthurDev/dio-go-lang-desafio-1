package main

import "fmt"

const ebulicaoF = 212

func main() {

	fusaoF := 32
	fusaoC := 0
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9
	fmt.Printf("A temperatura de ebulição da água em °F = %v, temperatura de ebulição da água em °C =%v.\nA temperatura de fusão da água em °F = %v, e a temperatura de fusão da água em °C é = %v.", tempF, tempC, fusaoF, fusaoC)
}
