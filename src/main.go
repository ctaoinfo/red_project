// CHUPIN Tao -- TARDELLI Andrea
// Creation 13/09/2022
// YTrack Ynov Campus
// RedProject

package main

import "src/packageFolder" // Import du fichier src en ten que dossier de package

func main() { // Fonction principal du lancement du projet
	var c1 packageFolder.Character                                                  // Création de la variable qui stop notre premier personnage
	c1.Init("Zaganor", "Devil", 1, 100, 40, []string{"Potion", "Potion", "Potion"}) // Initialisation du premier personnage

	c1.DisplayInfo()
}
