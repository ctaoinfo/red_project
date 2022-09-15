package packageFolder

import "fmt"

func (c Character) DisplayInfo() {
	fmt.Println(c.Name)
	fmt.Println(c.CurrentHealth, "HP")
}

func Menu() {
	var choice int

	switch choice {
	case 1:
		c.AccessInventory()
	}
}

func (c *Character) AccessInventory() {
	var item int
	var choice int
	fmt.Println("Vous avez actuellement ", c.CurrentHealth, "HP")
	fmt.Println("Acces à l'inventaire")

	fmt.Println("1. Afficher inventaire")
	fmt.Println("2. Utiliser un item")
	fmt.Println("3. Retour au menu")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		for _, itemPresent := range c.Inventory {
			fmt.Println(itemPresent)
		}
	case 2:
		fmt.Println("Quel item utiliser ?")

		for i, itemPresent := range c.Inventory {
			fmt.Println(i + 1)
			fmt.Println(itemPresent)
		}

		fmt.Scanln(&item)

		switch item {
		case 1:
			c.TakePot()
		}

		fmt.Println("Inventaire:", c.Inventory)
		fmt.Println("Vous avez ", c.CurrentHealth, "HP")
	case 3:
		Menu()
	}

}
