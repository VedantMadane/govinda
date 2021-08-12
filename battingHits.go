
// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"
// Declare a struct that represents a ball player.
type player struct{
	name string
	atBats int
	hits int
}
// Include fields called name, atBats and hits.

// Declare a method that calculates the batting average for a player.
func (p *player) average() float64 {
	return float64(p.hits)/float64(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{"Vedant", 10, 7},
		{"Kaustubh", 12, 6},
		{"Siddesh", 6, 4},
	}

	// Display the batting average for each player in the slice.
	for _, p := range players{
		fmt.Printf("%s: Average [.%.f]\n", p.name, p.average()*100)
	}
}
