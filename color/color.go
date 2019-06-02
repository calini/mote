package color

import "math/rand"

// HSVToRGB TODO for now just returns random values
func HSVToRGB(h, s, v uint8) (uint8, uint8, uint8) {
	r := uint8(rand.Uint32()) % 255
	g := uint8(rand.Uint32()) % 255
	b := uint8(rand.Uint32()) % 255
	return r, g, b
}
