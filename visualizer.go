// visualizer.go houses functions used for printing useful visualizations to
// the terminal
package main

import (
	"fmt"
	"time"
)

// visualizer prints output to the terminal, allowing us to visualize the
// intersections behavior
func visualize(r *road) {
	clearTerm(50)
	var r1, r2 *road
	if r.ID == x.Roads[0].ID {
		r1 = r
		r2 = x.Roads[1]
	} else {
		r1 = x.Roads[0]
		r2 = r
	}
	c1 := r1.State.Color
	c2 := r2.State.Color
	fmt.Println(c1.Sprint(r1.ID, r1.State.Msg))
	fmt.Println(c2.Sprint(r2.ID, r2.State.Msg))
	fmt.Println("               " + c1.Sprint(" ↓ ↑ "))
	fmt.Println("               " + c1.Sprint(" rd1 "))
	count := 6
	for i := 1; i <= count*2; i++ {
		if i > 1 {
			fmt.Println("               " + c1.Sprint(" ↓ ↑ "))
		}
		if i == count {
			c2.Println("→ → → → → → → → → → → → → → → → →  ")
			c2.Println(" ← ←rd_2 ← ← ← ← ← ← ← ← ← ← ← ← ← ")
		}
	}
	time.Sleep(100 * time.Millisecond)
}

// clearTerm clears the terminal between "frames" by printing the number of
// blank lines
func clearTerm(lines int) {
	for i := 0; i <= lines; i++ {
		fmt.Println()
	}
}
