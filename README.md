# Purple-Light (pl) - Generic Traffic Light Software

A model program supporting all of the traffic light functionality that our company 
"reverse engineered" over time spent on the road and observing the behavior of 
different traffic lights at different intersections across the United States.

[Screen Recording 2025-01-21 at 12.08.20 PM.webm](https://github.com/user-attachments/assets/9d12d4ed-c597-4636-b5be-3908bda38682)

## Colors Planned for Support by Feb 2025
 - Red
 - Yellow
 - Green

Support may be added for a pink or purple light. 

## Planned Functionality
 - Time-of-day specific configurations
 - Sensor/timing based configurations
 - Left Turns
 - Caution/Stop Lights
 - Stop Lights
 - Generic 2-road intersections only (for now)

## Running the Software 

    // cd to your home dir, clone the repo, and cd into the repo dir
    cd && git clone https://github.com/hartsfield/pl && cd pl

    // install dependencies (oh yeah you need the Go programming language)
    go mod tidy

    // run
    go run .

    // remember: lives are at stake here, take this seriously

## About the Software

We implement golangs <-time.Ticker.C (basically countdown channels) for timing
the lights so that we don't cause collisions.
