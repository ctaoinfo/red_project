// // CHUPIN Tao -- TARDELLI Andrea
// // // date
// // // YTrack Ynov Campus
// // // Red Project

package packageFolder // Package packageFolder

func (c *Character) CraftCheck(item map[string]int) bool { // Fonction de vérification des ressources
	allItem0 := 0                      // Variable de comptage des ressources
	for name, quantity := range item { // Boucle de vérification des ressources
		if c.Inventory[name] < quantity { //vérifier si le joueur possède le nombre d'item demandé
			DisplayFrame(" Vous n'avez pas assez de "+name+"!", []string{ // Affichage d'un message d'erreur
				"Vous n'avez pas assez de ressources pour acheter cet objet",
				"Dirigez-vous vers le marchand pour en acheter !"})
		} else { // Si le joueur possède le nombre d'item demandé
			if _, ok := c.Inventory[name]; ok { // Vérifier si le joueur possède l'item
				c.Inventory[name] -= quantity // Retirer l'item du joueur
				continue                      // Continuer la boucle
			}
			if (c.Inventory[name] - quantity) >= 0 { //éviter d'avoir un nombre négatif d'item
				c.Inventory[]++ // Ajouter 1 au comptage des ressources
			}
			for name, quantity := range c.Inventory { // Boucle de vérification des ressources
				c.Inventory[name] -= quantity //retirer la quantité d'item après le craft
				return true                   // Retourner vrai

			}
		}
	}
	return false // Retourner faux
}

func (c *Character) Forgeron(choice int, m *Mob) { // Fonction du forgeron

	switch choice { // Choix du joueur
	case 1: // Cas 1
		if c.CraftCheck(map[string]int{"PlumeCorbeau": 1, "CuirSanglier": 1}) { // Vérification des ressources
			c.AddItem("Casque de cuir")                                    // Ajouter l'item au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter un casque"}) // Affichage d'un message
		} else { // Si le joueur n'a pas les ressources
			DisplayFrame(" Vous n'avez pas assez de ressources!", []string{ // Affichage d'un message d'erreur
				"Vous n'avez pas assez de ressources pour acheter cet objet",
				"Dirigez-vous vers le marchand pour en acheter !"})
		}
	case 2: // Cas 2
		if c.CraftCheck(map[string]int{"FourrureLoup": 2, "PeauTroll": 1}) { // Vérification des ressources
			c.AddItem("Tunique de cuir")                                   // Ajouter l'item au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter une Hache"}) // Affichage d'un message
		} else { // Si le joueur n'a pas les ressources
			DisplayFrame(" Vous n'avez pas assez de ressources!", []string{ // Affichage d'un message d'erreur
				"Vous n'avez pas assez de ressources pour acheter cet objet",
				"Dirigez-vous vers le marchand pour en acheter !"})
		}
	case 3: // Cas 3
		if c.CraftCheck(map[string]int{"FourrureLoup": 1, "CuirSanglier": 1}) { // Vérification des ressources
			c.AddItem("Bottes de cuir")                                 // Ajouter l'item au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter un Arc"}) // Affichage d'un message
		} else { // Si le joueur n'a pas les ressources
			DisplayFrame(" Vous n'avez pas assez de ressources!", []string{ // Affichage d'un message d'erreur
				"Vous n'avez pas assez de ressources pour acheter cet objet",
				"Dirigez-vous vers le marchand pour en acheter !"})
		}
	case 4: // Cas 4
		c.MenuForgeron(m) // Retourner au menu du forgeron
	default: // Cas par défaut
		c.Forgeron(choice, m) // Retourner au menu du forgeron
	}

}
