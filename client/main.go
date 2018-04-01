package main

import (
	"encoding/json"
	"log"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	var canvas = js.Global.Get("document").Call("createElement", "canvas")
	canvas.Set("width", 500)
	canvas.Set("height", 500)
	canvas.Set("style", "border: 1px solid black")
	js.Global.Get("document").Get("body").Call("appendChild", canvas)
	var ctx = canvas.Call("getContext", "2d")
	var g = &game{ctx}
	js.Global.Set("run", g.run)
}

type game struct {
	ctx *js.Object
}

type event struct {
	Action string `json:"action"`
	Color  string `json:"color"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func (g *game) run(s string) {
	var events = make([]event, 0)
	var err = json.Unmarshal([]byte(s), &events)
	if err != nil {
		log.Fatal(err)
	}

	g.ctx.Call("save")
	for _, e := range events {
		switch e.Action {
		case "fillRect":
			g.ctx.Set("fillStyle", e.Color)
			g.ctx.Call("fillRect", e.X, e.Y, e.X+100, e.Y+100)
		}
	}
	g.ctx.Call("restore")
}
