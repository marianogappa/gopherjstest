package main

import "math/rand"

type player struct {
	color string
}

func newPlayer(color string) *player {
	return &player{color: color}
}

func (p *player) play(gc gameContext) event {
	return event{Action: "set", Color: p.color, X: rand.Intn(400), Y: rand.Intn(400)}
}
