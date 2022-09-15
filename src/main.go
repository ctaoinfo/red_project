// CHUPIN Tao -- TARDELLI Andrea
// Creation 13/09/2022
// YTrack Ynov Campus
// RedProject

package main

import "src/packageFolder" // Import du fichier src en ten que dossier de package

func main() { // Fonction principal du lancement du projet
	var c1 packageFolder.Character                                                  // Création de la variable qui stop notre premier personnage
	c1.Init("Zaganor", "Devil", 1, 100, 40, []string{"Potion", "Potion", "Potion"}) // Initialisation du premier personnage

	c1.DisplayInfo() // Affichage des infos du personnage

	for i := 1; i <= 4; i++ {
		switch i {
		case 1:
			c1.AccessInventory()
		case 2:
			c1.AccessInventory()
		case 3:
			c1.AccessInventory()
		case 4:
			c1.AccessInventory()

		}
	}
}
