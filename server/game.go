package main

import (
	"encoding/json"
	"time"

	"golang.org/x/net/websocket"
)

type gameContext struct{}

type game struct {
	players []*player
	ws      *websocket.Conn
}

func newGame() *game {
	return &game{players: []*player{newPlayer("#F00"), newPlayer("#0F0"), newPlayer("#00F")}}
}

func (g *game) setWs(ws *websocket.Conn) {
	if g.ws != nil {
		g.ws.Close()
	}
	g.ws = ws
	g.start()
}

func (g *game) context() gameContext {
	return gameContext{}
}

type event struct {
	Action string `json:"action"`
	Color  string `json:"color"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func (g *game) start() {
	for {
		for _, p := range g.players {
			var ev = p.play(g.context())
			if ev.Action == "set" {
				ev.Action = "fillRect"
				bs, _ := json.Marshal([]event{ev})
				websocket.Message.Send(g.ws, string(bs))
			}
		}
		time.Sleep(1 * time.Second)
	}
}
