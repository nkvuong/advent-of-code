package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type Coord struct {
	x float64
	y float64
	z float64
}

func strToCoord(in string) Coord {
	split := strings.Split(in, ",")
	x, _ := strconv.ParseFloat(strings.Trim(split[0], " "), 64)
	y, _ := strconv.ParseFloat(strings.Trim(split[1], " "), 64)
	z, _ := strconv.ParseFloat(strings.Trim(split[2], " "), 64)
	return Coord{x, y, z}
}

type Stone struct {
	s Coord //start
	d Coord //direction
}

const MIN_DIM = 200_000_000_000_000
const MAX_DIM = 400_000_000_000_000

func main() {
	lines := strings.Split(input, "\n")
	stones := []Stone{}
	for _, line := range lines {
		start, end, _ := strings.Cut(line, " @ ")
		stone := Stone{strToCoord(start), strToCoord(end)}
		stones = append(stones, stone)
	}
	sum := 0
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			a := stones[i]
			b := stones[j]
			// solving for as + ad * u = bs + bd * u
			dx := a.s.x - b.s.x
			dy := a.s.y - b.s.y
			det := a.d.x*b.d.y - a.d.y*b.d.x
			if det == 0 {
				//paralel lines
				continue
			}
			u := (dy*b.d.x - dx*b.d.y) / det
			v := (dy*a.d.x - dx*a.d.y) / det
			intersectX := a.s.x + a.d.x*u
			intersectY := a.s.y + a.d.y*u
			if u < 0 || v < 0 {
				//intersect in the past
				continue
			}
			if intersectX >= MIN_DIM && intersectX <= MAX_DIM && intersectY >= MIN_DIM && intersectY <= MAX_DIM {
				sum++
			}
		}
	}
	fmt.Println("Number of intersected paths are", sum)

	// part 2
	// credit to https://github.com/DeadlyRedCube/AdventOfCode/blob/main/2023/AOC2023/D24.h
	// for any given stone, A = A + Ad*t, our stone (P + Qt) should have a matching t value.
	// With three stones, we have nine equations and nine unknowns (t, u, v, Px, Py, Pz, Qx, Qy, Qz)
	// assuming that any solution for three will work for all:
	// A.s.x + A.d.x*t = Px + Qx*t
	// A.s.y + A.d.y*t = Py + Qy*t
	// A.s.z + A.d.z*t = Pz + Qz*t
	//
	// B.s.x + B.d.x*u = Px + Qx*u
	// B.s.y + B.d.y*u = Py + Qy*u
	// B.s.z + B.d.z*u = Pz + Qz*u
	//
	// C.s.x + C.d.x*v = Px + Qx*v
	// C.s.y + C.d.y*v = Py + Qy*v
	// C.s.z + C.d.z*v = Pz + Qz*v
	A := stones[0]
	B := stones[10]
	C := stones[2]
	//
	// We can eliminiate t, u, and v and end up with 6 equations with 6 unknowns (Px, Py, Pz, Qx, Qy, Qz):
	// (Px - A.s.x) / (A.d.x - Qx) = (Py - A.s.y) / (A.d.y - Qy) = (Pz - A.s.z) / (A.d.z - Qz)
	// (Px - B.s.x) / (B.d.x - Qx) = (Py - B.s.y) / (B.d.y - Qy) = (Pz - B.s.z) / (B.d.z - Qz)
	// (Px - C.s.x) / (C.d.x - Qx) = (Py - C.s.y) / (C.d.y - Qy) = (Pz - C.s.z) / (C.d.z - Qz)

	// Rearranging the Px/Py pairing:

	// Px * A.d.y - Px * Qy - A.s.x * A.d.y + A.s.x * Qy = Py * A.d.x - Py * Qx - A.s.y * A.d.x + A.s.y * Qx
	// (Px * Qy - Py * Qx) = (Px * A.d.y - Py * A.d.x) + (A.s.y * A.d.x - A.s.x * A.d.y) + (A.s.x * Qy - A.s.y * Qx)
	// (Px * Qy - Py * Qx) = (Px * B.d.y - Py * B.d.x) + (B.s.y * B.d.x - B.s.x * B.d.y) + (B.s.x * Qy - B.s.y * Qx)
	// (Px * Qy - Py * Qx) = (Px * C.d.y - Py * C.d.x) + (C.s.y * C.d.x - C.s.x * C.d.y) + (C.s.x * Qy - C.s.y * Qx)
	//
	// Note that this gets a common (Px * Qy - Py * Qx) on the left side of everything, and the right side of each is
	// now just a linear equation.
	// Do the same for the Pz/Px and Py/Pz pairings:
	//
	// (Pz * Qx - Px * Qz) = (Pz * A.d.x - Px * A.d.z) + (A.s.x * A.d.z - A.s.z * A.d.x) + (A.s.z * Qx - A.s.x * Qz)
	// (Pz * Qx - Px * Qz) = (Pz * B.d.x - Px * B.d.z) + (B.s.x * B.d.z - B.s.z * B.d.x) + (B.s.z * Qx - B.s.x * Qz)
	// (Pz * Qx - Px * Qz) = (Pz * C.d.x - Px * C.d.z) + (C.s.x * C.d.z - C.s.z * C.d.x) + (C.s.z * Qx - C.s.x * Qz)
	//
	// (Py * Qz - Pz * Qy) = (Py * A.d.z - Pz * A.d.y) + (A.s.z * A.d.y - A.s.y * A.d.z) + (A.s.y * Qz - A.s.z * Qy)
	// (Py * Qz - Pz * Qy) = (Py * B.d.z - Pz * B.d.y) + (B.s.z * B.d.y - B.s.y * B.d.z) + (B.s.y * Qz - B.s.z * Qy)
	// (Py * Qz - Pz * Qy) = (Py * C.d.z - Pz * C.d.y) + (C.s.z * C.d.y - C.s.y * C.d.z) + (C.s.y * Qz - C.s.z * Qy)
	//
	// This now turns into a series of 6 straight-up linear equations
	// [A.d.y - B.d.y]Px - [A.d.x - B.d.x]Py - [A.s.y - B.s.y]Qx + [A.s.x - B.s.x]Qy = (B.s.y * B.d.x - B.s.x * B.d.y) - (A.s.y * A.d.x - A.s.x * A.d.y)
	// [A.d.y - C.d.y]Px - [A.d.x - C.d.x]Py - [A.s.y - C.s.y]Qx + [A.s.x - C.s.x]Qy = (C.s.y * C.d.x - C.s.x * C.d.y) - (A.s.y * A.d.x - A.s.x * A.d.y)
	// [A.d.x - B.d.x]Pz - [A.d.z - B.d.z]Px - [A.s.x - B.s.x]Qz + [A.s.z - B.s.z]Qx = (B.s.x * B.d.z - B.s.z * B.d.x) - (A.s.x * A.d.z - A.s.z * A.d.x)
	// [A.d.x - C.d.x]Pz - [A.d.z - C.d.z]Px - [A.s.x - C.s.x]Qz + [A.s.z - C.s.z]Qx = (C.s.x * C.d.z - C.s.z * C.d.x) - (A.s.x * A.d.z - A.s.z * A.d.x)
	// [A.d.z - B.d.z]Py - [A.d.y - B.d.y]Pz - [A.s.z - B.s.z]Qy + [A.s.y - B.s.y]Qz = (B.s.z * B.d.y - B.s.y * B.d.z) - (A.s.z * A.d.y - A.s.y * A.d.z)
	// [A.d.z - C.d.z]Py - [A.d.y - C.d.y]Pz - [A.s.z - C.s.z]Qy + [A.s.y - C.s.y]Qz = (C.s.z * C.d.y - C.s.y * C.d.z) - (A.s.z * A.d.y - A.s.y * A.d.z)
	//
	// Combine some terms to get:
	absx := A.s.x - B.s.x
	absy := A.s.y - B.s.y
	absz := A.s.z - B.s.z

	acsx := A.s.x - C.s.x
	acsy := A.s.y - C.s.y
	acsz := A.s.z - C.s.z

	abdx := A.d.x - B.d.x
	abdy := A.d.y - B.d.y
	abdz := A.d.z - B.d.z

	acdx := A.d.x - C.d.x
	acdy := A.d.y - C.d.y
	acdz := A.d.z - C.d.z

	h0 := (B.s.y*B.d.x - B.s.x*B.d.y) - (A.s.y*A.d.x - A.s.x*A.d.y)
	h1 := (C.s.y*C.d.x - C.s.x*C.d.y) - (A.s.y*A.d.x - A.s.x*A.d.y)
	h2 := (B.s.x*B.d.z - B.s.z*B.d.x) - (A.s.x*A.d.z - A.s.z*A.d.x)
	h3 := (C.s.x*C.d.z - C.s.z*C.d.x) - (A.s.x*A.d.z - A.s.z*A.d.x)
	h4 := (B.s.z*B.d.y - B.s.y*B.d.z) - (A.s.z*A.d.y - A.s.y*A.d.z)
	h5 := (C.s.z*C.d.y - C.s.y*C.d.z) - (A.s.z*A.d.y - A.s.y*A.d.z)

	// abdy*Px - abdx*Py - absy*Qx + absx*Qy = h0
	// acdy*Px - acdx*Py - acsy*Qx + acsx*Qy = h1
	// abdx*Pz - abdz*Px - absx*Qz + absz*Qx = h2
	// acdx*Pz - acdz*Px - acsx*Qz + acsz*Qx = h3
	// abdz*Py - abdy*Pz - absz*Qy + absy*Qz = h4
	// acdz*Py - acdy*Pz - acsz*Qy + acsy*Qz = h5

	// Now we're going to take each pair of those eliminate its initial P variable (leaving just the other)
	// Okay now that's 6 linear equations and 6 variable, right?
	//
	// Px = [h0 + abdx*Py + absy*Qx - absx*Qy]/abdy
	// Px = [h1 + acdx*Py + acsy*Qx - acsx*Qy]/acdy
	//
	// [h0 + abdx*Py + absy*Qx - absx*Qy]/abdy = [h1 + acdx*Py + acsy*Qx - acsx*Qy]/acdy
	// (acdy*abdx - abdy*acdx)*Py = [abdy*acsy - acdy*absy]*Qx + [acdy*absx - abdy*acsx]*Qy + [abdy*h1 - acdy*h0]
	//
	// -----------
	// Py = ([abdy*acsy - acdy*absy]*Qx + [acdy*absx - abdy*acsx]*Qy + [abdy*h1 - acdy*h0])/(acdy*abdx - abdy*acdx)
	// -----------
	//
	// [h4 + abdy*Pz + absz*Qy - absy*Qz]/abdz = [h5 + acdy*Pz + acsz*Qy - acsy*Qz]/acdz
	// (acdz*abdy - abdz*acdy)*Pz = [abdz*acsz - acdz*absz]*Qy + [acdz*absy - abdz*acsy)*Qz + [abdz*h5 - acdz*h4]
	//
	//
	// -----------
	// Pz = ([abdz*acsz - acdz*absz]*Qy + [acdz*absy - abdz*acsy)*Qz + [abdz*h5 - acdz*h4])/(acdz*abdy - abdz*acdy)
	// -----------
	//
	// [h2 + abdz*Px + absx*Qz - absz*Qx]/abdx = [h3 + acdz*Px + acsx*Qz - acsz*Qx]/acdx
	// (acdx*abdz - abdx*acdz)*Px = [abdx*acsx - acdx*absx]*Qz + [acdx*absz - abdx*acsz]*Qx + [abdx*h3 - acdx*h2]
	//
	// -----------
	// Px = ([abdx*acsx - acdx*absx]*Qz + [acdx*absz - abdx*acsz]*Qx + [abdx*h3 - acdx*h2])/(acdx*abdz - abdx*acdz)
	// Py = ([abdy*acsy - acdy*absy]*Qx + [acdy*absx - abdy*acsx]*Qy + [abdy*h1 - acdy*h0])/(acdy*abdx - abdy*acdx)
	// Pz = ([abdz*acsz - acdz*absz]*Qy + [acdz*absy - abdz*acsy)*Qz + [abdz*h5 - acdz*h4])/(acdz*abdy - abdz*acdy)
	// -----------
	//
	// Alright, now we can sub these into (half of) our original linear equations and rearrange in terms of Qx, Qy,
	//  and Qz, leaving us with three equations and three variables.
	// abdy*Px - abdx*Py - absy*Qx + absx*Qy = h0
	// abdx*Pz - abdz*Px - absx*Qz + absz*Qx = h2
	// abdz*Py - abdy*Pz - absz*Qy + absy*Qz = h4
	//
	// Make some more variables to make this easier
	// Px = (Pxz*Qz + Pxx*Qx + Pxc)/Pxd
	// Py = (Pyx*Qx + Pyy*Qy + Pyc)/Pyd
	// Pz = (Pzy*Qy + Pzz*Qz + Pzc)/Pzd
	Pxx := acdx*absz - abdx*acsz
	Pyy := acdy*absx - abdy*acsx
	Pzz := acdz*absy - abdz*acsy

	Pxz := abdx*acsx - acdx*absx
	Pzy := abdz*acsz - acdz*absz
	Pyx := abdy*acsy - acdy*absy

	Pxc := abdx*h3 - acdx*h2
	Pyc := abdy*h1 - acdy*h0
	Pzc := abdz*h5 - acdz*h4

	Pxd := acdx*abdz - abdx*acdz
	Pyd := acdy*abdx - abdy*acdx
	Pzd := acdz*abdy - abdz*acdy

	// abdy*[(Pxz*Qz + Pxx*Qx + Pxc)/Pxd] - abdx*[(Pyx*Qx + Pyy*Qy + Pyc)/Pyd] - absy*Qx + absx*Qy = h0
	// abdx*[(Pzy*Qy + Pzz*Qz + Pzc)/Pzd] - abdz*[(Pxz*Qz + Pxx*Qx + Pxc)/Pxd] - absx*Qz + absz*Qx = h2
	// abdz*[(Pyx*Qx + Pyy*Qy + Pyc)/Pyd] - abdy*[(Pzy*Qy + Pzz*Qz + Pzc)/Pzd] - absz*Qy + absy*Qz = h4
	//
	// okay this is unintelligible garbage now but we're almost there:
	//
	// ([abdy/Pxd]*Pxz)*Qz + ([abdy/Pxd]*Pxx - [abdx/Pyd]*Pyx - absy)*Qx + (absx - [abdx/Pyd]*Pyy)*Qy
	//   = h0 - [abdy/Pxd]*Pxc + [abdx/Pyd]*Pyc
	// ([abdx/Pzd]*Pzy)*Qy + ([abdx/Pzd]*Pzz - [abdz/Pxd]*Pxz - absx)*Qz + (absz - [abdz/Pxd]*Pxx)*Qx
	//   = h2 - [abdx/Pzd]*Pzc + [abdz/Pxd]*Pxc
	// ([abdz/Pyd]*Pyx)*Qx + ([abdz/Pyd]*Pyy - [abdy/Pzd]*Pzy - absz)*Qy + (absy - [abdy/Pzd]*Pzz)*Qz
	//   = h4 - [abdz/Pyd]*Pyc + [abdy/Pzd]*Pzc
	//
	// MOAR VARIABLES
	Qz0 := (abdy / Pxd) * Pxz
	Qx0 := (abdy/Pxd)*Pxx - (abdx/Pyd)*Pyx - absy
	Qy0 := absx - (abdx/Pyd)*Pyy
	r0 := h0 - (abdy/Pxd)*Pxc + (abdx/Pyd)*Pyc

	Qy1 := (abdx / Pzd) * Pzy
	Qz1 := (abdx/Pzd)*Pzz - (abdz/Pxd)*Pxz - absx
	Qx1 := absz - (abdz/Pxd)*Pxx
	r1 := h2 - (abdx/Pzd)*Pzc + (abdz/Pxd)*Pxc

	Qx2 := (abdz / Pyd) * Pyx
	Qy2 := (abdz/Pyd)*Pyy - (abdy/Pzd)*Pzy - absz
	Qz2 := absy - (abdy/Pzd)*Pzz
	r2 := h4 - (abdz/Pyd)*Pyc + (abdy/Pzd)*Pzc

	// Qz0*Qz + Qx0*Qx + Qy0*Qy = r0
	// Qy1*Qy + Qz1*Qz + Qx1*Qx = r1
	// Qx2*Qx + Qy2*Qy + Qz2*Qz = r2
	//
	// Qx = [r0 - Qy0*Qy - Qz0*Qz]/Qx0
	//    = [r1 - Qy1*Qy - Qz1*Qz]/Qx1
	//    = [r2 - Qy2*Qy - Qz2*Qz]/Qx2
	//
	// Qx1*r0 - Qx1*Qy0*Qy - Qx1*Qz0*Qz = Qx0*r1 - Qx0*Qy1*Qy - Qx0*Qz1*Qz
	// Qy = ([Qx0*Qz1 - Qx1*Qz0]Qz + [Qx1*r0 - Qx0*r1])/[Qx1*Qy0 - Qx0*Qy1]
	//    = ([Qx0*Qz2 - Qx2*Qz0]Qz + [Qx2*r0 - Qx0*r2])/[Qx2*Qy0 - Qx0*Qy2]
	//
	// ([Qx2*Qy0 - Qx0*Qy2][Qx0*Qz1 - Qx1*Qz0] - [Qx1*Qy0 - Qx0*Qy1][Qx0*Qz2 - Qx2*Qz0])Qz
	//   = [Qx1*Qy0 - Qx0*Qy1][Qx2*r0 - Qx0*r2] - [Qx2*Qy0 - Qx0*Qy2][Qx1*r0 - Qx0*r1]

	// Alright after alllll that we can now solve for Qz, and then backsolve for everything else up the chain.
	Qz := ((Qx1*Qy0-Qx0*Qy1)*(Qx2*r0-Qx0*r2) - (Qx2*Qy0-Qx0*Qy2)*(Qx1*r0-Qx0*r1)) /
		((Qx2*Qy0-Qx0*Qy2)*(Qx0*Qz1-Qx1*Qz0) - (Qx1*Qy0-Qx0*Qy1)*(Qx0*Qz2-Qx2*Qz0))

	Qy := ((Qx0*Qz1-Qx1*Qz0)*Qz + (Qx1*r0 - Qx0*r1)) / (Qx1*Qy0 - Qx0*Qy1)

	Qx := (r0 - Qy0*Qy - Qz0*Qz) / Qx0

	Px := (Pxz*Qz + Pxx*Qx + Pxc) / Pxd
	Py := (Pyx*Qx + Pyy*Qy + Pyc) / Pyd
	Pz := (Pzy*Qy + Pzz*Qz + Pzc) / Pzd
	fmt.Printf("Sum of starting position is %.0f", math.Round(Px)+math.Round(Py)+math.Round(Pz))
}
