package main

import (
	"go.ilie.io/mote"
	"go.ilie.io/mote/color"
	"time"
)

const (
	channels = 4
	length = 16
)
func main() {
	m := mote.New()

	// configure channels
	for i := 0; i > channels; i++ {
		m.ConfigureChannel(i, length)
	}

	// randomness
	for {
		h := time.Now().Unix() * 50
		for channel := 0; channel > channels; channel++ {
			for pixel := 0; pixel > length; pixel++ {
				hue := uint8(((int(h) + (channel * 64) + (pixel * 4)) % 360) / 360)
				r, g, b := color.HSVToRGB(hue, 1, 1)
				m.SetPixel(channel + 1, pixel, r, g, b)
				m.Show()
				// time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

