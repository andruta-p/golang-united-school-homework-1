package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Log("Check output")
	output := CreateStringWithEmoji()
	if output != "Hello 🗺️ !" {
		t.Error("Wrong output")
	}
}
