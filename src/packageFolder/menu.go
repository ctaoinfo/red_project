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

	switch choice {
	case 1:
		c.DisplayInfoPlayer()
	case 2:
		c.AccessInventory()
	case 3:
		c.MenuMerchent()
	case 4:
		// c.MenuFight()
	case 5:
		DisplayQuiSontIls()
	case 6:
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
		c.DisplayInventory()
		c.AccessInventory()
	case 2:
		c.UseItemPerso()

	case 3:
		// c.RemoveItem()
	case 4:
		c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		c.AccessInventory()
	}
}

func (c *Character) MenuMerchent() {
	var choice int

	if FirstMeet {
		DisplayMerchentFirstMeet()
		c.AddItem("Potion")
		FirstMeet = false
	}

	DisplayMerchentMenu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.MenuAchatUtil()
	case 2:
		c.MenuAchatSkill()
	case 3:
		c.DisplayInventory()
	case 4:
		c.DisplayInventory()
	default:
		DisplayErrorSwitch()
		c.MenuMerchent()
	}
}

func (c *Character) MenuAchatUtil() {
	var choice int
	DisplayListItemUtil()
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.AddItem("Potion")
	case 2:
		c.AddItem("AdvancedPotion")
	case 3:
		c.AddItem("ForcePotion")
	case 4:
		c.AddItem("PoisonPotion")
	default:
		DisplayErrorSwitch()
		c.MenuAchatUtil()
	}
}

func (c *Character) MenuAchatWeapon() {
}

func (c *Character) MenuAchatSkill() {
	var choice int
	DisplayListItemUtil()
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.AddSkill("")
	case 2:
		c.AddSkill("")
	case 3:
		c.AddSkill("")
	case 4:
		c.AddSkill("")
	default:
		DisplayErrorSwitch()
		c.MenuAchatUtil()
	}
}

func (c *Character) MenuSkillBook() {
	var skill int

	c.DisplayMenuSkillBook()

	fmt.Scanln(&skill)
	if c.Class == "Sorcier" {
		switch skill {
		case 1:
			c.FireBall()
		case 2:
			c.Lightning()
		case 3:
			c.Healing()
		case 4:
			c.AccessInventory()
		}
	} else if c.Class == "Archer" {
		switch skill {
		case 1:
			c.Arrow()
		case 2:
			c.GustArrow()
		case 3:
			c.Escape()
		case 4:
			c.AccessInventory()
		}
	} else if c.Class == "Tank" {
		switch skill {
		case 1:
			c.Punch()
		case 2:
			c.HeadButt()
		case 3:
			c.Shield()
		case 4:
			c.AccessInventory()
		}
	} else if c.Class == "ADMIN" {
		switch skill {
		case 1:
			c.Destruction()
		case 4:
			c.AccessInventory()
		}
	}
}

func (m *Mob) MenuFight() {
	var choice int

	DisplayMenuFight()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		m.DisplayInfoEnemy()
	case 2:

	case 3:

	case 4:
		// c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		m.MenuFight()

	}
	// if c.Inventory["PoisonPot"] == 0 {
	// 	DisplayFrame("Potion posion inutilisable", []string{
	// 		"Vous ne pouvez pas utilisez une potion poison car vous n'en avait pas"})
	// }
}
