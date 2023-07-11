package main

import (
	"log"

	"github.com/Nirespire/swn/content"
	"github.com/Nirespire/swn/content/npc"
	"github.com/Nirespire/swn/content/player/name"
	"github.com/Nirespire/swn/modules/gpt"
	"github.com/alecthomas/kong"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	var CLI struct {
		Npc struct {
			Portrait bool `help:"Generate a portrait image" default:true`
			Stats    bool `help:"Generates a stat block for the NPC" default:false`
		} `cmd:"" help:"Generates content for NPCs"`
		Player struct {
		} `cmd:"" help:"Generates content for Player Characters"`
	}

	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "npc":
		npc := npc.NPC{}
		npc.RandomRoll()
		spew.Dump(npc)

		req := gpt.ImageRequest{npc.GenerateImagePrompt()}
		imageURL := req.GenerateImage()
		log.Println(imageURL)

	case "player":
		pc := content.PlayerCharacter{}
		pc.RandomRoll(true, name.Any, name.None)
		spew.Dump(pc)
	default:
		panic(ctx.Command())
	}
}
