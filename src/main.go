// CHUPIN Tao -- TARDELLI Andrea
// Creation 13/09/2022
// YTrack Ynov Campus
// RedProject

package main // Package mains

import "src/packageFolder" // Importation du package

// Import du fichier src en tant que dossier de package
func main() { // Fonction principal du lancement du projet
	gameStart := true // Variable jeu commencer

	// Initialisation personnage
	var c1 packageFolder.Character // Création du personnage
	c1.InitCharCreation()          // Initialisation de la création du personnage

	var m1 packageFolder.Mob // Création du mob
	m1.InitMobCreation(1)    // Initialisation de la création du mob

	for gameStart { // Boucle de jeu
		c1.PrincipalMenu(&m1) // Menu principal
	}
}
