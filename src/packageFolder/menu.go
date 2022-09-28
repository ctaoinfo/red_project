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
		c.AccessInventory()
	case 3:
		c.DisplayMoney()
		c.AccessInventory()
	case 4:
		c.UseItemPerso()
		c.AccessInventory()
	case 5:
		//c.L
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

	if !c.VerifFullInventoryItems() {
		c.DisplayFullInventory()
		c.MenuMerchent()
	}

	fmt.Scanln(&choice)
	switch choice {
	case 1:
		if !c.VerifFullInventoryItemsAmount("PlumeCorbeau") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		}
		c.AddItem("PlumeCorbeau")
		c.Money -= 3
	case 2:
		if !c.VerifFullInventoryItemsAmount("FourrureLoup") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		}
		c.AddItem("FourrureLoup")
		c.Money -= 5
	case 3:
		if !c.VerifFullInventoryItemsAmount("PeauTroll") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		}
		c.AddItem("PeauTroll")
		c.Money -= 8
	case 4:
		if !c.VerifFullInventoryItemsAmount("CuirSanglier") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		}
		c.AddItem("CuirSanglier")
		c.Money -= 11
	case 5:
		if !c.VerifFullInventoryItemsAmount("FilAraigne") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		}
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
	if !c.VerifFullInventoryItems() {
		c.DisplayFullInventory()
		c.MenuMerchent()
	}
	switch choice {
	case 1:
		if !c.VerifFullInventoryItemsAmount("Potion") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		} else {
			c.AddItem("Potion")
			c.Money -= 6
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Soin Basic"})
		}
	case 2:
		if !c.VerifFullInventoryItemsAmount("AdvancedPotion") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		} else {
			c.AddItem("AdvancedPotion")
			c.Money -= 16
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Soin Avancée"})
		}
	case 3:
		if !c.VerifFullInventoryItemsAmount("ForcePotion") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		} else {
			c.AddItem("ForcePotion")
			c.Money -= 12
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Force"})
		}
	case 4:
		if !c.VerifFullInventoryItemsAmount("PosionPotion") {
			c.DisplayFullInventory()
			c.MenuMerchent()
		} else {
			c.AddItem("PoisonPotion")
			c.Money -= 9
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Poison"})
		}
	case 5:
		c.PrincipalMenu()
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
	c.DisplayListSpell()
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.LivreDeSort("Boule De Feu")
	case 2:
		c.LivreDeSort("")
	case 3:
		c.LivreDeSort("")
	case 4:
		c.LivreDeSort("")
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
	return len(c.Inventory) < 10
}
func (c *Character) VerifFullInventoryItemsAmount(item string) bool {

	return c.Inventory[item] < 10

}
