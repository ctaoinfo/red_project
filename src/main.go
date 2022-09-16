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
	var c1 packageFolder.Character                                                                   // Création de la variable qui stop notre premier personnage
	c1.Init("Zaganor", "Devil", 1, 100, 40, 100, []string{"Potion", "Potion", "Potion"}, []string{}) // Initialisation du premier personnage

	// Initialisation Mob
	/*
		var m1 packageFolder.Mob
		m1.InitMob("Zaganor1", "", 1, 100, 40)

		var m2 packageFolder.Mob
		m2.InitMob("Zaganor2", "Angel", 1, 100, 40)

		var m3 packageFolder.Mob
		m3.InitMob("Zaganor3", "Angel", 1, 100, 40)

		var m4 packageFolder.Mob
		m4.InitMob("Zaganor4", "Angel", 1, 100, 40)

		var m5 packageFolder.Mob
		m5.InitMob("Zaganor5", "Angel", 1, 100, 40)

		var m6 packageFolder.Mob
		m6.InitMob("Zaganor6", "Angel", 1, 100, 40)
	*/

	for gameStart {

		c1.PrincipalMenu()

	}

}
