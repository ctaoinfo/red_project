// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
)

func (c *Character) TakePot() {

	fmt.Println("Vous avez ", c.CurrentHealth, "HP")
	if c.CurrentHealth == c.Health {
		fmt.Println("Vous ne pouvez pas utiliser la potion vous avez les HP maximum")
	} else {
		if c.CurrentHealth >= c.Health-50 /* ET si il reste potion dans inventaire*/ {
			c.CurrentHealth = c.Health
		} else {
			c.CurrentHealth += 50
		}
		fmt.Println("Tu as utilisé ", c.Inventory[Item-1])
		c.Inventory = append(c.Inventory[:Item-1], c.Inventory[Item:]...)
		fmt.Println("Maintenant vous avez ", c.CurrentHealth, "HP")
	}
}

func (c *Character) UseItem() {

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

func (c *Character) AddItem(item string) {
	c.Inventory = append(c.Inventory, item)
}

func (c Character) RemoveItem(item []string) {
	c.Inventory = append(c.Inventory[:Item-1], c.Inventory[Item:]...)
}

func (c Character) Count(list []string) map[string]int {
	duplicate := make(map[string]int)

	for _, item := range list {
		_, exist := duplicate[item]
		if exist {
			duplicate[item] += 1
		} else {
			duplicate[item] = 1
		}
	}
	return duplicate
}
