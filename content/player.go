package content

import (
	"log"

	"github.com/Nirespire/swn/util"
	"github.com/justinian/dice"
)

type Background struct {
	name string
}

type Cyberware struct {
	name string
}

type Focus struct {
	name string
}

type Gear struct {
	name string
}

type Weapon struct {
	name string
}

type Stat struct {
	value    int
	modifier int
}

type Stats struct {
	strength     Stat
	dexterity    Stat
	constitution Stat
	intelligence Stat
	wisdom       Stat
	charisma     Stat
}

type Health struct {
	current int
	max     int
}

type Skill struct {
	name       string
	decription string
	level      int
}

type PsychicSkill struct {
	name        string
	description string
	level       int
}

type Effort struct {
	current int
	active  int
	max     int
}

type SavingThrows struct {
	PhysicalEffect int
	MentalEffect   int
	Evasion        int
}

type PlayerCharacter struct {
	Name          string
	Level         int
	Stats         Stats
	Homeworld     string
	Background    Background
	Faction       string
	Foci          []Focus
	Cyberware     Cyberware
	Gear          []Gear
	Weapons       []Weapon
	Skills        []Skill
	PsychicSkills []PsychicSkill
	Credits       int
	Debt          int
	Effort        Effort
	SavingThrows  SavingThrows
}

type PlayerCharacterGenerator interface {
	RandomRoll()
}

func (pc *PlayerCharacter) RandomRoll(pickArray bool) {

	var arrayStats []int

	// 1
	if pickArray {
		log.Print("Using array stats")
		arrayStats = util.Shuffle([]int{14, 12, 11, 10, 9, 7})
	} else {
		log.Print("Using rolled stats")
		arrayStats = make([]int, 5)

		for i := range arrayStats {
			rollResult, _, _ := dice.Roll("3d6")
			arrayStats[i] = rollResult.Int()
		}
	}

	pc.Stats.strength = statsAndMod(arrayStats[0])
	pc.Stats.dexterity = statsAndMod(arrayStats[1])
	pc.Stats.constitution = statsAndMod(arrayStats[2])
	pc.Stats.intelligence = statsAndMod(arrayStats[3])
	pc.Stats.wisdom = statsAndMod(arrayStats[4])
	pc.Stats.charisma = statsAndMod(arrayStats[5])

}

func statsAndMod(value int) Stat {
	stat := Stat{}

	stat.value = value

	if value <= 3 {
		stat.modifier = -2
	} else if value >= 4 && value <= 7 {
		stat.modifier = -1
	} else if value >= 8 && value <= 13 {
		stat.modifier = 0
	} else if value >= 14 && value <= 17 {
		stat.modifier = 1
	} else if value >= 18 {
		stat.modifier = 2
	}

	return stat
}
