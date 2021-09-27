package board

import (
	"strings"

	"github.com/dangermike/cups_n_saucers/cupsaucer"
)

// B is where we put our cups and saucers
type B [16]cupsaucer.CS

func (b B) String() string {
	var sb strings.Builder
	sb.Grow(16 * 3)
	for r := 0; r < 16; r += 4 {
		if r > 0 {
			sb.WriteRune('\n')
		}
		for c := 0; c < 4; c++ {
			if c > 0 {
				sb.WriteRune(' ')
			}
			sb.WriteString(b[r+c].String())
		}
	}
	return sb.String()
}

// IsValid returns if the board is in a valid state
func (b B) IsValid() bool {
	// we don't have to check for CupSaucer items because
	// we are going to try to place each item once

	// check across rows
	for r := 0; r < 16; r += 4 {
		acc := cupsaucer.CS(0) // blank
		for c := 0; c < 4; c++ {
			if acc&b[r+c] > 0 {
				return false
			}
			acc |= b[r+c]
		}
	}

	// check down columns
	for c := 0; c < 4; c++ {
		acc := cupsaucer.CS(0) // blank
		for r := 0; r < 16; r += 4 {
			if acc&b[r+c] > 0 {
				return false
			}
			acc |= b[r+c]
		}
	}

	return true
}

// TryPlace attempts to put a cup/saucer onto the board. Returns whether
// that placement is valid
func (b B) TryPlace(cs cupsaucer.CS, position int) (B, bool) {
	if b[position] > cupsaucer.None {
		return b, false
	}
	b[position] = cs
	return b, b.IsValid()
}

// Equal tests if two boards are identical
func Equal(a, b B) bool {
	a1, a2 := a.toUints()
	b1, b2 := b.toUints()
	return a1 == b1 && a2 == b2
}

func (b B) toUints() (uint64, uint64) {
	l := uint64(b[0]) +
		uint64(b[1])<<8 +
		uint64(b[2])<<16 +
		uint64(b[3])<<24 +
		uint64(b[4])<<32 +
		uint64(b[5])<<40 +
		uint64(b[6])<<48 +
		uint64(b[7])<<56
	h := uint64(b[8]) +
		uint64(b[9])<<8 +
		uint64(b[10])<<16 +
		uint64(b[11])<<24 +
		uint64(b[12])<<32 +
		uint64(b[13])<<40 +
		uint64(b[14])<<48 +
		uint64(b[15])<<56
	return l, h

}

// Merge merges a board of saucers with a board of cups
func Merge(a, b B) B {
	var r B
	for x := 0; x < 16; x++ {
		r[x] = a[x] | b[x]
	}
	return r
}
