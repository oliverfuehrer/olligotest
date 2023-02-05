package ollipackage

import "strconv"

var sum int = 0

func internalText() string {
	sum++
	return "helloText: " + strconv.Itoa(sum)
}
