package main

import (
	"github.com/Nirespire/swn/content"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	pc :=  content.PlayerCharacter{}
	pc.RandomRoll(true)
	spew.Dump(pc)
}