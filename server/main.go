package main

import (
	"log"
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	mustServe(mustNewListener(9000), newGame())
}

func mustServe(listener *net.TCPListener, g *game) {
	mux := http.NewServeMux()
	mux.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) { g.setWs(ws) }))

	if err := http.Serve(listener, mux); err != nil {
		log.Println("Server went down: ", err)
	}
}

func mustNewListener(port int) *net.TCPListener {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP("localhost").To4(),
		Port: port,
	})
	if err != nil {
		log.Fatalf("Could not open listener on port %v. err=%v", port, err)
	}
	return listener
}
