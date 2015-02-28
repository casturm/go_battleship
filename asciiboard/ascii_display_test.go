package asciiboard

import (
	"battleship/game"
	"testing"
)

type testboard struct {
	player game.Player
	result string
}

var testboards = []testboard{
	{game.Player{"", "", []*game.Ship{}, []game.Point{}},
		`~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
`},
	{game.Player{"", "", []*game.Ship{game.MakeShip(game.Point{2, 4}, "right", 5)}, []game.Point{}},
		`~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~<BBB>~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
`},
	{game.Player{"", "", []*game.Ship{&game.Ship{[]game.Point{}, []game.Point{game.Point{1, 3}, game.Point{5, 7}}}}, []game.Point{}},
		`~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~*~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~~~~~~
~~~~~*~~~~
~~~~~~~~~~
~~~~~~~~~~
`},
}

func TestAsciiDisplay(t *testing.T) {
	for i, testboard := range testboards {
		result := AsciiDisplay(testboard.player)
		expected := testboard.result
		if result != expected {
			t.Error("testboard", i, "Expected", expected, "got", result)
		}
	}
}
