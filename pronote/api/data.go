package api

// Subject defines the pretty-print style of a school subject
type Subject struct {
	Name  string
	Emoji string
}

// Subjects holds the list of all available subjects in Pronote
var Subjects = map[string]Subject{
	// Core curriculum
	"HISTOIRE & GEOGRAPHIE":             {Name: "Histoire-Géographie", Emoji: "🌎"},
	"ENS. MORAL & CIVIQUE":              {Name: "EMC", Emoji: "🏛"},
	"SC PHYSIQ & CHIMIQ":                {Name: "Sciences physiques", Emoji: "🔭"},
	"SCIENCES DE LA VIE ET DE LA TERRE": {Name: "SVT", Emoji: "☘"},
	"ED.PHYSIQUE & SPORT.":              {Name: "EPS", Emoji: "⚽"},
	"PHILOSOPHIE":                       {Name: "Philosophie", Emoji: "✒"},
	"PARCOURS REUSSITE ORIENT":          {Name: "MAP PRO (Vie de classe)", Emoji: "🪑"},

	// Living languages
	"ANGLAIS":            {Name: "Anglais", Emoji: "🍵"},
	"ESPAGNOL":           {Name: "Espagnol", Emoji: "🌮"},
	"DNL SI":             {Name: "Anglais Euro", Emoji: "🇪🇺"},
	"ANGLAIS SECT.EUROP": {Name: "Anglais Euro", Emoji: "🇪🇺"},

	// Specialties and options
	"MATHEMATIQUES":        {Name: "Mathématiques", Emoji: "🔢"},
	"MATHS EXP":            {Name: "Maths expertes", Emoji: "🧮"},
	"SC.INGEN. & SC.PHYS.": {Name: "Sciences de l'ingénieur", Emoji: "⚙"},
}

type Clock struct {
	Hour   int
	Minute int
	Emoji  rune
}

var Hours = []Clock{
	// 🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚🕛🕜🕝🕞🕟🕠🕡🕢🕣🕤🕥🕦🕧
	{Hour: 1, Minute: 0, Emoji: '🕐'},
	{Hour: 2, Minute: 0, Emoji: '🕐'},
	{Hour: 3, Minute: 0, Emoji: '🕐'},
	{Hour: 4, Minute: 0, Emoji: '🕐'},
	{Hour: 5, Minute: 0, Emoji: '🕐'},
	{Hour: 6, Minute: 0, Emoji: '🕐'},
	{Hour: 7, Minute: 0, Emoji: '🕐'},
	{Hour: 8, Minute: 0, Emoji: '🕐'},
	{Hour: 9, Minute: 0, Emoji: '🕐'},
	{Hour: 10, Minute: 0, Emoji: '🕐'},
	{Hour: 11, Minute: 0, Emoji: '🕐'},
	{Hour: 12, Minute: 0, Emoji: '🕐'},
}
