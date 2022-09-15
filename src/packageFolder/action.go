package packageFolder

import "fmt"

func (c *Character) TakePot() {
	var usePot bool = false
	var item int

	if c.CurrentHealth <= 90 /* ET SI POTION EST DANS INVENTAIRE*/ {
		usePot = true
	}

	if usePot {
		switch c.CurrentHealth {
		case 90:
			c.CurrentHealth += 10
		case 80:
			c.CurrentHealth += 20
		case 70:
			c.CurrentHealth += 30
		case 60:
			c.CurrentHealth += 40
		default:
			c.CurrentHealth += 50
		}

		fmt.Println("Tu as utilisé ", c.Inventory[item-1])
		c.Inventory = append(c.Inventory[:item-1], c.Inventory[item:]...)

	}
	usePot = false

}
