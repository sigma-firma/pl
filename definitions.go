package main

import (
	"time"

	"github.com/fatih/color"
)

// xing is a two (2) road intersection, the type of which a traffic light
// might be used to direct
type xing struct {
	Roads       []*road
	DefaultRoad *road
	ActiveRoad  *road
	Hours       []float64
}

// road is a regular road, the type you might find anywhere
type road struct {
	// ID is the roads identifier
	ID string
	// State is where we store ther state of the roads light and sensor
	// activity
	State *state
	// LeftBoundSensor is triggered when traffic triggers the sensor in the
	// left bound lane
	LeftBoundSensor bool
	// RightBoundSensor is triggered when traffic triggers the sensor in
	// the right bound lane
	RightBoundSensor bool
	// MinGreenTime is the minimum number os second a road should stay
	// green
	MinGreenTime int
	// YellowTime is the number of seconds a light should stay on yellow
	YellowTime int
	// RedTime is the number of seconds a light should stay red after
	// turning red from yellow, butbefore the green traffic light is
	// activated on the partner road
	RedTime int
	// ColorTimes helps reduce the code base a little, allowing us to map
	// a *color.Color to it's corresponding time, making for easy access
	ColorTimes map[*color.Color]int
}

// state is the stateful mechanism of a *road
type state struct {
	Color *color.Color
	Timer *time.Ticker
	Msg   string
}

var (
	// colors to be used for visual output
	red, yellow, green color.Color

	// here we set up our roads with some initial variables,
	// see: road{}
	r1 *road = &road{
		ID:           "r1",
		MinGreenTime: 3,
		YellowTime:   1,
		RedTime:      1,
	}
	r2 *road = &road{
		ID:           "r2",
		MinGreenTime: 3,
		YellowTime:   1,
		RedTime:      1,
	}
	x *xing = &xing{
		Roads:       []*road{r1, r2},
		DefaultRoad: r1,
	}
)

func (r *road) CheckSetAndReset() {
	if r.LeftBoundSensor || r.RightBoundSensor {
		<-x.ActiveRoad.SetState(&yellow).C
		<-x.ActiveRoad.SetState(&red).C
		<-r.SetState(&green).C
		r.LeftBoundSensor = false
		r.RightBoundSensor = false
		return
	}
	r.LeftBoundSensor = true
}

// SetState is used to set the state of a roads light based on the color passed
// to it, while also setting the opposing roads lights to the proper color
func (r *road) SetState(c *color.Color) *time.Ticker {
	r.State.Timer = time.NewTicker(time.Second * time.Duration(r.ColorTimes[c]))
	if c == &green {
		x.ActiveRoad.State.Msg = " INACTIVE"
		x.ActiveRoad.State.Color = &red
		r.State.Msg = "   ACTIVE"
		r.State.Color = c
		x.ActiveRoad = r
		visualize(r)
		return r.State.Timer
	}
	r.State.Color = c
	r.State.Msg = " CHANGING"
	visualize(r)
	return r.State.Timer
}
