package background

type Background struct {
	name        string
	description string
}

var BackgroundsList = []Background{
	{"Barbarian", "born of a primitive world"},
	{"Clergy", "a consecreated man or woman"},
	{"Courtesan", "trading on pleasurable company"},
	{"Criminal", "thief, rogue, liar, or worse 6 Entertainer, artful and beguiling"},
	{"Noble", "by blood or by social capital"},
	{"Peasant", "whether primitive or high-tech 12 Pilot, or rider, or sailor, or vehicle-driver 14 Scholar, a scientist or academic"},
	{"Spacer", "dwelling in the deep-space habs 18 Thug, ruffian, or strong arm of the people"},
	{"Worker", "a cube drone or day laborer"},
	{"Dilettante", "with money if not purpose"},
	{"Merchant", "whether peddler or far trader"},
	{"Official", "a functionary of some greater state"},
	{"Physician", "a healer of the sick and maimed"},
	{"Politician", "aspiring to leadership and control"},
	{"Soldier", "whether mercenary or conscript"},
	{"Technician", "artisan, engineer, or builder"},
	{"Vagabond", "roaming without a home"},
}
