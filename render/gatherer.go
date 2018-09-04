package render

import (
	"fmt"
	"time"

	"github.com/wahtye/go-pathtracer/geometry"
)

type Pixel struct {
	accumulation float64
	samples      uint64
}

func NewPixel() *Pixel {
	return &Pixel{0, 0}
}

type Gatherer struct {
	width, height int
	canvas        []*Pixel
	canvasChannel chan []*Pixel
	photonChannel chan []*geometry.Photon
}

func NewGatherer(width, height int, canvasChannel chan []*Pixel, photonChannel chan []*geometry.Photon) *Gatherer {
	canvas := make([]*Pixel, width*height)
	for idx := range canvas {
		canvas[idx] = NewPixel()
	}

	return &Gatherer{
		width, height,
		canvas,
		canvasChannel,
		photonChannel,
	}
}

func (gatherer *Gatherer) Gather() {
	iterations := 0
	start := time.Now()
	total := 0
	for {
		photons := <-gatherer.photonChannel
		for _, photon := range photons {
			position := photon.Y*gatherer.width + photon.X
			gatherer.canvas[position].accumulation += photon.Intensity
			gatherer.canvas[position].samples++
			total++
		}

		iterations++
		if iterations%5 == 0 {
			raysPerSecond := int(float64(total) / time.Since(start).Seconds())
			fmt.Printf("%d rays per second, %d total\n", raysPerSecond, total)
			gatherer.canvasChannel <- gatherer.canvas
		}
	}
}
