// CHUPIN Tao -- TARDELLI Andrea
// Creation 13/09/2022
// YTrack Ynov Campus
// RedProject

package main

import "red_project/src"

func main() {
	var c1 src.Character
	c1.Init("toto", "elf", 1, 100, 40, []string{"Potion"})
	c1.DisplayInfo()
	// src.DisplayInfo()
	src.AccessInventory()
}
