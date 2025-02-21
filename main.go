// This package is designed to model traffic light software, which I've reverse
// engineered in my time driving over the years. It is provided as is, with no
// warranty, under the MIT license. Yes its so good MIT gave it a license.
//
// WARNING: Not designed to run on actual traffic light hardware. Using the
// default timing will result in DEATH. You have been warned :^|
package main

import (
	"time"

	"github.com/fatih/color"
)

func main() {
	loop(1) // see: loop()
}

// loop checks the sensors of the inactive road to see if vehicles have pulled
// up to the light.
func loop(i int) {
	for ; x.Roads[i] != x.ActiveRoad; loop(1 - i) {
		x.Roads[i].CheckSetAndReset()
	}
}

// init is used to initialize/tune startup variables.
func init() {
	// fine tuning the colors
	red.AddBgRGB(255, 118, 97).AddRGB(255, 255, 255)
	yellow.AddBgRGB(255, 255, 0).AddRGB(102, 102, 0)
	green.AddBgRGB(182, 212, 108).AddRGB(255, 255, 255)

	// set the xing.ActiveRoad road to the xing.DefaultRoad
	x.ActiveRoad = x.DefaultRoad

	// set up the ColorTime map, which maps *color.Color to int for use in
	// our timing mechanisms
	for _, r := range x.Roads {
		r.ColorTimes = make(map[*color.Color]int)
		r.ColorTimes[&red] = r.RedTime
		r.ColorTimes[&yellow] = r.YellowTime
		r.ColorTimes[&green] = r.MinGreenTime

		// set the initial state of each road
		r.State = &state{
			Color: &red,                // red
			Timer: time.NewTicker(100), // placeholder, does nothing
			Msg:   " INACTIVE",
		}
	}
}
