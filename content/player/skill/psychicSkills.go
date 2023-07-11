package skill

type PsychicSkill int

type Biopsionics PsychicSkill
type Metapsionics PsychicSkill
type Precognition PsychicSkill
type Telekinesis PsychicSkill
type Telepathy PsychicSkill
type Teleportation PsychicSkill

type PsychicSkills struct {
	Biopsionics   Biopsionics
	Metapsionics  Metapsionics
	Precognition  Precognition
	Telekinesis   Telekinesis
	Telepathy     Telepathy
	Teleportation Teleportation
}
