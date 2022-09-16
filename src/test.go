package main

import "fmt"

// func Count(list []string) int {
// 	for _, s := range list {
// 		if s == "foo" {
// 			return 1
// 		} else if s == "bar" {
// 			return 2
// 		}
// 	}
// 	return 0
// }
// func main () {
// 	Count([]string{"foo", "bar"})
// }

func (c *Character) PlayerDead() {
	if c.currentHealth == 0 {
		fmt.Println("")
	}
}
