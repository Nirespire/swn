package main

import (
	"github.com/Nirespire/swn/content"
	"github.com/Nirespire/swn/content/name"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	pc :=  content.PlayerCharacter{}
	pc.RandomRoll(true, name.Any, name.None)
	spew.Dump(pc)
}