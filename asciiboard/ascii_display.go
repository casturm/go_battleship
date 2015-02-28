package asciiboard

import (
	"battleship/game"
	"bytes"
)

func AsciiDisplay(player game.Player) string {
	var buffer bytes.Buffer

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			ascii := "~"
			for _, ship := range player.Ships {
				for p, point := range ship.Location {
					if (point == game.Point{x, y}) {
						switch p {
						case 0:
							ascii = "<"
						case len(ship.Location) - 1:
							ascii = ">"
						default:
							ascii = "B"
						}
					}
				}
				for _, point := range ship.Hits {
					if (point == game.Point{x, y}) {
						ascii = "*"
					}
				}
			}
			buffer.WriteString(ascii)
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
