package npc

import (
	"log"

	"github.com/Nirespire/swn/util"
)

type NPC struct {
	Background     string
	RoleInSociety  string
	BiggestProblem string
	Age            string
	Desire         string
	CharacterTrait string
}

var backgroundsList = []string{
	"The local underclass or porest natives",
	"Common laborers or cube workers",
	"Aspiring bourgeoise or upper class",
	"The elite of this society",
	"Minority or foreigners",
	"Offworlders or exotics",
}

var roleInSocietyList = []string{
	"Criminal, thug, thief, swindler",
	"Menial, cleaner, retail worker, servant",
	"Unskilled heavy labor, porter, construction",
	"Skilled trade, electrician, mechanic, pilot",
	"Idea worker, programmer, writer",
	"Merchant, business owner, trader, banker",
	"Official, bureaucrat, courtier, clerk",
	"Military, soldier, enforcer, law officer",
}

var biggestProblemList = []string{
	"They have significant debt or money woes",
	"A loved one is in trouble; rerolled for it",
	"Romantic failure with a desired person",
	"Drug or behavioral addiction",
	"Their superior dislikes or resents them",
	"They have a persistent sickness",
	"They hate their job or live situation",
	"Someone dangerous is targeting them",
	"They're pursuing a disasterous purpose",
	"They have no problems worth mentioning",
}

var ageList = []string{
	"Unusually young or old for their role",
	"Young adult",
	"Mature prime",
	"Middle-aged or elderly",
}

var greatestDesiredList = []string{
	"They want a particular romantic partner",
	"They want money for them or a loved one",
	"They want a promotion in their job",
	"They want answers about a past trauma",
	"They want revenge on an enemy",
	"They want to help a beleaguered friend",
	"They want an entirely different job",
	"They want protection from an enemy",
	"They want to leave their current life",
	"They want fame and glory",
	"They want power over those around them",
	"They have everything they want from life",
}

var chacterTraitList = []string{
	"Ambition",
	"Avarice",
	"Bitterness",
	"Courage",
	"Cowardice",
	"Curiosity",
	"Deceitfulness",
	"Determination",
	"Devotion to a cause",
	"Filiality",
	"Hatred",
	"Honesty",
	"Hopefulness",
	"Love of a person",
	"Nihilism",
	"Paternalism",
	"Pessimism",
	"Protectiveness",
	"Resentment",
	"Shame",
}

func (npc *NPC) RandomRoll() {
	npc.Background = util.GetRandomArrayString(backgroundsList)
	npc.RoleInSociety = util.GetRandomArrayString(roleInSocietyList)
	npc.BiggestProblem = util.GetRandomArrayString(biggestProblemList)
	npc.Age = util.GetRandomArrayString(ageList)
	npc.Desire = util.GetRandomArrayString(greatestDesiredList)
	npc.CharacterTrait = util.GetRandomArrayString(chacterTraitList)
}

func (npc *NPC) GenerateImagePrompt() string {
	portraitPrompt := "Portrait of a " + npc.Age + ". Their background is " + npc.Background + " and their role in society is " + npc.RoleInSociety + ". Their expression is one of " + npc.CharacterTrait +
		". Realistic. Sci-fi style."
	log.Println(portraitPrompt)
	return portraitPrompt
}
