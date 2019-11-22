package main

import "testing"

func TestCountingValleys1(t *testing.T) {
	ar := "UDDDUDUU"
	n := int32(8)

	valleys := countingValleys(n, ar)
	if valleys != 1 {
		t.Errorf("got %d instead of 1", valleys)
	}

}
