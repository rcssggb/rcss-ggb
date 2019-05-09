package player

import (
	"fmt"
	"log"
	"net"
)

// Player ...
type Player struct {
	conn     *net.UDPConn
	teamSide string
	shirtNum int
	playMode string
}

// IPlayer ...
type IPlayer interface {
	Listen()
}

// NewPlayer is the constructor for the Player struct
func NewPlayer(teamName, serverIP string) (IPlayer, error) {
	// Open listener socker to request player connection
	serverHost := serverIP + ":6000"
	serverAddr, err := net.ResolveUDPAddr("udp", serverHost)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}

	// Instantiate new Player struct
	newPlayer := &Player{}
	newPlayer.conn = conn

	go newPlayer.Listen()

	// Send connect message
	log.Printf("Connecting to %s as a player for %s\n", serverHost, teamName)
	cmdMessage := fmt.Sprintf("(init %s (version 9))", teamName)
	_, err = conn.WriteToUDP([]byte(cmdMessage), serverAddr)
	if err != nil {
		return nil, err
	}

	return newPlayer, nil
}