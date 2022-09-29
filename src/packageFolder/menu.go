// CHUPIN Tao -- TARDELLI Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
	"os"
)

// Variable Globale

var FirstMeet bool = true // Pour la fonction marchand

func (c *Character) PrincipalMenu(m *Mob) { // Menu principal
	var choice int         // Variable pour le choix
	c.PlayerDead()         // Vérifie si le joueur est mort
	DisplayPrincipalMenu() // Affiche le menu principal

	fmt.Scanln(&choice) // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		c.DisplayInfoPlayer() // Affiche les infos du joueur
	case 2: // Si le choix est 2
		c.AccessInventory(m) // Affiche l'inventaire du joueur
	case 3: // Si le choix est 3
		c.MenuMerchent(m) // Affiche le menu du marchand
	case 4: // Si le choix est 4
		c.MenuForgeron(m) // Affiche le menu du forgeron
	case 5: // Si le choix est 5
		m.FightMob(1, c) // Affiche le menu du combat
	case 6: // Si le choix est 6
		DisplayQuiSontIls() // Affiche les infos des créateurs
	case 7: // Si le choix est 7
		DisplayImageAsciiFin() // Affiche l'image de fin
		os.Exit(1)             // Quitte le programme
	default:
		DisplayErrorSwitch() // Affiche un message d'erreur
		c.PrincipalMenu(m)   // Affiche le menu principal
	}
}

func (c *Character) AccessInventory(m *Mob) { // Menu de l'inventaire
	var choice int // Variable pour le choix

	c.PlayerDead()         // Vérifie si le joueur est mort
	DisplayInventoryMenu() // Affiche le menu de l'inventaire

	fmt.Scanln(&choice) // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		c.DisplayInventory() // Affiche l'inventaire du joueur
		c.AccessInventory(m) // Affiche le menu de l'inventaire
	case 2:
		c.AccessStuffInventory(m) // Affiche le menu de l'inventaire
		c.AccessInventory(m)      // Affiche le menu de l'inventaire
	case 3:
		c.DisplayMoney()
		c.AccessInventory(m)
	case 4:
		c.UseItemPerso(m)
		c.AccessInventory(m)
	case 5:
		c.MenuSkillBook(m)
		c.AccessInventory(m)
	case 6:
		c.PrincipalMenu(m)
	default:
		DisplayErrorSwitch()
		c.AccessInventory(m)
	}
}

func (c *Character) AccessStuffInventory(m *Mob) { // Menu de l'inventaire stuff
	var choice int // Variable pour le choix

	c.PlayerDead()              // Vérifie si le joueur est mort
	DisplayMenuStuffInventory() // Affiche le menu de l'inventaire stuff

	fmt.Scanln(&choice) // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		// c.DisplayStuff()
		c.AccessStuffInventory(m) // Affiche le menu de l'inventaire stuff
	case 2:
		// c.AddStuff

	case 3:
		// c.RemoveStuff()
	case 4:
		c.PrincipalMenu(m) // Affiche le menu principal
	default:
		DisplayErrorSwitch() // Affiche un message d'erreur
		c.AccessInventory(m) // Affiche le menu de l'inventaire
	}
}

func (c *Character) MenuMerchent(m *Mob) { // Menu du marchand
	var choice int // Variable pour le choix

	c.PlayerDead() // Vérifie si le joueur est mort

	if FirstMeet { // Si c'est la première fois que le joueur rencontre le marchand
		DisplayMerchentFirstMeet() // Affiche le message de la première rencontre
		c.AddItem("Potion")        // Ajoute une potion dans l'inventaire du joueur gratuitement
		FirstMeet = false          // Change la valeur de FirstMeet
	}

	DisplayMerchentMenu() // Affiche le menu du marchand

	fmt.Scanln(&choice) // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		c.MenuBuyUtil(m) // Affiche le menu d'achat d'utilitaire
	case 2: // Si le choix est 2
		c.MenuBuySkill(m) // Affiche le menu d'achat de compétence
	case 3: // Si le choix est 3
		c.MenuBuyStuff(m) // Affiche le menu d'achat d'équipement
	case 4: // Si le choix est 4
		c.MenuSellItem(m) // Affiche le menu de vente d'item
	case 5: // Si le choix est 5
		c.DisplayInventory() // Affiche l'inventaire du joueur
	case 6: // Si le choix est 6
		c.PrincipalMenu(m) // Affiche le menu principal
	default: // Si le choix est autre
		DisplayErrorSwitch() // Affiche un message d'erreur
		c.MenuMerchent(m)    // Affiche le menu du marchand
	}
}

func (c *Character) MenuForgeron(m *Mob) { // Menu du forgeron
	var choice int // Variable pour le choix

	c.PlayerDead()          // Vérifie si le joueur est mort
	c.DisplayMenuForgeron() // Affiche le menu du forgeron

	fmt.Scanln(&choice) // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		c.MenuBuyEquipment(m) // Affiche le menu d'achat d'équipement
	case 2: // Si le choix est 2
		c.AccessInventory(m) // Affiche le menu de l'inventaire
	case 3: // Si le choix est 3
		c.PrincipalMenu(m) // Affiche le menu principal
	default: // Si le choix est autre
		DisplayErrorSwitch() // Affiche un message d'erreur
		c.MenuForgeron(m)    // Affiche le menu du forgeron
	}
}

func (c *Character) MenuBuyStuff(m *Mob) { // Menu d'achat d'équipement
	var choice int // Variable pour le choix

	c.PlayerDead()     // Vérifie si le joueur est mort
	DisplayListStuff() // Affiche la liste des équipements

	if !c.VerifFullInventoryItems() { // Si l'inventaire du joueur n'est pas plein
		c.DisplayFullInventory() // Affiche un message d'inventaire plein
		c.MenuMerchent(m)        // Affiche le menu du marchand
	}

	fmt.Scanln(&choice) // Récupère le choix
	switch choice {     // Switch pour le choix
	case 1: // Si le choix est 1
		if !c.VerifFullInventoryItemsAmount("PlumeCorbeau") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		}
		c.AddItem("PlumeCorbeau") // Ajoute une plume de corbeau dans l'inventaire du joueur
		c.Money -= 36             // Retire 36 pièces d'or au joueur
	case 2: // Si le choix est 2
		if !c.VerifFullInventoryItemsAmount("FourrureLoup") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		}
		c.AddItem("FourrureLoup") // Ajoute une fourrure de loup dans l'inventaire du joueur
		c.Money -= 45             // Retire 45 pièces d'or au joueur
	case 3: // Si le choix est 3
		if !c.VerifFullInventoryItemsAmount("PeauTroll") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		}
		c.AddItem("PeauTroll") // Ajoute une peau de troll dans l'inventaire du joueur
		c.Money -= 32          // Retire 32 pièces d'or au joueur
	case 4: // Si le choix est 4
		if !c.VerifFullInventoryItemsAmount("CuirSanglier") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		}
		c.AddItem("CuirSanglier") // Ajoute un cuir de sanglier dans l'inventaire du joueur
		c.Money -= 28             // Retire 11 pièces d'or au joueur
	case 5: // Si le choix est 5
		if !c.VerifFullInventoryItemsAmount("FilAraigne") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		}
		c.AddItem("FilAraigne") // Ajoute un fil d'araignée dans l'inventaire du joueur
		c.Money -= 34           // Retire 15 pièces d'or au joueur
	case 6: // Si le choix est 6
		c.MenuMerchent(m) // Affiche le menu du marchand
	default: // Si le choix est autre
		DisplayErrorSwitch() // Affiche un message d'erreur
		c.MenuBuyStuff(m)    // Affiche le menu d'achat d'équipement
	}
}

func (c *Character) MenuSellItem(m *Mob) { // Menu de vente d'item
	var choice int // Variable pour le choix

	c.PlayerDead()    // Vérifie si le joueur est mort
	DisplaySellMenu() // Affiche le menu de vente d'item

	fmt.Scanln(&choice) // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		c.RemoveItem() // Retire un item de l'inventaire du joueur
	case 5: // Si le choix est 5
		c.DisplayInventory() // Affiche l'inventaire du joueur
	case 6: // Si le choix est 6
		c.PrincipalMenu(m) // Affiche le menu principal
	default: // Si le choix est autre
		DisplayErrorSwitch() // Affiche un message d'erreur
		c.MenuSellItem(m)    // Affiche le menu de vente d'item
	}
}

func (c *Character) MenuBuyUtil(m *Mob) { // Menu d'achat d'utilitaire
	var choice int // Variable pour le choix

	c.PlayerDead()                    // Vérifie si le joueur est mort
	DisplayListItemUtil()             // Affiche la liste des utilitaires
	fmt.Scanln(&choice)               // Récupère le choix
	if !c.VerifFullInventoryItems() { // Si l'inventaire du joueur n'est pas plein
		c.DisplayFullInventory() // Affiche un message d'inventaire plein
		c.MenuMerchent(m)        // Affiche le menu du marchand
	}
	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		if !c.VerifFullInventoryItemsAmount("Potion") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		} else { // Si l'inventaire du joueur est plein
			c.AddItem("Potion")                                                           // Ajoute une potion dans l'inventaire du joueur
			c.Money -= 6                                                                  // Retire 6 pièces d'or au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Soin Basic"}) // Affiche un message de confirmation
		}
	case 2: // Si le choix est 2
		if !c.VerifFullInventoryItemsAmount("AdvancedPotion") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		} else { // Si l'inventaire du joueur est plein
			c.AddItem("AdvancedPotion")                                                     // Ajoute une potion dans l'inventaire du joueur
			c.Money -= 16                                                                   // Retire 16 pièces d'or au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Soin Avancée"}) // Affiche un message de confirmation
		}
	case 3: // Si le choix est 3
		if !c.VerifFullInventoryItemsAmount("ForcePotion") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		} else { // Si l'inventaire du joueur est plein
			c.AddItem("ForcePotion")                                                 // Ajoute une potion dans l'inventaire du joueur
			c.Money -= 12                                                            // Retire 12 pièces d'or au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Force"}) // Affiche un message de confirmation
		}
	case 4: // Si le choix est 4
		if !c.VerifFullInventoryItemsAmount("PosionPotion") { // Si l'inventaire du joueur n'est pas plein
			c.DisplayFullInventory() // Affiche un message d'inventaire plein
			c.MenuMerchent(m)        // Affiche le menu du marchand
		} else { // Si l'inventaire du joueur est plein
			c.AddItem("PoisonPotion")                                                 // Ajoute une potion dans l'inventaire du joueur
			c.Money -= 9                                                              // Retire 9 pièces d'or au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter une Potion de Poison"}) // Affiche un message de confirmation
		}
	case 5: // Si le choix est 5
		c.PrincipalMenu(m) // Affiche le menu principal
	default: // Si le choix est autre
		DisplayErrorSwitch() // Affiche un message d'erreur
		c.MenuBuyUtil(m)     // Affiche le menu d'achat d'utilitaire
	}
}

func (c *Character) MenuBuyEquipment(m *Mob) { // Menu d'achat d'équipement
	var choice int // Variable pour le choix

	c.PlayerDead()         // Vérifie si le joueur est mort
	DisplayListEquipment() // Affiche la liste des équipements
	fmt.Scanln(&choice)    // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		if c.Money >= 10 { // Si le joueur a assez d'argent
			c.Money -= 10         // Retire 10 pièces d'or au joueur
			c.Forgeron(choice, m) // Affiche le menu du forgeron
		} else { // Si le joueur n'a pas assez d'argent
			DisplayFrame("Achat", []string{"Vous n'avez pas assez d'argent"}) // Affiche un message d'erreur
		}
	case 2: // Si le choix est 2
		if c.Money >= 15 { // Si le joueur a assez d'argent
			c.Money -= 15         // Retire 15 pièces d'or au joueur
			c.Forgeron(choice, m) // Affiche le menu du forgeron

		} else { // Si le joueur n'a pas assez d'argent
			DisplayFrame("Achat", []string{"Vous n'avez pas assez d'argent"}) // Affiche un message d'erreur
		}
	case 3: // Si le choix est 3
		if c.Money >= 20 { // Si le joueur a assez d'argent
			c.Money -= 20         // Retire 20 pièces d'or au joueur
			c.Forgeron(choice, m) // Affiche le menu du forgeron

		} else { // Si le joueur n'a pas assez d'argent
			DisplayFrame("Achat", []string{"Vous n'avez pas assez d'argent"}) // Affiche un message d'erreur
		}
	case 4: // Si le choix est 4
		c.MenuForgeron(m) // Affiche le menu du forgeron
	default: // Si le choix est autre
		DisplayErrorSwitch()  // Affiche un message d'erreur
		c.MenuBuyEquipment(m) // Affiche le menu d'achat d'équipement
	}
}

func (c *Character) MenuBuySkill(m *Mob) { // Menu d'achat de compétence
	var choice int // Variable pour le choix

	c.PlayerDead()               // Vérifie si le joueur est mort
	c.DisplayListBoutiqueSpell() // Affiche la liste des compétences
	fmt.Scanln(&choice)          // Récupère le choix

	if c.Class == "Sorcier" { // Si le joueur est un sorcier
		switch choice { // Switch pour le choix
		case 1: // Si le choix est 1
			c.spellBook("Boule de feu")     // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuBouleDeFeu() // Affiche un message de confirmation
		case 2: // Si le choix est 2
			c.spellBook("Eclair")       // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuEclair() // Affiche un message de confirmation
		case 3: // Si le choix est 3
			c.spellBook("Anneau soin")      // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuAnneauSoin() // Affiche un message de confirmation
		case 4: // Si le choix est 4
			c.spellBook("") // Ajoute la compétence dans le grimoire du joueur
		default: // Si le choix est autre
			DisplayErrorSwitch() // Affiche un message d'erreur
			c.MenuBuyUtil(m)     // Affiche le menu d'achat de compétence
		}
	} else if c.Class == "Archer" { // Si le joueur est un archer
		switch choice { // Switch pour le choix
		case 1: // Si le choix est 1
			c.spellBook("Flèche")       // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuFleche() // Affiche un message de confirmation
		case 2: // Si le choix est 2
			c.spellBook("Rafale flèche") // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuRafale()  // Affiche un message de confirmation
		case 3: // Si le choix est 3
			c.spellBook("Escive")        // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuEsquive() // Affiche un message de confirmation
		case 4: // Si le choix est 4
			c.spellBook("") // Ajoute la compétence dans le grimoire du joueur
		default: // Si le choix est autre
			DisplayErrorSwitch() // Affiche un message d'erreur
			c.MenuBuyUtil(m)     // Affiche le menu d'achat de compétence
		}
	} else if c.Class == "Tank" { // Si le joueur est un tank
		switch choice { // Switch pour le choix
		case 1: // Si le choix est 1
			c.spellBook("Coup de poing")     // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuCoupDePoing() // Affiche un message de confirmation
		case 2: // Si le choix est 2
			c.spellBook("Coup de tête")     // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuCoupDeTete() // Affiche un message de confirmation
		case 3: // Si le choix est 3
			c.spellBook("Bouclier")       // Ajoute la compétence dans le grimoire du joueur
			c.DisplaySortObtenuBouclier() // Affiche un message de confirmation
		case 4: // Si le choix est 4
			c.spellBook("") // Ajoute la compétence dans le grimoire du joueur
		default: // Si le choix est autre
			DisplayErrorSwitch() // Affiche un message d'erreur
			c.MenuBuyUtil(m)     // Affiche le menu d'achat de compétence
		}
	} else if c.Class == "ADMIN" { // Si le joueur est un admin
		switch choice { // Switch pour le choix
		case 1: // Si le choix est 1
			c.Destruction() // Détruit le joueur
		case 4: // Si le choix est 4
			c.AccessInventory(m) // Affiche l'inventaire du joueur
		}
	}
}

func (c *Character) MenuSkillBook(m *Mob) { // Menu du grimoire
	var skill int // Variable pour le choix

	c.PlayerDead()           // Vérifie si le joueur est mort
	c.DisplayMenuSkillBook() // Affiche le menu du grimoire

	fmt.Scanln(&skill)        // Récupère le choix
	if c.Class == "Sorcier" { // Si le joueur est un sorcier
		switch skill { // Switch pour le choix
		case 1: // Si le choix est 1
			c.FireBall() // Lance la compétence
		case 2: // Si le choix est 2
			c.Lightning() // Lance la compétence
		case 3: // Si le choix est 3
			c.Healing() // Lance la compétence
		case 4: // Si le choix est 4
			c.AccessInventory(m) // Affiche l'inventaire du joueur
		}
	} else if c.Class == "Archer" { // Si le joueur est un archer
		switch skill { // Switch pour le choix
		case 1: // Si le choix est 1
			c.Arrow() // Lance la compétence
		case 2: // Si le choix est 2
			c.GustArrow() // Lance la compétence
		case 3: // Si le choix est 3
			c.Escape() // Lance la compétence
		case 4: // Si le choix est 4
			c.AccessInventory(m) // Affiche l'inventaire du joueur
		}
	} else if c.Class == "Tank" { // Si le joueur est un tank
		switch skill { // Switch pour le choix
		case 1: // Si le choix est 1
			c.Punch() // Lance la compétence
		case 2: // Si le choix est 2
			c.HeadButt() // Lance la compétence
		case 3: // Si le choix est 3
			c.Shield() // Lance la compétence
		case 4: // Si le choix est 4
			c.AccessInventory(m) // Affiche l'inventaire du joueur
		}
	} else if c.Class == "ADMIN" { // Si le joueur est un admin
		switch skill { // Switch pour le choix
		case 1: // Si le choix est 1
			c.Destruction() // Détruit le joueur
		case 4: // Si le choix est 4
			c.AccessInventory(m) // Affiche l'inventaire du joueur
		}
	}
}

func (m *Mob) MenuFight(c *Character) { // Menu de combat
	var choice int // Variable pour le choix

	c.PlayerDead()     // Vérifie si le joueur est mort
	DisplayMenuFight() // Affiche le menu de combat

	fmt.Scanln(&choice) // Récupère le choix

	switch choice { // Switch pour le choix
	case 1: // Si le choix est 1
		m.DisplayGoblinInfo() // Affiche les informations du gobelin
	case 2: // Si le choix est 2

	case 3: // Si le choix est 3
		DisplayFuir()                  // Affiche un message de fuite
		c.CurrentHealth = c.Health / 2 // Met la vie du joueur à la moitié
		c.PrincipalMenu(m)             // Affiche le menu principal
		c.Money -= 20                  // Retire 20 pièces d'or au joueur
	case 4: // Si le choix est 4
		c.AccessInventory(m) // Affiche l'inventaire du joueur
	default: // Si le choix est autre
		DisplayErrorSwitch() // Affiche un message d'erreur
		m.MenuFight(c)       // Affiche le menu de combat

	}
}

func (c *Character) VerifFullInventoryItems() bool { // Vérifie si l'inventaire est plein
	return len(c.Inventory) < 10 // Retourne vrai si l'inventaire est plein
}

func (c *Character) VerifFullInventoryItemsAmount(item string) bool { // Vérifie si l'inventaire est plein

	return c.Inventory[item] < 10 // Retourne vrai si l'inventaire est plein

}

// func (c Character) SpellBook() {
// 	for index, sp := range c.Skill {
// 		// Hmmmmm bonne grosse bite entre les orteils la
// 	}
// }
