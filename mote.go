package mote

import log "github.com/sirupsen/logrus"

// Mote represents the interface to the Pimoroni Mote
type Mote struct { }

// New initializes a new Mote interface
func New() *Mote {
	// nothing
	log.Info("Init new mote interface...\n")
	return &Mote{}
}

// TODO gamma correction
// ConfigureChannel sets the size of the channel and optionally the gamma correction
func (m *Mote) ConfigureChannel(channel, size int) {
	log.Info("Configuring chnanel %d with size %d...\n", channel, size)
}

// SetPixel sets the RGB values of a pixel on a channel
func (m *Mote) SetPixel(channel, pixel int, r, g, b uint8) {
	log.Info("Setting channel %d, pixel %d to value [r: %d, g: %d, b: %d]...\n", r, g, b, channel, pixel)
}

// Show displays the pixels on the Mote
func (m *Mote) Show() {
	log.Info("Displaying new config on Mote...")
}

// Clear sets values of all pixels to 0, 0, 0
func (m *Mote) Clear() {
	log.Info("Clearing the Mote...")
}