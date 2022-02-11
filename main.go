package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

// Main method
func main() {
	fmt.Println(CreateStringWithEmoji())
}

// Method for creating text "Hello" with world map emoji
func CreateStringWithEmoji() string {
	return emoji.Sprint("Hello :world_map:!")
}
