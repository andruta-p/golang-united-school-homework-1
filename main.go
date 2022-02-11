package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	fmt.Println(CreateStringWithEmoji())
}

func CreateStringWithEmoji() string {
	return emoji.Sprint("Hello :world_map:!")
}
