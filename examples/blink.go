package main

import (
	"go.ilie.io/mote"
	"time"
)

func main() {
	m := mote.New()

	m.ConfigureChannel(1, 16)

	for {
		m.SetPixel(1, 1, 255, 255, 255) // set pixel 1, 1 to white
		m.Show()
		time.Sleep(1 * time.Second)

		m.Clear()
		m.Show()
		time.Sleep(1 * time.Second)
	}
}

