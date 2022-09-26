// CHUPIN Tao -- TARDELLI Andrea
// Creation 13/09/2022
// YTrack Ynov Campus
// RedProject

package main

import "src/packageFolder"

// Import du fichier src en ten que dossier de package
func main() { // Fonction principal du lancement du projet
	gameStart := true

	// Initialisation personnage
	var c1 packageFolder.Character
	c1.InitCharCreation()

	for gameStart {

		c1.PrincipalMenu()

	}
}
