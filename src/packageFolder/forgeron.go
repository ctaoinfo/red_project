// // CHUPIN Tao -- TARDELLI Andrea
// // // date
// // // YTrack Ynov Campus
// // // Red Project

package packageFolder // Package packageFolder

func (c *Character) CraftCheck(item map[string]int) bool { // Fonction de vérification des ressources                     // Variable de comptage des ressources
	var count int
	for name, quantity := range item { // Boucle de vérification des ressources
		if c.Inventory[name] < quantity { //vérifier si le joueur possède le nombre d'item demandé
			DisplayFrame(" Vous n'avez pas assez de "+name+"!", []string{ // Affichage d'un message d'erreur
				"Vous n'avez pas assez de ressources pour acheter cet objet",
				"Dirigez-vous vers le marchand pour en acheter !"})
			return false
		}
		// Si le joueur possède le nombre d'item demandé
		count++ // On incrémente le compteur
	}
	return count == len(item)
}

func (c *Character) Forgeron(choice int, m *Mob) { // Fonction du forgeron

	switch choice { // Choix du joueur
	case 1: // Cas 1
		Map := map[string]int{"PlumeCorbeau": 1, "CuirSanglier": 1}
		if c.CraftCheck(Map) { // Vérification des ressources
			for item, count := range Map { // Boucle de suppression des ressources
				c.RemoveItem(item, count)
			}
			c.AddItem("Casque de cuir")                                    // Ajouter l'item au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter un casque"}) // Affichage d'un message
		}
	case 2: // Cas 2
		Map := map[string]int{"FourrureLoup": 2, "PeauTroll": 1}
		if c.CraftCheck(Map) { // Vérification des ressources
			for item, count := range Map { // Boucle de suppression des ressources
				c.RemoveItem(item, count)
			}
			c.AddItem("Tunique de cuir")                                   // Ajouter l'item au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter une Hache"}) // Affichage d'un message
		}
	case 3: // Cas 3
		Map := map[string]int{"FourrureLoup": 1, "CuirSanglier": 1}
		if c.CraftCheck(Map) { // Vérification des ressources
			for item, count := range Map { // Boucle de suppression des ressources
				c.RemoveItem(item, count)
			}
			c.AddItem("Bottes de cuir")                                 // Ajouter l'item au joueur
			DisplayFrame("Achat", []string{"Vous avez acheter un Arc"}) // Affichage d'un message
		}
	case 4: // Cas 4
		c.MenuForgeron(m) // Retourner au menu du forgeron
	default: // Cas par défaut
		c.Forgeron(choice, m) // Retourner au menu du forgeron
	}

}
