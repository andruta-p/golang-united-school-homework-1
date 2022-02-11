package main

import (
	"testing"
)

func TestCreateStringWithEmoji(t *testing.T) {
	t.Log("Check output")
	output := CreateStringWithEmoji()
	if output != "Hello 🗺️ !" {
		t.Error("Wrong output")
	}
}
