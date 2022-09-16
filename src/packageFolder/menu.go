// CHUPIN Tao -- CORBEL Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
	"os"
)

// Variable Global
var FirstMeet bool = true // Pour la fonction marchand

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
		os.Exit(1)
	default:
		DisplayErrorSwitch()
		c.PrincipalMenu()
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
		c.UseItemPerso()

	case 3:
		// c.RemoveItem([]string{})
	default:
		DisplayErrorSwitch()
		c.AccessInventory()
	}
}

func (c *Character) MenuMerchent() {
	var choice int

	if FirstMeet {
		DisplayMerchentFirstMeet()
		FirstMeet = false
	}

	DisplayMerchentMenu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.MenuAchatUtil()
	case 2:

	case 3:
		c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		c.MenuMerchent()
	}
}

func (c Character) MenuAchatUtil() {
	var choice int
	DisplayListItemUtil()
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.AddItem("Potion")
	case 2:
		// Ajout d'une potion advanced
	case 3:
		// Potion de poison
	case 4:
		// Totem réanimation*
	default:
		DisplayErrorSwitch()
	}
}

func (c Character) MenuAchatWeapon() {
}

func (c *Character) MenuSpellBook() {
	var Spell int

	DisplayMenuSpell()

	fmt.Scanln(&Spell)

	switch Spell {
	case 1:
		// c.FireBall()
	case 2:
		// c.Lightning()
	case 3:
		// c.Shield()
	case 4:
		// c.Action()
	}
}
