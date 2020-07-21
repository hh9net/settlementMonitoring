package utils

import "testing"

func TestStrTimeTotime(t *testing.T) {
	StrTimeTotime("1212-12-12 12:12:12")
}

func TestStrTimeToNowtime(t *testing.T) {
	StrTimeToNowtime()
}
