// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
)

func (c *Character) TakePot() {
	var usePot bool = false

	if c.CurrentHealth <= 90 /* ET SI POTION EST DANS INVENTAIRE*/ {
		usePot = true
	}

	if usePot {
		fmt.Println(c.CurrentHealth)
		switch c.CurrentHealth {
		case 90:
			c.CurrentHealth += 10
			fmt.Println("Vous avez gagné 10 HP")
			fmt.Println(c.CurrentHealth)
		case 80:
			c.CurrentHealth += 20
			fmt.Println("Vous avez gagné 20 HP")
			fmt.Println(c.CurrentHealth)
		case 70:
			c.CurrentHealth += 30
			fmt.Println("Vous avez gagné 30 HP")
			fmt.Println(c.CurrentHealth)
		case 60:
			c.CurrentHealth += 40
			fmt.Println("Vous avez gagné 40 HP")
			fmt.Println(c.CurrentHealth)
		default:
			c.CurrentHealth += 50
			fmt.Println("Vous avez gagné 50 HP")
			fmt.Println(c.CurrentHealth)
		}

		fmt.Println("Tu as utilisé ", c.Inventory[Item-1])
		c.Inventory = append(c.Inventory[:Item-1], c.Inventory[Item:]...)
		fmt.Println(c.Inventory)
	}
	usePot = false
}

func (c Character) UseItem() {

	fmt.Println("Quel item utiliser ?")

	// for i, itemPresent := range c.Inventory {
	// 	fmt.Println(i + 1)
	// 	fmt.Println(c.Inventory[1], "x")
	// }

	fmt.Scanln(&Item)

	switch Item {
	case 1:
		c.TakePot()
	case 2:
	case 3:
	case 4:

	}

	fmt.Println("Inventaire:", c.Inventory)

}

func (c Character) buyItem() {

}

func (c Character) sellItem() {

}

func (c Character) throwItem() {

}
