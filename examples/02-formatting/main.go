package main

import "github.com/go-schild/logging"

func main() {
	var x = 1
	var y = 2
	var z = 3

	// This will just log a space separated string of your input.
	logging.Info("Coordinates:", x, y, z)

	// But sometimes you may want to do more formatting.
	// In this case you can use the format methods indicated by an added F at the end.
	// They use fmt.Sprintf internally.
	logging.InfoF("Coordinates: (%d, %d, %d)", x, y, z)
}
