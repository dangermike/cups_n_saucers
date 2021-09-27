package cupsaucer

import (
	"github.com/dangermike/cups_n_saucers/color"
)

// CS represents a cup and a saucer
type CS byte

const (
	// None is an unset cup/saucer
	None = CS(0)
	// SaucerWhite is a white saucer
	SaucerWhite = CS(color.White)
	// SaucerRed is a red saucer
	SaucerRed = CS(color.Red)
	// SaucerBlue is a blue saucer
	SaucerBlue = CS(color.Blue)
	// SaucerGreen is a green saucer
	SaucerGreen = CS(color.Green)
	// CupWhite is a white cup
	CupWhite = CS(color.White << 4)
	// CupRed is a red cup
	CupRed = CS(color.Red << 4)
	// CupBlue is a blue cup
	CupBlue = CS(color.Blue << 4)
	// CupGreen is a green cup
	CupGreen = CS(color.Green << 4)
)

// All is all possible combinations of cups and saucers
var All = genAllCupSaucer()

// Cup tells you the color of the cup
func (cs CS) Cup() color.C {
	return color.C(cs >> 4)
}

// Saucer tells you the color of the saucers
func (cs CS) Saucer() color.C {
	return color.C(cs) & color.All
}

// New makes a new cup and saucer pair
func New(cup, saucer color.C) CS {
	return CS((cup << 4) + saucer)
}

func (cs CS) String() string {
	return cs.Cup().String() + cs.Saucer().String()
}

func genAllCupSaucer() []CS {
	cnt := len(color.Colors) * len(color.Colors)
	b := make([]CS, cnt, cnt)
	ix := 0
	for _, cup := range color.Colors {
		for _, saucer := range color.Colors {
			b[ix] = New(cup, saucer)
			ix++
		}
	}
	return b
}
