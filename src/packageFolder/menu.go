// CHUPIN Tao -- TARDELLI Andrea
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
	c.PlayerDead()
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

	c.PlayerDead()
	DisplayInventoryMenu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.DisplayInventory()
		c.AccessInventory()
	case 2:
		c.AccessStuffInventory()
	case 3:
		c.DisplayMoney()
	case 4:
		c.UseItemPerso()

	case 5:
		// c.RemoveItem()
	case 6:
		c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		c.AccessInventory()
	}
}

func (c *Character) AccessStuffInventory() {
	var choice int

	c.PlayerDead()
	DisplayStuffInventory()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// c.DisplayStuff()
		c.AccessStuffInventory()
	case 2:
		// c.AddStuff

	case 3:
		// c.RemoveStuff()
	case 4:
		c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		c.AccessInventory()
	}
}

func (c *Character) MenuMerchent() {
	var choice int

	c.PlayerDead()

	if FirstMeet {
		DisplayMerchentFirstMeet()
		c.AddItem("Potion")
		FirstMeet = false
	}

	DisplayMerchentMenu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.MenuBuyUtil()
	case 2:
		c.MenuBuySkill()
	case 3:
		c.MenuBuyStuff()
	case 4:
		c.MenuSellItem()
	case 5:
		c.DisplayInventory()
	case 6:
		c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		c.MenuMerchent()
	}
}

func (c *Character) MenuBuyStuff() {
	var choice int

	c.PlayerDead()
	DisplayListStuff()

	if c.VerifFullInventoryItems() {
		c.MenuBuyStuff()
	}
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		c.AddItem("PlumeCorbeau")
		c.Money -= 3
	case 2:
		c.AddItem("FourrureLoup")
		c.Money -= 5
	case 3:
		c.AddItem("PeauTroll")
		c.Money -= 8
	case 4:
		c.AddItem("CuirSanglier")
		c.Money -= 11
	case 5:
		c.AddItem("FilAraigne")
		c.Money -= 15
	case 6:
		c.MenuMerchent()
	default:
		DisplayErrorSwitch()
		c.MenuBuyStuff()
	}
}

func (c *Character) MenuSellItem() {
	var choice int

	c.PlayerDead()
	DisplaySellMenu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:

		c.RemoveItem()
	case 5:
		c.DisplayInventory()
	case 6:
		c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		c.MenuSellItem()
	}
}

func (c *Character) MenuBuyUtil() {
	var choice int

	c.PlayerDead()
	DisplayListItemUtil()
	fmt.Scanln(&choice)
	if c.VerifFullInventoryItems() {
		c.MenuMerchent()
	}
	switch choice {
	case 1:
		c.AddItem("Potion")
		c.Money -= 6
	case 2:
		c.AddItem("AdvancedPotion")
		c.Money -= 16
	case 3:
		c.AddItem("ForcePotion")
		c.Money -= 12
	case 4:
		c.AddItem("PoisonPotion")
		c.Money -= 9
	default:
		DisplayErrorSwitch()
		c.MenuBuyUtil()
	}
}

func (c *Character) MenuBuyWeapon() {
}

func (c *Character) MenuBuySkill() {
	var choice int

	c.PlayerDead()
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
		c.MenuBuyUtil()
	}
}

func (c *Character) MenuSkillBook() {
	var skill int

	c.PlayerDead()
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

func (c *Character) MenuFight() {
	var choice int

	c.PlayerDead()
	DisplayMenuFight()

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// mob.DisplayInfoEnemy()
	case 2:

	case 3:

	case 4:
		// c.PrincipalMenu()
	default:
		DisplayErrorSwitch()
		// mob.MenuFight()

	}
}

func (c *Character) VerifFullInventoryItems() bool {
	if len(c.Inventory.Items) <= 10 {
		return false
	} else {
		c.DisplayFullInventory()
		return true
	}
}

func (c *Character) VerifFullInventoryEquipment() bool {
	if len(c.Inventory.Equipment.Chestplate) <= 10 {
		return false
	} else {
		c.DisplayFullInventory()
		return true
	}
}
