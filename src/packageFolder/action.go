// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
	"time"
)

func (m *Mob) PoisonPot() {
	fmt.Println("vous avez bu une potion de poison")
	TakeDamage := 10

	for i := 0; i < 3; i++ {
		m.CurrentHealth -= TakeDamage
		fmt.Println("Vous avez perdu 10 points de vie !")
		time.Sleep(1 * time.Second)
	}
}

func (c *Character) TakePot() {
	var item int
	fmt.Println("Vous avez ", c.CurrentHealth, "HP")

	if c.CurrentHealth == c.Health {

		fmt.Println("Vous ne pouvez pas utiliser la potion vous avez les HP maximum")

	} else {

		if c.CurrentHealth >= c.Health-50 /* ET si il reste potion dans inventaire*/ {
			c.CurrentHealth = c.Health
		} else {
			c.CurrentHealth += 50
		}
		fmt.Println("Tu as utilisé ", c.Inventory[item-1])
		c.Inventory = append(c.Inventory[:item-1], c.Inventory[item:]...)
		fmt.Println("Maintenant vous avez ", c.CurrentHealth, "HP")
	}
}

func (c *Character) UseItemPerso() {
	var item int
	fmt.Println("Quel item utiliser ?")

	// for i, itemPresent := range c.Inventory {
	// 	fmt.Println(i + 1)
	// 	fmt.Println(c.Inventory[1], "x")
	// }

	fmt.Scanln(&item)

	switch item {
	case 1:
		c.TakePot()
	case 2:
	case 3:
	case 4:
	default:
	}
	fmt.Println("Inventaire:", c.Inventory)
}

func (m *Mob) UseItemMob() {
	var item int
	fmt.Println("Quel item utiliser ?")

	// for i, itemPresent := range c.Inventory {
	// 	fmt.Println(i + 1)
	// 	fmt.Println(c.Inventory[1], "x")
	// }

	fmt.Scanln(&item)

	switch item {
	case 1:
		m.PoisonPot()
	case 2:
	case 3:
	case 4:
	default:
	}
}

func (c *Character) AddItem(AddItem string) {
	c.Inventory = append(c.Inventory, AddItem)
}

func (c *Character) RemoveItem(RemoveItem string) {
	c.Inventory = append(c.Inventory, RemoveItem)
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
func (m *Mob) BouleDeFeu() {
	fmt.Println("Vous utilisez Boule de feu !")
	m.CurrentHealth -= 20
	fmt.Println("Vous avez infligé 20 points de dégats !")
}
