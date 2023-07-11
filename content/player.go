package content

import (
	"log"

	"github.com/Nirespire/swn/content/player/background"
	"github.com/Nirespire/swn/content/player/name"
	"github.com/Nirespire/swn/content/player/skill"
	"github.com/Nirespire/swn/util"
	"github.com/justinian/dice"
)

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
	Background    background.Background
	Faction       string
	Foci          []Focus
	Cyberware     Cyberware
	Gear          []Gear
	Weapons       []Weapon
	Skills        skill.Skills
	PsychicSkills skill.PsychicSkills
	Credits       int
	Debt          int
	Effort        Effort
	SavingThrows  SavingThrows
}

func (pc *PlayerCharacter) RandomRoll(pickArray bool, culture name.Culture, gender name.Gender) {

	// 1 - Roll your six attributes or assign them from an array
	// 2 - Mark down your attribute modifiers for each score.
	pc.Stats = generateStats(pickArray)
	
	// 3 - Pick a Background
	pc.Background = generateBackground()
	
	// misc

	// Pick Name
	pc.Name = name.GetNameByCultureGender(culture, gender)
}

func generateBackground() background.Background {
	return background.BackgroundsList[util.GetNewRandom().Intn(len(background.BackgroundsList))]
}

func generateStats(pickArray bool) Stats {
	stats := Stats{}

	var arrayStats []int

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

	stats.strength = generateStatsAndMod(arrayStats[0])
	stats.dexterity = generateStatsAndMod(arrayStats[1])
	stats.constitution = generateStatsAndMod(arrayStats[2])
	stats.intelligence = generateStatsAndMod(arrayStats[3])
	stats.wisdom = generateStatsAndMod(arrayStats[4])
	stats.charisma = generateStatsAndMod(arrayStats[5])

	return stats
}

func generateStatsAndMod(value int) Stat {
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
