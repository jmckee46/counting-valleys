package main

func countingValleys(n int32, s string) int32 {
	var altitude int32
	var previousAltitude int32
	var valleys int32

	for _, char := range s {
		previousAltitude = altitude

		if string(char) == "U" {
			altitude++
		} else {
			altitude--
		}
		if previousAltitude == 0 && altitude < 0 {
			valleys++
		}
	}
	return valleys
}

func main() {

}
