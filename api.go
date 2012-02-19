package main

// The API wraps a player and only exports
// methods meant to be called by the user.
type API struct {
	player *Player
}
