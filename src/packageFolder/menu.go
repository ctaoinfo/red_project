// CHUPIN Tao -- CORBEL Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder

import "fmt"

// Variable Global
var FirstMeet bool = true

func (c *Character) PrincipalMenu() {
	var choice int

	DisplayPrincipalMenu()

	fmt.Scanln(&choice)
	fmt.Print("\n")

	switch choice {
	case 1:
		c.DisplayInfoPlayer()
	case 2:
		c.AccessInventory()
	case 3:
		c.MenuMerchent()
	case 4:
		DisplayQuiSontIls()
	case 5:
		return
	}

}

func (c *Character) AccessInventory() {

	var choice int
	DisplayInventoryMenu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		firstPass := true
		result := c.Count(c.Inventory)
		for index, itemPresent := range result {
			if firstPass {
				fmt.Println("+----------+")
				firstPass = false
			}
			fmt.Printf("| ")
			fmt.Printf("%d %s", itemPresent, index)
			fmt.Println(" |")
			fmt.Println("+----------+")

		}
	case 2:
		c.UseItem()

	case 3:
		c.ThrowItem()

	default:
		return
	}
}

func (c Character) MenuMerchent() {
	var choice int

	if FirstMeet {
		DisplayMerchentFirstMeet()
		FirstMeet = false
	}

	DisplayMerchentMenu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:

	case 2:

	case 3:

	}

}
