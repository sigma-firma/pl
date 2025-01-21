package main

import (
	"time"

	"github.com/fatih/color"
)

func main() {
	loop(1) // magic, see: loop()
}

// loop checks the sensors of the inactive road to see if vehicles have pulled
// up to the light
func loop(i int) {
	for ; x.Roads[i] != x.ActiveRoad; loop(1 - i) {
		x.Roads[i].CheckSetAndReset()
	}
}

// init is used to initialize/tune some variables
func init() {
	// fine tuning the colors
	red.AddBgRGB(255, 118, 97).AddRGB(255, 255, 255)
	yellow.AddBgRGB(255, 255, 0).AddRGB(102, 102, 0)
	green.AddBgRGB(182, 212, 108).AddRGB(255, 255, 255)

	// set the default road to the xing.ActiveRoad
	x.ActiveRoad = x.DefaultRoad

	// set up the ColorTime map, which maps *color.Color to int to use in
	// our timing mechanisms
	for _, r := range x.Roads {
		r.ColorTimes = make(map[*color.Color]int)
		r.ColorTimes[&red] = r.RedTime
		r.ColorTimes[&yellow] = r.YellowTime
		r.ColorTimes[&green] = r.MinGreenTime
		// set up the initial state of each roads light to red
		r.State = &state{
			Color: &red,
			Timer: time.NewTicker(100),
			Msg:   " INACTIVE",
		}
	}
}
