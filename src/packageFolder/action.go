// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
	"strconv"
)

func (m *Mob) PoisonPot() {
	DisplayFrame("Dégat Poison", []string{
		"Vous allez infliger 30 de dégats sur 2 round !",
		"Il restera " + strconv.Itoa(m.CurrentHealth-15) + "HP au monstre au prochain round",
		"Et " + strconv.Itoa(m.CurrentHealth-30) + "HP au monstre dans 2 round"})
	TakeDamage := 10

	for i := 0; i < 3; i++ {
		m.CurrentHealth -= TakeDamage
	}
}

func (c *Character) TakePot() {

	if c.CurrentHealth == c.Health || c.Inventory["Potion"] == 0 {
		if c.CurrentHealth == c.Health {
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion vous avez les HP maximum"})
		} else if c.Inventory["Potion"] == 0 {
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion  car vous en avez plus"})
		}

	} else {
		if c.CurrentHealth >= c.Health-50 || c.Inventory["Potion"] > 0 {
			c.CurrentHealth = c.Health
		} else {
			c.CurrentHealth += 50
		}
		c.Inventory["Potion"] -= 1
		DisplayFrame(" Soin applique", []string{"Tu as utilise une potion soin basic", "Vous avez a present " + strconv.Itoa(c.CurrentHealth) + "HP"})
	}
}

func (c *Character) TakeAdvancedPot() {
	if c.CurrentHealth == c.Health || c.Inventory["AdvancedPotion"] == 0 {
		if c.CurrentHealth == c.Health {
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion vous avez les HP maximum"})
		} else {
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion car vous en avez plus"})
		}

	} else {
		if c.CurrentHealth >= c.Health-50 || c.Inventory["AdvancedPotion"] > 0 {
			c.CurrentHealth = c.Health
		} else {
			c.CurrentHealth += 100
		}
		c.Inventory["AdvancedPotion"] -= 1
		DisplayFrame(" Soin applique", []string{"Tu as utilise une potion soin basic", "Vous avez a present " + strconv.Itoa(c.CurrentHealth) + "HP"})
	}
}

func (c *Character) UseItemPerso() {
	var item int

	var index int
	var txt string
	for i, it := range c.Inventory {
		if i == "Potion" {
			txt = "Potion soin basic"
			index = 1
		} else if i == "AdvancedPotion" {
			txt = "Potion soin avancé"
			index = 2
		} else if i == "ForcePotion" {
			txt = "Potion de force"
			index = 3
		}

		DisplayFrame(" Quel item utiliser ?", []string{
			"- " + strconv.Itoa(index) + ". " + txt + " - " + strconv.Itoa(it) + " disponible",
			"- 6 . Retour Accés inventaire"})
	}

	fmt.Scanln(&item)

	switch item {
	case 1:
		c.TakePot()
	case 2:
		c.TakeAdvancedPot()
	case 3:
		// c.TakeForcePot()
	case 4:
	case 6:
		c.AccessInventory()
	default:
		DisplayErrorSwitch()
		c.AccessInventory()
	}
	c.AccessInventory()
}

func (c *Character) FireBall() {
	DisplayFireBall()
}

func (c *Character) Lightning() {
	DisplayLightning()
}

func (c *Character) Healing() {
	DisplayHealing()
}

func (c *Character) Arrow() {
	DisplayArrow()
}

func (c *Character) GustArrow() {
	DisplayGustArrow()
}

func (c *Character) Escape() {
	DisplayEscape()
}

func (c *Character) Punch() {
	DisplayPunch()
}

func (c *Character) HeadButt() {
	DisplayHeadButt()
}

func (c *Character) Shield() {
	DisplayShield()
}

func (c *Character) Destruction() {
	DisplayDestruction()
}

// func (m *Mob) UseItemMob() {
// 	var item int

// 	fmt.Scanln(&item)

// 	switch item {
// 	case 1:
// 		m.PoisonPot()
// 	case 2:
// 	case 3:
// 	case 4:
// 	default:
// 	}
// }

func (c *Character) AddItem(AddItem string) {
	for item := range c.Inventory {
		if AddItem == item {
			c.Inventory[AddItem]++
		} else {
			c.Inventory[AddItem] = 1
		}
	}
}

func (c *Character) RemoveItem(RmItem string) {
	for item := range c.Inventory {
		if RmItem == item {
			c.Inventory[RmItem]--
		} else {
			c.Inventory[RmItem] = 1
		}
	}
}

func (c *Character) AddSkill(AddSkill string) {
	for skill := range c.Skill {
		if AddSkill == skill {
			c.Skill[AddSkill]++
		} else {
			c.Skill[AddSkill] = 1
		}
	}
}

func (c *Character) RemoveSkill(RmSkill string) {
	for skill := range c.Inventory {
		if RmSkill == skill {
			c.Inventory[RmSkill]--
		} else {
			c.Inventory[RmSkill] = 1
		}
	}
}

// func (m *Mob) BouleDeFeu() {
// 	fmt.Println("Vous utilisez Boule de feu !")
// 	m.CurrentHealth -= 20
// 	fmt.Println("Vous avez infligé 20 points de dégats !")
// }
