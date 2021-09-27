package color

// C is a color
type C int16

const (
	// None is an unset color
	None C = 0

	// White with fear
	White C = 1
	// Red for the your ruby lips
	Red C = 2
	// Blue for Miles Davis
	Blue C = 4
	// Green for Mother Earth Gaia
	Green C = 8
	// All for the rainbow
	All = White | Red | Blue | Green
)

// Colors  is a list of all the colors we know
var Colors = []C{White, Red, Blue, Green}

func (c C) String() string {
	if c == Red {
		return "R"
	}
	if c == Blue {
		return "B"
	}
	if c == Green {
		return "G"
	}
	if c == White {
		return "W"
	}
	return "?"
}
