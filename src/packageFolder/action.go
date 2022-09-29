// CHUPIN Tao -- TARDELLI Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder // packageFolder

import ( // import
	"fmt"
	"strconv"
)

func (m *Character) PoisonPot() { // PoisonPot
	DisplayFrame("Dégat Poison", []string{ // Affichage dégâts Potion poison
		"Vous allez infliger 30 de dégats sur 2 round !",
		"Il restera " + strconv.Itoa(m.CurrentHealth-15) + "HP au monstre au prochain round",
		"Et " + strconv.Itoa(m.CurrentHealth-30) + "HP au monstre dans 2 round"})
	TakeDamage := 10 // Dégâts Potion poison

	for i := 0; i < 3; i++ { // Boucle pour les 3 round
		m.CurrentHealth -= TakeDamage // Dégâts Potion poison
	}
}

func (c *Character) TakePot() { // fonction prendre une potion

	if c.CurrentHealth == c.Health || c.Inventory["Potion"] == 0 { // si les HP sont au max ou si on a plus de potion
		if c.CurrentHealth == c.Health { // si les HP sont au max
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion vous avez les HP maximum"}) // affichage
		} else if c.Inventory["Potion"] == 0 { // si on a plus de potion
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion  car vous en avez plus"}) // affichage
		}

	} else { // si on a des potions et que les HP ne sont pas au max
		if c.CurrentHealth >= c.Health-50 && c.Inventory["Potion"] > 0 { // si les HP sont supérieur ou égale à 50 et qu'on a des potions
			c.CurrentHealth = c.Health // les HP sont au max
		} else { // si les HP sont inférieur à 50
			c.CurrentHealth += 50 // on ajoute 50 HP
		}
		c.Inventory["Potion"] -= 1                                                                                                                      // on retire une potion
		DisplayFrame(" Soin applique", []string{"Tu as utilisé une potion soin basic", "Vous avez a présent " + strconv.Itoa(c.CurrentHealth) + " HP"}) // affichage
	}
}

func (c *Character) TakeAdvancedPot() { // fonction prendre une potion avancé
	if c.CurrentHealth == c.Health || c.Inventory["AdvancedPotion"] == 0 { // si les HP sont au max ou si on a plus de potion
		if c.CurrentHealth == c.Health { // si les HP sont au max
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion vous avez les HP maximum"}) // affichage
		} else { // si on a plus de potion
			DisplayFrame(" Soin impossible", []string{"Vous avez " + strconv.Itoa(c.CurrentHealth) + "HP", "Vous ne pouvez pas utiliser la potion car vous en avez plus"}) // affichage
		}

	} else { // si on a des potions et que les HP ne sont pas au max
		if c.CurrentHealth >= c.Health-50 || c.Inventory["AdvancedPotion"] > 0 { // si les HP sont supérieur ou égale à 50 et qu'on a des potions
			c.CurrentHealth = c.Health // les HP sont au max
		} else { // si les HP sont inférieur à 50
			c.CurrentHealth += 100 // on ajoute 100 HP
		}
		c.Inventory["AdvancedPotion"] -= 1                                                                                                             // on retire une potion
		DisplayFrame(" Soin applique", []string{"Tu as utilise une potion soin basic", "Vous avez a present " + strconv.Itoa(c.CurrentHealth) + "HP"}) // affichage
	}
}

func (c *Character) UseItemPerso(m *Mob) { // fonction utiliser un item
	var item int // variable item

	var index int                    // variable index
	var txt string                   // variable txt
	for i, it := range c.Inventory { // boucle pour afficher les items
		if i == "Potion" { // si l'item est une potion
			txt = "Potion soin basic" // txt = potion soin basic
			index = 1                 // index = 1
		} else if i == "AdvancedPotion" { // si l'item est une potion avancé
			txt = "Potion soin avancé" // txt = potion soin avancé
			index = 2                  // index = 2
		} else if i == "ForcePotion" { // si l'item est une potion de force
			txt = "Potion de force" // txt = potion de force
			index = 3               // index = 3
		}

		DisplayFrame(" Quel item utiliser ?", []string{ // affichage
			"- " + strconv.Itoa(index) + ". " + txt + " - " + strconv.Itoa(it) + " disponible",
			"- 6 . Retour Accés inventaire"})
	}

	fmt.Scanln(&item) // scan de l'item

	switch item { // switch de l'item
	case 1: // si l'item est 1
		c.TakePot() // on prend la potion
	case 2: // si l'item est 2
		c.TakeAdvancedPot() // on prend la potion avancé
	case 3: // si l'item est 3
		// c.TakeForcePot()
	case 4:
	case 6: // si l'item est 6
		c.AccessInventory(m) // on retourne à l'accés inventaire
	default: // si l'item n'est pas 1, 2, 3, 4 ou 6
		DisplayErrorSwitch() // affichage erreur
		c.AccessInventory(m) // on retourne à l'accés inventaire
	}
	c.AccessInventory(m) // on retourne à l'accés inventaire
}

func (c *Character) FireBall() { // fonction boule de feu
	DisplayFireBall() // affichage boule de feu
}

func (c *Character) Lightning() { // fonction éclair
	DisplayLightning() // affichage éclair
}

func (c *Character) Healing() { // fonction soin
	DisplayHealing() // affichage soin
}

func (c *Character) Arrow() { // fonction flèche
	DisplayArrow() // affichage flèche
}

func (c *Character) GustArrow() { // fonction rafale de flèches
	DisplayGustArrow() // affichage rafale de flèches
}

func (c *Character) Escape() { // fonction fuir
	DisplayEscape() // affichage fuir
}

func (c *Character) Punch() { // fonction coup de poing
	DisplayPunch() // affichage coup de poing
}

func (c *Character) HeadButt() { // fonction coup de tête
	DisplayHeadButt() // affichage coup de tête
}

func (c *Character) Shield() { // fonction bouclier
	DisplayShield() // affichage bouclier
}

func (c *Character) Destruction() { // fonction destruction
	DisplayDestruction() // affichage destruction
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

func (c *Character) AddItem(AddItem string) { // fonction ajouter un item
	verif := false                  // variable verif
	for item := range c.Inventory { // boucle pour vérifier si l'item est dans l'inventaire
		if AddItem == item { // si l'item est dans l'inventaire
			c.Inventory[AddItem]++ // on ajoute 1 à l'item
			verif = true           // verif = true
			break                  // on sort de la boucle
		} else { // si l'item n'est pas dans l'inventaire
			verif = false // verif = false
		}
	}
	if !verif { // si verif = false
		c.Inventory[AddItem] = 1 // on ajoute l'item
	}

}

func (c *Character) RemoveItem() { // fonction retirer un item
	var choice int    // variable choix
	var RmItem string // variable retirer un item

	DisplaySellItem()   // affichage vendre un item
	fmt.Scanln(&choice) // scan du choix
	switch choice {     // switch du choix
	case 1: // si le choix est 1
		RmItem = "Potion" // retirer un item = potion
	case 2: // si le choix est 2
		RmItem = "AdvancedPotion" // retirer un item = potion avancé
	case 3: // si le choix est 3
		RmItem = "ForcePotion" // retirer un item = potion de force
	case 4: // si le choix est 4
		RmItem = "PoisonPotion" // retirer un item = potion de poison
	}
	for item := range c.Inventory { // boucle pour vérifier si l'item est dans l'inventaire
		if RmItem == item { // si l'item est dans l'inventaire
			c.Inventory[RmItem]-- // on retire 1 à l'item
		} else { // si l'item n'est pas dans l'inventaire
			c.Inventory[RmItem] = 0 // on retire l'item
		}
	}
	if RmItem == "Potion" { // si retirer un item = potion
		c.Money += 3 // on ajoute 3 au money
	} else if RmItem == "AdvancedPotion" { // si retirer un item = potion avancé
		c.Money += 8 // on ajoute 8 au money
	} else if RmItem == "ForcePotion" { // si retirer un item = potion de force
		c.Money += 6 // on ajoute 6 au money
	} else if RmItem == "PoisonPotion" { // si retirer un item = potion de poison
		c.Money += 6 // on ajoute 6 au money
	}
}

func (c *Character) spellBook(nom string) { // fonction livre de sort
	if c.Skill[nom] == 0 { // si le sort n'est pas appris
		c.Skill[nom] += 1     // on apprend le sort
		c.Inventory[nom] -= 1 // on retire 1 au sort
	} else { // si le sort est appris
		DisplayFrame("Sort", []string{"Vous avez déjà appris ce sort !"}) // affichage vous avez déjà appris ce sort
	}
}

// func (m *Mob) BouleDeFeu() {
// 	fmt.Println("Vous utilisez Boule de feu !")
// 	m.CurrentHealth -= 20
// 	fmt.Println("Vous avez infligé 20 points de dégats !")
// }

func (c *Character) Fuir() { // fonction fuir
	c.CurrentHealth = 0 // on met la vie du personnage à 0
}
